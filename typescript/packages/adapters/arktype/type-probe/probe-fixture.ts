/**
 * ArkType Type Probe Fixture - AUTO-GENERATED
 *
 * This file contains generated ArkType code for representative schemas of each
 * targeted construct. The inferred types reveal where type.unknown degrades the
 * TypeScript type information.
 *
 * Generated: 2026-02-11
 * Do NOT edit manually - regenerate with: bun run tasks/type-fidelity-baseline/arktype/generate-probe.ts
 */
import { type } from "arktype";

// oneOf with string | number - should ideally be string | number, currently unknown
const probe_oneOfStringOrNumber = type.unknown.narrow((val, ctx) => {
    const schemas = [type.string, type.number];
    const validCount = schemas.filter(s => s.allows(val)).length;
    if (validCount === 0) return ctx.mustBe("matching exactly one schema (matched none)");
    if (validCount > 1) return ctx.mustBe("matching exactly one schema (matched multiple)");
    return true;
  });
type Probe_oneOfStringOrNumber = typeof probe_oneOfStringOrNumber.infer;

// oneOf with discriminated objects - currently unknown
const probe_oneOfObjects = type.unknown.narrow((val, ctx) => {
    const schemas = [type({ kind: type.unit("a"), value: type.string }).narrow((val, ctx) => !Array.isArray(val) || ctx.mustBe("an object, not an array")), type({ kind: type.unit("b"), value: type.number }).narrow((val, ctx) => !Array.isArray(val) || ctx.mustBe("an object, not an array"))];
    const validCount = schemas.filter(s => s.allows(val)).length;
    if (validCount === 0) return ctx.mustBe("matching exactly one schema (matched none)");
    if (validCount > 1) return ctx.mustBe("matching exactly one schema (matched multiple)");
    return true;
  });
type Probe_oneOfObjects = typeof probe_oneOfObjects.infer;

// not string - should be unknown, currently unknown
const probe_notString = type.unknown.narrow((val, ctx) => !type.string.allows(val) || ctx.mustBe("not matching the excluded schema"));
type Probe_notString = typeof probe_notString.infer;

// not boolean - should be unknown, currently unknown
const probe_notBoolean = type.unknown.narrow((val, ctx) => !type.boolean.allows(val) || ctx.mustBe("not matching the excluded schema"));
type Probe_notBoolean = typeof probe_notBoolean.infer;

// if/then/else - should be unknown, currently unknown
const probe_conditionalIfThenElse = type.unknown.narrow((val, ctx) => {
      if (type({ kind: type.unit("a") }).narrow((val, ctx) => !Array.isArray(val) || ctx.mustBe("an object, not an array")).allows(val)) {
        return type({ kind: type.unit("a"), value: type.string }).narrow((val, ctx) => !Array.isArray(val) || ctx.mustBe("an object, not an array")).allows(val) || ctx.mustBe("valid for then branch");
      } else {
        return type({ kind: type.unit("b"), value: type.number }).narrow((val, ctx) => !Array.isArray(val) || ctx.mustBe("an object, not an array")).allows(val) || ctx.mustBe("valid for else branch");
      }
    });
type Probe_conditionalIfThenElse = typeof probe_conditionalIfThenElse.infer;

// if/then only - should be unknown, currently unknown
const probe_conditionalIfThen = type.unknown.narrow((val, ctx) => {
      if (type.string.narrow((s, ctx) => ([...new Intl.Segmenter().segment(s)].length >= 1) || ctx.mustBe("a string with valid length")).allows(val)) {
        return type.string.narrow((s, ctx) => ([...new Intl.Segmenter().segment(s)].length <= 10) || ctx.mustBe("a string with valid length")).allows(val) || ctx.mustBe("valid for then branch");
      }
      return true;
    });
type Probe_conditionalIfThen = typeof probe_conditionalIfThen.infer;

// type-guarded object properties without explicit type - currently unknown
const probe_typeGuardedObject = type.unknown.narrow((val, ctx) => {
    if (typeof val === "object" && val !== null && !Array.isArray(val)) {
          if (!type({ name: type.string }).narrow((val, ctx) => !Array.isArray(val) || ctx.mustBe("an object, not an array")).allows(val)) return ctx.mustBe("a valid object");
        }
    return true;
  });
type Probe_typeGuardedObject = typeof probe_typeGuardedObject.infer;

