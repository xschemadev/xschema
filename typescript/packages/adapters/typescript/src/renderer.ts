import type {
  ArrayNode,
  ConditionalNode,
  EnumNode,
  IntersectionNode,
  LiteralNode,
  NotNode,
  NullableNode,
  NumberNode,
  ObjectNode,
  OneOfNode,
  RefNode,
  SchemaNode,
  StringNode,
  TupleNode,
  TypeGuardedNode,
  UnionNode,
} from "@xschemadev/core";

/**
 * Renders a SchemaNode IR to a TypeScript type expression.
 *
 * This is a TYPE-ONLY renderer - it generates TypeScript types, not runtime validation code.
 * Runtime constraints (minLength, pattern, minimum, etc.) cannot be expressed in TS types
 * and are ignored. Only structural types are rendered.
 */
export function render(node: SchemaNode): string {
  switch (node.kind) {
    case "string":
      return renderString(node);
    case "number":
      return renderNumber(node);
    case "boolean":
      return "boolean";
    case "null":
      return "null";
    case "object":
      return renderObject(node);
    case "array":
      return renderArray(node);
    case "tuple":
      return renderTuple(node);
    case "union":
      return renderUnion(node);
    case "intersection":
      return renderIntersection(node);
    case "oneOf":
      return renderOneOf(node);
    case "not":
      return renderNot(node);
    case "literal":
      return renderLiteral(node);
    case "enum":
      return renderEnum(node);
    case "any":
      return "unknown";
    case "never":
      return "never";
    case "ref":
      return renderRef(node);
    case "conditional":
      return renderConditional(node);
    case "typeGuarded":
      return renderTypeGuarded(node);
    case "nullable":
      return renderNullable(node);
    default: {
      const _exhaustive: never = node;
      throw new Error(`Unknown node kind: ${(_exhaustive as SchemaNode).kind}`);
    }
  }
}

// Constraints cannot be expressed in TS types - return base type
function renderString(_node: StringNode): string {
  return "string";
}

// Constraints cannot be expressed in TS types - return base type
function renderNumber(node: NumberNode): string {
  // integer constraint cannot be enforced at type level
  return "number";
}

function renderObject(node: ObjectNode): string {
  const propKeys = Array.from(node.properties.keys());

  // No defined properties
  if (propKeys.length === 0) {
    // Pure record type based on additionalProperties
    if (node.additionalProperties === false) {
      // No properties allowed - empty object
      return "Record<string, never>";
    } else if (
      typeof node.additionalProperties === "object" &&
      node.additionalProperties.kind !== "any"
    ) {
      const valueType = render(node.additionalProperties);
      return `Record<string, ${valueType}>`;
    } else {
      return "Record<string, unknown>";
    }
  }

  // Build property types
  const props = propKeys.map((key) => {
    const prop = node.properties.get(key)!;
    const propType = render(prop.schema as SchemaNode);
    const safeKey = needsQuotes(key) ? JSON.stringify(key) : key;
    const optional = prop.required ? "" : "?";
    return `${safeKey}${optional}: ${propType}`;
  });

  // Handle additional properties
  let indexSignature = "";
  if (node.additionalProperties === true) {
    indexSignature = "[key: string]: unknown";
  } else if (
    typeof node.additionalProperties === "object" &&
    node.additionalProperties.kind !== "any"
  ) {
    const valueType = render(node.additionalProperties);
    indexSignature = `[key: string]: ${valueType}`;
  }
  // additionalProperties: false means strict object - no index signature

  const allMembers = indexSignature ? [...props, indexSignature] : props;
  return `{ ${allMembers.join("; ")} }`;
}

// Check if a property key needs to be quoted
function needsQuotes(key: string): boolean {
  // Valid JS identifier regex (simplified)
  return !/^[a-zA-Z_$][a-zA-Z0-9_$]*$/.test(key);
}

function renderArray(node: ArrayNode): string {
  // Runtime constraints (minItems, maxItems, uniqueItems, contains) cannot be
  // expressed in TypeScript types - only structural type is rendered
  const itemType = render(node.items);
  return `${itemType}[]`;
}

