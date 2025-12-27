// String variations - testing different quote styles
import { createXSchemaClient } from "@xschema/client";
import { adapter } from "@xschema/zod";

const xschema = createXSchemaClient({});

// Double quotes
xschema.fromURL("DoubleQuote", "https://example.com/a.json", adapter);

// Single quotes
xschema.fromURL('SingleQuote', 'https://example.com/b.json', adapter);

// Template literal without interpolation (should work)
xschema.fromURL("TemplateLit", `https://example.com/c.json`, adapter);

// Template literal with interpolation (should be SKIPPED)
const version = "v1";
xschema.fromURL("Interpolated", `https://example.com/${version}/d.json`, adapter);

// Name with interpolation (should be SKIPPED)
xschema.fromURL(`Schema${version}`, "https://example.com/e.json", adapter);
