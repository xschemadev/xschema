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
  // TODO: implement in typescript-renderer-containers task
  return "Record<string, unknown>";
}

function renderArray(node: ArrayNode): string {
  // TODO: implement in typescript-renderer-containers task
  return "unknown[]";
}

function renderTuple(node: TupleNode): string {
  // TODO: implement in typescript-renderer-containers task
  return "unknown[]";
}

function renderUnion(node: UnionNode): string {
  // TODO: implement in typescript-renderer-combinators task
  return "unknown";
}

function renderIntersection(node: IntersectionNode): string {
  // TODO: implement in typescript-renderer-combinators task
  return "unknown";
}

function renderOneOf(node: OneOfNode): string {
  // TODO: implement in typescript-renderer-combinators task
  // oneOf renders same as union at type level (no runtime distinction)
  return "unknown";
}

function renderNot(_node: NotNode): string {
  // 'not' cannot be expressed in TypeScript types
  return "unknown";
}

function renderLiteral(node: LiteralNode): string {
  // TODO: implement in typescript-renderer-special task
  return "unknown";
}

function renderEnum(node: EnumNode): string {
  // TODO: implement in typescript-renderer-special task
  return "unknown";
}

function renderRef(node: RefNode): string {
  // $ref nodes have already been resolved - render the resolved schema
  return render(node.resolved);
}

function renderConditional(node: ConditionalNode): string {
  // if/then/else cannot be expressed in TypeScript types at compile time
  // Best approximation: union of then and else types (or unknown if neither)
  // TODO: implement in typescript-renderer-special task
  return "unknown";
}

function renderTypeGuarded(node: TypeGuardedNode): string {
  // TypeGuardedNode contains type-specific schemas
  // At type level, this becomes a union of all possible types
  // TODO: implement in typescript-renderer-special task
  return "unknown";
}

function renderNullable(node: NullableNode): string {
  const inner = render(node.inner);
  return `${inner} | null`;
}
