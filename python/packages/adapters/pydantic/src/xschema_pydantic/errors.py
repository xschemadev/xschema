"""Custom exception classes for error handling."""


class XSchemaError(Exception):
    """Base exception for all xschema-pydantic errors."""

    def __init__(self, message: str, schema_path: str | None = None):
        """Initialize error with message and optional schema path.

        Args:
            message: Human-readable error message
            schema_path: JSON pointer to location in schema where error occurred
        """
        self.schema_path = schema_path
        if schema_path:
            super().__init__(f"{message} (at {schema_path})")
        else:
            super().__init__(message)


class InvalidSchemaError(XSchemaError):
    """Raised when schema is malformed or invalid."""

    pass


class UnsupportedFeatureError(XSchemaError):
    """Raised when schema uses features that cannot be compiled."""

    pass


class ConversionError(XSchemaError):
    """Raised when schema cannot be converted to Pydantic code."""

    pass
