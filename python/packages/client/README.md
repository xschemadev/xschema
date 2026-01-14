# xschema-client

Runtime client for xschema-generated validators with type-safe schema lookup.

## Installation

```bash
pip install xschema-client
```

## Usage

```python
from xschema_client import create_xschema_client
from .xschema.xschema_gen import schemas

# Create client with default namespace
xschema = create_xschema_client(schemas, default_namespace="user")

# Full key lookup
user_validator = xschema("user:Profile")

# Shorthand with default namespace
profile_validator = xschema("Profile")  # Resolves to "user:Profile"

# Validate data
profile_validator.model_validate({"name": "Alice", "email": "alice@example.com"})
```
