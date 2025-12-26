from dataclasses import dataclass
from typing import Literal
from xschema import xschema


@dataclass
class PydanticAdapter:
    name: Literal["pydantic"] = "pydantic"
    brand: Literal["xschema-adapter"] = "xschema-adapter"


pydantic_adapter = PydanticAdapter()

# From URL
xschema.from_url("User", "https://cdn.my/user.json", pydantic_adapter)
print("User:", xschema.User)

# From file
xschema.from_file("Post", "./schemas/post.json", pydantic_adapter)
print("Post:", xschema.Post)

# From inline schema
xschema.from_schema(
    "Comment",
    {
        "type": "object",
        "properties": {
            "id": {"type": "string", "format": "uuid"},
            "body": {"type": "string"},
            "author_id": {"type": "string", "format": "uuid"},
        },
        "required": ["id", "body", "author_id"],
    },
    pydantic_adapter,
)
print("Comment:", xschema.Comment)

# If haven't run the CLI, here is the DX
xschema.from_url("Unknown", "https://cdn.my/unknown.json", pydantic_adapter)
