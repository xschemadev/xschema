"""Runtime client for xschema-generated validators with type-safe schema lookup."""

from typing import Any, Callable, Dict, Generic, Optional, TypeVar

__all__ = ["create_xschema_client", "XSchemaError"]

T = TypeVar("T")


class XSchemaError(Exception):
    """Raised when a schema cannot be found."""

    pass


class XSchemaClient(Generic[T]):
    """Type-safe client for looking up schemas by namespace:id."""

    def __init__(self, schemas: Dict[str, T]) -> None:
        """
        Initialize the xschema client.

        Args:
            schemas: Dictionary of schemas keyed by "namespace:id"
        """
        self._schemas = schemas

    def __call__(self, key: str) -> T:
        """
        Look up a schema by key.

        Args:
            key: Schema key in "namespace:id" format

        Returns:
            The schema validator (e.g., Pydantic model)

        Raises:
            XSchemaError: If schema not found
        """
        if key not in self._schemas:
            msg = f"Unknown schema: {key}. Run `xschema generate`."
            raise XSchemaError(msg)

        return self._schemas[key]


def create_xschema_client(schemas: Dict[str, T]) -> Callable[[str], T]:
    """
    Create an xschema client for looking up schemas by namespace:id.

    Provides runtime schema lookup with type safety.

    Args:
        schemas: Dictionary of schemas keyed by "namespace:id"

    Returns:
        A callable that looks up schemas by key

    Example:
        >>> from xschema_client import create_xschema_client
        >>> schemas = {"user:Profile": ProfileModel, "another:TSConfig": TSConfigModel}
        >>> xschema = create_xschema_client(schemas)
        >>>
        >>> # Look up by full key
        >>> profile = xschema("user:Profile")
        >>>
        >>> # Validate data
        >>> profile.model_validate({"name": "Alice"})
    """
    client = XSchemaClient(schemas)
    return client
