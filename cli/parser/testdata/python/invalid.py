# Invalid calls - these should NOT be parsed
from xschema import create_xschema_client
from xschema_pydantic import adapter

xschema = create_xschema_client({})
dynamic_name = "Dynamic"
dynamic_url = "https://example.com/dynamic.json"

# Dynamic name (variable) - correctly skipped
xschema.from_url(dynamic_name, "https://example.com/a.json", adapter)

# Dynamic URL (variable) - correctly skipped
xschema.from_url("StaticName", dynamic_url, adapter)

# Wrong object (not xschema) - correctly skipped
not_xschema.from_url("WrongObject", "https://example.com/b.json", adapter)

# Wrong method name - correctly skipped
xschema.from_uri("WrongMethod", "https://example.com/c.json", adapter)

# Missing arguments - correctly skipped
xschema.from_url("MissingArgs", adapter)

# Extra arguments (4 args instead of 3) - correctly skipped
xschema.from_url("ExtraArgs", "https://example.com/d.json", adapter, "extra")

# Nested call (xschema not direct object) - correctly skipped
foo.xschema.from_url("NestedXschema", "https://example.com/e.json", adapter)

# Keyword arguments (not positional) - correctly skipped
xschema.from_url(name="KeywordArgs", url="https://example.com/h.json", adapter=adapter)

# f-string in name - correctly skipped
xschema.from_url(f"FString{1}", "https://example.com/f.json", adapter)

# f-string in URL - correctly skipped
xschema.from_url("Valid", f"https://example.com/{1}/g.json", adapter)

# ============================================================
# KNOWN ISSUES
# ============================================================
# - Inline comments between arguments break parsing
#   xschema.from_file("Name", # comment <- breaks
#       "./path.json", adapter)
