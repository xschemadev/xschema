"""Pydantic Renderer - Converts SchemaNode IR to Pydantic code strings."""

from dataclasses import dataclass, field
from typing import Any

from xschema_core import (
    SchemaNode,
    StringNode,
    NumberNode,
    BooleanNode,
    NullNode,
    LiteralNode,
    EnumNode,
    AnyNode,
    NeverNode,
    ArrayNode,
    TupleNode,
    ObjectNode,
    UnionNode,
    IntersectionNode,
    RefNode,
)


@dataclass
class RenderResult:
    """Result of rendering a schema node."""

    code: str  # Generated Python code (class definition or type alias)
    type_expr: str  # Type expression (e.g., "str", "MyClass", "list[int]")
    imports: set[str] = field(default_factory=set)  # Required imports


def render(node: SchemaNode, name: str) -> RenderResult:
    """Render a SchemaNode to Pydantic code.

    Args:
        node: The IR node to render
        name: The name to use for the generated type (PascalCase)

    Returns:
        RenderResult with generated code, type expression, and required imports
    """
    match node.kind:
        case "string":
            return render_string(node)
        case "number":
            return render_number(node)
        case "boolean":
            return render_boolean(node)
        case "null":
            return render_null(node)
        case "literal":
            return render_literal(node)
        case "enum":
            return render_enum(node)
        case "any":
            return render_any(node)
        case "never":
            return render_never(node)
        case "array":
            return render_array(node, name)
        case "tuple":
            return render_tuple(node, name)
        case "object":
            return render_object(node, name)
        case "union":
            return render_union(node, name)
        case "intersection":
            return render_intersection(node, name)
        case "ref":
            return render_ref(node, name)
        case _:
            # Fallback for unknown node types
            return RenderResult(code="", type_expr="Any", imports={"from typing import Any"})


def render_string(node: StringNode) -> RenderResult:
    """Render StringNode to Pydantic type."""
    imports: set[str] = set()
    has_constraints = (
        node.min_length is not None
        or node.max_length is not None
        or node.pattern is not None
    )

    if has_constraints:
        # Use Annotated with StringConstraints for constrained strings
        imports.add("from typing import Annotated")
        imports.add("from pydantic import StringConstraints")

        constraints: list[str] = []
        if node.min_length is not None:
            constraints.append(f"min_length={node.min_length}")
        if node.max_length is not None:
            constraints.append(f"max_length={node.max_length}")
        if node.pattern is not None:
            # Escape the pattern for Python string
            escaped_pattern = node.pattern.replace("\\", "\\\\").replace("'", "\\'")
            constraints.append(f"pattern=r'{escaped_pattern}'")

        constraint_str = ", ".join(constraints)
        type_expr = f"Annotated[str, StringConstraints({constraint_str})]"
    else:
        type_expr = "str"

    return RenderResult(code="", type_expr=type_expr, imports=imports)


def render_number(node: NumberNode) -> RenderResult:
    """Render NumberNode to Pydantic type."""
    imports: set[str] = set()
    base_type = "int" if node.integer else "float"

    has_constraints = (
        node.minimum is not None
        or node.maximum is not None
        or node.exclusive_minimum is not None
        or node.exclusive_maximum is not None
        or node.multiple_of is not None
    )

    if has_constraints:
        # Use Annotated with annotated-types constraints
        imports.add("from typing import Annotated")
        annotations: list[str] = []

        if node.minimum is not None:
            imports.add("from annotated_types import Ge")
            annotations.append(f"Ge({node.minimum})")
        if node.exclusive_minimum is not None:
            imports.add("from annotated_types import Gt")
            annotations.append(f"Gt({node.exclusive_minimum})")
        if node.maximum is not None:
            imports.add("from annotated_types import Le")
            annotations.append(f"Le({node.maximum})")
        if node.exclusive_maximum is not None:
            imports.add("from annotated_types import Lt")
            annotations.append(f"Lt({node.exclusive_maximum})")
        if node.multiple_of is not None:
            imports.add("from annotated_types import MultipleOf")
            annotations.append(f"MultipleOf({node.multiple_of})")

        annotation_str = ", ".join(annotations)
        type_expr = f"Annotated[{base_type}, {annotation_str}]"
    else:
        type_expr = base_type

    return RenderResult(code="", type_expr=type_expr, imports=imports)


def render_boolean(node: BooleanNode) -> RenderResult:
    """Render BooleanNode to Pydantic type."""
    return RenderResult(code="", type_expr="bool", imports=set())


def render_null(node: NullNode) -> RenderResult:
    """Render NullNode to Pydantic type."""
    return RenderResult(code="", type_expr="None", imports=set())


def render_literal(node: LiteralNode) -> RenderResult:
    """Render LiteralNode to Pydantic type."""
    imports: set[str] = {"from typing import Literal"}

    # Format the literal value appropriately
    value = node.value
    if value is None:
        literal_repr = "None"
    elif isinstance(value, bool):
        literal_repr = "True" if value else "False"
    elif isinstance(value, str):
        # Escape quotes in string
        escaped = value.replace("\\", "\\\\").replace("'", "\\'")
        literal_repr = f"'{escaped}'"
    else:
        literal_repr = repr(value)

    type_expr = f"Literal[{literal_repr}]"
    return RenderResult(code="", type_expr=type_expr, imports=imports)


