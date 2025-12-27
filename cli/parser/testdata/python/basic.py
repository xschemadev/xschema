# Basic xschema calls - all should be parsed
from xschema import create_xschema_client
from xschema_pydantic import pydantic_adapter

xschema = create_xschema_client({})

# from_url with double quotes
xschema.from_url("User", "https://api.example.com/user.json", pydantic_adapter)

# from_file with double quotes
xschema.from_file("Post", "./schemas/post.json", pydantic_adapter)
