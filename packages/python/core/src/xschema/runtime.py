from typing import Protocol, Literal


class XSchemaAdapter(Protocol):
    @property
    def name(self) -> str: ...

    @property
    def __brand__(self) -> Literal["xschema-adapter"]: ...


class XSchemaBase:
    """Base class for generated xschema namespace."""

    pass