def render_enum(node: EnumNode) -> RenderResult:
    """Render EnumNode to Pydantic type using Literal."""
    imports: set[str] = {"from typing import Literal"}

    # Build Literal type with all enum values
    literal_values: list[str] = []
    for value in node.values:
        if value is None:
            literal_values.append("None")
        elif isinstance(value, bool):
            literal_values.append("True" if value else "False")
        elif isinstance(value, str):
            escaped = value.replace("\\", "\\\\").replace("'", "\\'")
            literal_values.append(f"'{escaped}'")
        else:
            literal_values.append(repr(value))

    values_str = ", ".join(literal_values)
    type_expr = f"Literal[{values_str}]"
    return RenderResult(code="", type_expr=type_expr, imports=imports)


def render_any(node: AnyNode) -> RenderResult:
    """Render AnyNode to Pydantic type."""
    imports: set[str] = {"from typing import Any"}
    return RenderResult(code="", type_expr="Any", imports=imports)


def _never_validator(v: Any) -> Any:
    """Validator that always fails - used for Never type."""
    raise ValueError("Never type: no value is valid")


def render_never(node: NeverNode) -> RenderResult:
    """Render NeverNode to Pydantic type.

    Note: Pydantic's TypeAdapter doesn't support typing.Never directly.
    We use Annotated with BeforeValidator that always fails for runtime validation.
    """
    imports: set[str] = {
        "from typing import Annotated, Any",
        "from pydantic import BeforeValidator",
    }
    # The validator function needs to be defined in the generated code
    # We generate it inline since we need it to be available
    code = '''def _never_validator(v):
    raise ValueError("Never type: no value is valid")'''
    type_expr = "Annotated[Any, BeforeValidator(_never_validator)]"
    return RenderResult(code=code, type_expr=type_expr, imports=imports)


# Placeholder implementations for complex types (to be implemented in subsequent tasks)

def render_array(node: ArrayNode, name: str) -> RenderResult:
    """Render ArrayNode to Pydantic type (placeholder)."""
    # Will be fully implemented in adapter-renderer-arrays task
    item_result = render(node.items, f"{name}Item")
    imports = item_result.imports.copy()
    type_expr = f"list[{item_result.type_expr}]"
    return RenderResult(code=item_result.code, type_expr=type_expr, imports=imports)


def render_tuple(node: TupleNode, name: str) -> RenderResult:
    """Render TupleNode to Pydantic type (placeholder)."""
    # Will be fully implemented in adapter-renderer-arrays task
    imports: set[str] = set()
    item_types: list[str] = []

    for i, item in enumerate(node.prefix_items):
        item_result = render(item, f"{name}Item{i}")
        imports.update(item_result.imports)
        item_types.append(item_result.type_expr)

    type_expr = f"tuple[{', '.join(item_types)}]"
    return RenderResult(code="", type_expr=type_expr, imports=imports)


def render_object(node: ObjectNode, name: str) -> RenderResult:
    """Render ObjectNode to Pydantic type (placeholder)."""
    # Will be fully implemented in adapter-renderer-objects task
    imports: set[str] = {"from pydantic import BaseModel"}
    code = f"class {name}(BaseModel):\n    pass  # TODO: implement properties"
    return RenderResult(code=code, type_expr=name, imports=imports)


def render_union(node: UnionNode, name: str) -> RenderResult:
    """Render UnionNode to Pydantic type (placeholder)."""
    # Will be fully implemented in adapter-renderer-unions task
    imports: set[str] = set()
    variant_types: list[str] = []

    for i, variant in enumerate(node.variants):
        variant_result = render(variant, f"{name}Variant{i}")
        imports.update(variant_result.imports)
        variant_types.append(variant_result.type_expr)

    type_expr = " | ".join(variant_types)
    return RenderResult(code="", type_expr=type_expr, imports=imports)


def render_intersection(node: IntersectionNode, name: str) -> RenderResult:
    """Render IntersectionNode to Pydantic type (placeholder)."""
    # Will be fully implemented in adapter-renderer-unions task
    # For now, just render the first schema
    if node.schemas:
        return render(node.schemas[0], name)
    imports: set[str] = {"from typing import Any"}
    return RenderResult(code="", type_expr="Any", imports=imports)


def render_ref(node: RefNode, name: str) -> RenderResult:
    """Render RefNode to Pydantic type."""
    # If resolved, render the resolved schema
    if node.resolved is not None:
        return render(node.resolved, name)

    # Otherwise, just use the path as a forward reference
    # Extract the type name from the path (last segment after /)
    ref_name = node.path.split("/")[-1]
    return RenderResult(code="", type_expr=f"'{ref_name}'", imports=set())
