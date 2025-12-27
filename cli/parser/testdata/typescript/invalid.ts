// Invalid calls - these should NOT be parsed
const adapter = { name: "zod" };
const dynamicName = "Dynamic";

// Dynamic name (variable) - correctly skipped
xschema.fromURL(dynamicName, "https://example.com/a.json", adapter);

// Wrong object (not xschema) - correctly skipped
notXschema.fromURL("WrongObject", "https://example.com/b.json", adapter);

// Wrong method name - correctly skipped
xschema.fromURI("WrongMethod", "https://example.com/c.json", adapter);

// Missing arguments - correctly skipped
xschema.fromURL("MissingArgs", adapter);

// Extra arguments (4 args instead of 3) - correctly skipped
xschema.fromURL("ExtraArgs", "https://example.com/d.json", adapter, "extra");

// Nested call (xschema not direct object) - correctly skipped
foo.xschema.fromURL("NestedXschema", "https://example.com/e.json", adapter);

// Template literal with interpolation in name - correctly skipped
xschema.fromURL(`User${1}`, "https://example.com/f.json", adapter);

// Template literal with interpolation in URL - correctly skipped
xschema.fromURL("Valid", `https://example.com/${1}/g.json`, adapter);

// ============================================================
// KNOWN ISSUES: The following ARE parsed but SHOULDN'T be
// ============================================================

// Dynamic URL (variable) - BUG: captured as loc="dynamicURL"
// xschema.fromURL("StaticName", dynamicURL, adapter);