function renderTuple(node: TupleNode): string {
  // Runtime constraints (minItems, maxItems, uniqueItems, contains) cannot be
  // expressed in TypeScript types
  const prefixTypes = node.prefixItems.map((item) => render(item));

  if (node.restItems === false) {
    // Strict tuple - no additional items allowed
    return `[${prefixTypes.join(", ")}]`;
  }

  // Variadic tuple with rest element
  const restType = render(node.restItems);
  if (prefixTypes.length === 0) {
    // Just rest items = array
    return `${restType}[]`;
  }
  return `[${prefixTypes.join(", ")}, ...${restType}[]]`;
}

function renderUnion(node: UnionNode): string {
  if (node.variants.length === 0) {
    return "never";
  }
  if (node.variants.length === 1) {
    return render(node.variants[0]);
  }
  const types = node.variants.map((s) => render(s));
  return types.join(" | ");
}

function renderIntersection(node: IntersectionNode): string {
  if (node.schemas.length === 0) {
    return "unknown";
  }
  if (node.schemas.length === 1) {
    return render(node.schemas[0]);
  }
  const schemas = node.schemas.map((s) => render(s));
  return schemas.join(" & ");
}

function renderOneOf(node: OneOfNode): string {
  // oneOf renders same as union at type level - runtime "exactly one" constraint
  // cannot be expressed in TypeScript types
  if (node.schemas.length === 0) {
    return "never";
  }
  if (node.schemas.length === 1) {
    return render(node.schemas[0]);
  }
  const variants = node.schemas.map((s) => render(s));
  return variants.join(" | ");
}

function renderNot(_node: NotNode): string {
  // 'not' cannot be expressed in TypeScript types
  return "unknown";
}

function renderLiteral(node: LiteralNode): string {
  return valueToLiteralType(node.value);
}

/**
 * Converts a JavaScript value to its TypeScript literal type representation.
 * - strings: "hello" → '"hello"'
 * - numbers: 42 → '42'
 * - booleans: true → 'true'
 * - null: null → 'null'
 * - arrays: [1, 2] → 'readonly [1, 2]'
 * - objects: {a: 1} → '{ readonly a: 1 }'
 */
function valueToLiteralType(value: unknown): string {
  if (value === null) {
    return "null";
  }
  if (typeof value === "string") {
    return JSON.stringify(value);
  }
  if (typeof value === "number" || typeof value === "boolean") {
    return String(value);
  }
  if (Array.isArray(value)) {
    if (value.length === 0) {
      return "readonly []";
    }
    const elements = value.map((v) => valueToLiteralType(v));
    return `readonly [${elements.join(", ")}]`;
  }
  if (typeof value === "object") {
    const entries = Object.entries(value);
    if (entries.length === 0) {
      return "Record<string, never>";
    }
    const props = entries.map(([k, v]) => {
      const safeKey = needsQuotes(k) ? JSON.stringify(k) : k;
      return `readonly ${safeKey}: ${valueToLiteralType(v)}`;
    });
    return `{ ${props.join("; ")} }`;
  }
  // Fallback for undefined or other types
  return "unknown";
}

function renderEnum(node: EnumNode): string {
  if (node.values.length === 0) {
    return "never";
  }
  if (node.values.length === 1) {
    return valueToLiteralType(node.values[0]);
  }
  // Union of literal types
  const literals = node.values.map((v) => valueToLiteralType(v));
  return literals.join(" | ");
}

function renderRef(node: RefNode): string {
  // $ref nodes have already been resolved - render the resolved schema
  return render(node.resolved);
}

function renderConditional(node: ConditionalNode): string {
  // if/then/else is runtime validation - cannot be expressed in TypeScript types.
  // Best approximation: union of then and else types (covers possible output types).
  const types: string[] = [];
  if (node.then) {
    types.push(render(node.then));
  }
  if (node.else) {
    types.push(render(node.else));
  }
  if (types.length === 0) {
    return "unknown";
  }
  if (types.length === 1) {
    return types[0];
  }
  return types.join(" | ");
}

function renderTypeGuarded(node: TypeGuardedNode): string {
  // TypeGuardedNode contains type-specific schemas that apply based on runtime type.
  // At type level: union of all guarded schemas (all possible types the value could be).
  if (node.guards.length === 0) {
    return "unknown";
  }
  if (node.guards.length === 1) {
    return render(node.guards[0].schema);
  }
  const types = node.guards.map((g) => render(g.schema));
  return types.join(" | ");
}

function renderNullable(node: NullableNode): string {
  const inner = render(node.inner);
  return `${inner} | null`;
}
