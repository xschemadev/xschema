# Basic xschema calls - all should be parsed
pydantic_adapter = {"name": "pydantic", "__brand": "xschema-adapter"}

# from_url with double quotes
xschema.from_url("User", "https://api.example.com/user.json", pydantic_adapter)

# from_file with double quotes
xschema.from_file("Post", "./schemas/post.json", pydantic_adapter)