// type-guarded array constraint without explicit type - currently unknown
const probe_typeGuardedArrayMinItems = type.unknown.narrow((val, ctx) => {
    if (Array.isArray(val)) {
          if (!type.unknown.array().narrow((arr, ctx) => arr.length >= 1 || ctx.mustBe("an array with at least 1 items")).allows(val)) return ctx.mustBe("a valid array");
        }
    return true;
  });
type Probe_typeGuardedArrayMinItems = typeof probe_typeGuardedArrayMinItems.infer;

// tuple [string, number] - should be [string, number, ...unknown[]], currently unknown[]
const probe_tupleStringNumber = type.unknown.array().narrow((arr, ctx) => {
      const schemas = [type.string, type.number];
      for (let i = 0; i < Math.min(arr.length, schemas.length); i++) {
        if (!schemas[i].allows(arr[i])) return ctx.mustBe("valid tuple items");
      }
      return true;
    });
type Probe_tupleStringNumber = typeof probe_tupleStringNumber.infer;

// closed tuple [string, number] - should be [string, number], currently unknown[]
const probe_tupleStringNumberClosed = type.unknown.array().narrow((arr, ctx) => {
      const schemas = [type.string, type.number];
      if (arr.length > schemas.length) return ctx.mustBe("an array with at most " + schemas.length + " items");
      for (let i = 0; i < arr.length; i++) {
        if (!schemas[i].allows(arr[i])) return ctx.mustBe("valid tuple items");
      }
      return true;
    });
type Probe_tupleStringNumberClosed = typeof probe_tupleStringNumberClosed.infer;

// tuple [string, number, ...boolean[]] - currently unknown[]
const probe_tupleMixedWithRest = type.unknown.array().narrow((arr, ctx) => {
      const schemas = [type.string, type.number];
      for (let i = 0; i < Math.min(arr.length, schemas.length); i++) {
        if (!schemas[i].allows(arr[i])) return ctx.mustBe("valid tuple items");
      }
      const restSchema = type.boolean;
      for (let i = schemas.length; i < arr.length; i++) {
        if (!restSchema.allows(arr[i])) return ctx.mustBe("valid rest items");
      }
      return true;
    });
type Probe_tupleMixedWithRest = typeof probe_tupleMixedWithRest.infer;

// const object - should be narrow type, currently unknown
const probe_constObject = type.unknown.narrow((val, ctx) => JSON.stringify(val, Object.keys(val as object).sort()) === "{\"age\":30,\"name\":\"alice\"}" || ctx.mustBe("equal to the const value"));
type Probe_constObject = typeof probe_constObject.infer;

// const array - should be narrow type, currently unknown
const probe_constArray = type.unknown.narrow((val, ctx) => JSON.stringify(val, Object.keys(val as object).sort()) === "[1,2,3]" || ctx.mustBe("equal to the const value"));
type Probe_constArray = typeof probe_constArray.infer;

// enum with object values - currently unknown
const probe_enumWithObjects = type.unknown.narrow((val, ctx) => {
      const normalized = JSON.stringify(val, val != null && typeof val === 'object' ? Object.keys(val).sort() : undefined);
      const validValues = ["{\"a\":1}", "{\"b\":2}", "\"simple\""];
      return validValues.includes(normalized) || ctx.mustBe("one of the enum values");
    });
type Probe_enumWithObjects = typeof probe_enumWithObjects.infer;

// enum with array values - currently unknown
const probe_enumWithArrays = type.unknown.narrow((val, ctx) => {
      const normalized = JSON.stringify(val, val != null && typeof val === 'object' ? Object.keys(val).sort() : undefined);
      const validValues = ["[1,2]", "[3,4]", "null"];
      return validValues.includes(normalized) || ctx.mustBe("one of the enum values");
    });
type Probe_enumWithArrays = typeof probe_enumWithArrays.infer;

// Export all probe types for external inspection
export type {
  Probe_oneOfStringOrNumber,
  Probe_oneOfObjects,
  Probe_notString,
  Probe_notBoolean,
  Probe_conditionalIfThenElse,
  Probe_conditionalIfThen,
  Probe_typeGuardedObject,
  Probe_typeGuardedArrayMinItems,
  Probe_tupleStringNumber,
  Probe_tupleStringNumberClosed,
  Probe_tupleMixedWithRest,
  Probe_constObject,
  Probe_constArray,
  Probe_enumWithObjects,
  Probe_enumWithArrays,
};
