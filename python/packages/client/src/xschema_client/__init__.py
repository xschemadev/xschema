"""Runtime client for xschema-generated validators with type-safe schema lookup."""

from typing import Any, Callable, Dict, Generic, Optional, TypeVar

__all__ = ["create_xschema_client", "XSchemaError"]

T = TypeVar("T")


class XSchemaError(Exception):
    """Raised when a schema cannot be found."""

    pass


class XSchemaClient(Generic[T]):
    """Type-safe client for looking up schemas by namespace:id."""

    def __init__(
        self, schemas: Dict[str, T], default_namespace: Optional[str] = None
    ) -> None:
        """
        Initialize the xschema client.

        Args:
            schemas: Dictionary of schemas keyed by "namespace:id"
            default_namespace: Optional default namespace for shorthand lookups
        """
        self._schemas = schemas
        self._default_namespace = default_namespace

    def __call__(self, key: str) -> T:
        """
        Look up a schema by key.

        Args:
            key: Schema key in "namespace:id" format, or just "id" if default_namespace is set

        Returns:
            The schema validator (e.g., Pydantic model)

        Raises:
            XSchemaError: If schema not found
        """
        # If key includes ":", use as-is; otherwise prepend default_namespace
        if ":" in key:
            resolved_key = key
        elif self._default_namespace:
            resolved_key = f"{self._default_namespace}:{key}"
        else:
            resolved_key = key

        if resolved_key not in self._schemas:
            msg = f"Unknown schema: {resolved_key}. Run `xschema generate`."
            raise XSchemaError(msg)

        return self._schemas[resolved_key]


def create_xschema_client(
    schemas: Dict[str, T], default_namespace: Optional[str] = None
) -> Callable[[str], T]:
    """
    Create an xschema client for looking up schemas by namespace:id.

    Provides runtime schema lookup with optional shorthand notation.

    Args:
        schemas: Dictionary of schemas keyed by "namespace:id"
        default_namespace: Optional default namespace for shorthand lookups

    Returns:
        A callable that looks up schemas by key

    Example:
        >>> from xschema_client import create_xschema_client
        >>> schemas = {"user:Profile": ProfileModel, "another:TSConfig": TSConfigModel}
        >>> xschema = create_xschema_client(schemas, default_namespace="user")
        >>>
        >>> # Full key lookup
        >>> profile = xschema("user:Profile")
        >>>
        >>> # Shorthand with default namespace
        >>> profile2 = xschema("Profile")  # Resolves to "user:Profile"
        >>>
        >>> # Validate data
        >>> profile.model_validate({"name": "Alice"})
    """
    client = XSchemaClient(schemas, default_namespace)
    return client
