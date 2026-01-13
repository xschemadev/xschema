// Adapter protocol types (existing)
export type { ConvertInput, ConvertResult } from "./types.js";
export { createAdapterCLI } from "./cli.js";

// JSON Schema types
export type { JSONSchema } from "./schema/index.js";

// Intermediate Representation (IR)
export type {
  // Main types
  SchemaNode,
  NodeKind,
  NodeOfKind,

  // Node types
  StringNode,
  NumberNode,
  BooleanNode,
  NullNode,
  ObjectNode,
  ArrayNode,
  TupleNode,
  UnionNode,
  IntersectionNode,
  OneOfNode,
  NotNode,
  LiteralNode,
  EnumNode,
  AnyNode,
  NeverNode,
  RefNode,
  ConditionalNode,
  TypeGuardedNode,
  NullableNode,
  TypeGuard,
  TypeCheck,

  // Constraints
  StringConstraints,
  NumberConstraints,
  ArrayConstraints,
  ContainsConstraint,
  PropertyDef,
  PatternPropertyDef,
  Dependency,
  PropertyDependency,
  SchemaDependency,
} from "./ir/index.js";

// Parser
export { parse, parseSchema, type ParseContext } from "./parser/index.js";

// Utilities
export {
  // Primitives
  isPrimitive,
  escapeString,
  isEmptyObject,
  getOwnProperty,
  PROTOTYPE_PROPERTY_NAMES,
  hasPrototypeProperties,

  // JSON Pointer
  resolveJsonPointer,
  getRefName,

  // Code builder
  chain,
  buildUnion,
  buildIntersection,
  buildSuperRefine,
  buildRefine,
  buildLiteral,
  buildSafeParseCheck,
  buildPropertyCheck,
  sortedStringify,
} from "./utils/index.js";
