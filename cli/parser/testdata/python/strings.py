# String variations - testing different quote styles
adapter = {"name": "pydantic"}

# Double quotes
xschema.from_url("DoubleQuote", "https://example.com/a.json", adapter)

# Single quotes
xschema.from_url("SingleQuote", "https://example.com/b.json", adapter)

# Triple double quotes
xschema.from_url("""TripleDouble""", """https://example.com/c.json""", adapter)

# Triple single quotes
xschema.from_url("""TripleSingle""", """https://example.com/d.json""", adapter)

# Raw string for file path
xschema.from_file("RawString", r"./schemas/raw.json", adapter)
