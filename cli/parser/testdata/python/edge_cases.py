# Edge cases - whitespace, comments, multiline
from xschema import create_xschema_client
from xschema_pydantic import adapter

xschema = create_xschema_client({})

# Extra whitespace
xschema.from_url("ExtraSpaces", "https://example.com/spaces.json", adapter)

# Multiline call (no inline comments)
xschema.from_url("Multiline", "https://example.com/multiline.json", adapter)

# Very long URL
xschema.from_url(
    "LongURL",
    "https://very-long-domain-name.example.com/api/v1/schemas/users/definitions/extended-profile.json",
    adapter,
)

# Unicode in name
xschema.from_url("Schéma", "https://example.com/unicode.json", adapter)

# ============================================================
# KNOWN ISSUES - things that DON'T work in Python:
# ============================================================
# - Inline comments between arguments break parsing
#   xschema.from_file("Name", # comment <- breaks
#       "./path.json", adapter)
# - Raw strings (r"...") captured with r prefix
# - Triple quotes have extra quotes in value
# - f-strings are incorrectly captured
