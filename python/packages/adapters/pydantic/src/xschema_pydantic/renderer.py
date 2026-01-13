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
    OneOfNode,
    IntersectionNode,
    NotNode,
    ConditionalNode,
    TypeGuardedNode,
    NullableNode,
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
        case "oneOf":
            # Placeholder: treat like union for now
            return render_union(UnionNode(variants=node.schemas), name)
        case "intersection":
            return render_intersection(node, name)
        case "not":
            # Placeholder: will implement in later task
            return RenderResult(
                code="", type_expr="Any", imports={"from typing import Any"}
            )
        case "conditional":
            # Placeholder: will implement in later task
            return RenderResult(
                code="", type_expr="Any", imports={"from typing import Any"}
            )
        case "typeGuarded":
            # Placeholder: will implement in later task
            return RenderResult(
                code="", type_expr="Any", imports={"from typing import Any"}
            )
        case "nullable":
            # Handle nullable by wrapping inner type with | None
            inner_result = render(node.inner, name)
            if inner_result.type_expr.endswith(" | None"):
                return inner_result
            return RenderResult(
                code=inner_result.code,
                type_expr=f"{inner_result.type_expr} | None",
                imports=inner_result.imports,
            )
        case "ref":
            return render_ref(node, name)
        case _:
            # Fallback for unknown node types
            return RenderResult(
                code="", type_expr="Any", imports={"from typing import Any"}
            )


def render_string(node: StringNode) -> RenderResult:
    """Render StringNode to Pydantic type."""
    imports: set[str] = set()
    has_constraints = node.constraints is not None and (
        node.constraints.min_length is not None
        or node.constraints.max_length is not None
        or node.constraints.pattern is not None
    )

    if has_constraints:
        # Use Annotated with StringConstraints for constrained strings
        imports.add("from typing import Annotated")
        imports.add("from pydantic import StringConstraints")

        constraints: list[str] = []
        if node.constraints.min_length is not None:
            constraints.append(f"min_length={node.constraints.min_length}")
        if node.constraints.max_length is not None:
            constraints.append(f"max_length={node.constraints.max_length}")
        if node.constraints.pattern is not None:
            # Escape the pattern for Python string
            escaped_pattern = node.constraints.pattern.replace("\\", "\\\\").replace(
                "'", "\\'"
            )
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

    has_constraints = node.constraints is not None and (
        node.constraints.minimum is not None
        or node.constraints.maximum is not None
        or node.constraints.exclusive_minimum is not None
        or node.constraints.exclusive_maximum is not None
        or node.constraints.multiple_of is not None
    )

    if has_constraints:
        # Use Annotated with annotated-types constraints
        imports.add("from typing import Annotated")
        annotations: list[str] = []

        if node.constraints.minimum is not None:
            imports.add("from annotated_types import Ge")
            annotations.append(f"Ge({node.constraints.minimum})")
        if node.constraints.exclusive_minimum is not None:
            imports.add("from annotated_types import Gt")
            annotations.append(f"Gt({node.constraints.exclusive_minimum})")
        if node.constraints.maximum is not None:
            imports.add("from annotated_types import Le")
            annotations.append(f"Le({node.constraints.maximum})")
        if node.constraints.exclusive_maximum is not None:
            imports.add("from annotated_types import Lt")
            annotations.append(f"Lt({node.constraints.exclusive_maximum})")
        if node.constraints.multiple_of is not None:
            imports.add("from annotated_types import MultipleOf")
            annotations.append(f"MultipleOf({node.constraints.multiple_of})")

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
    code = """def _never_validator(v):
    raise ValueError("Never type: no value is valid")"""
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
    """Render ObjectNode to Pydantic BaseModel class."""
    imports: set[str] = {"from pydantic import BaseModel"}
    nested_classes: list[str] = []
    field_lines: list[str] = []

    # Process each property
    for prop_name, prop_def in node.properties:
        # Render the property's schema
        # Use PascalCase nested class name based on property name
        nested_name = f"{name}{_to_pascal_case(prop_name)}"
        prop_result = render(prop_def.schema, nested_name)
        imports.update(prop_result.imports)

        # Collect nested class definitions
        if prop_result.code:
            nested_classes.append(prop_result.code)

        # Determine if property is required
        is_required = prop_def.required

        # Check if the property schema is a union containing null (nullable)
        is_nullable = _is_nullable_type(prop_def.schema)

        # Build the field definition
        type_expr = prop_result.type_expr

        if is_required:
            if is_nullable:
                # Required but nullable: field: Type | None (no default)
                if not type_expr.endswith(" | None"):
                    type_expr = f"{type_expr} | None"
                field_lines.append(f"    {prop_name}: {type_expr}")
            else:
                # Required and not nullable: field: Type
                field_lines.append(f"    {prop_name}: {type_expr}")
        else:
            # Optional: field: Type | None = None
            if not type_expr.endswith(" | None") and type_expr != "None":
                type_expr = f"{type_expr} | None"
            field_lines.append(f"    {prop_name}: {type_expr} = None")

    # Handle additionalProperties
    config_line = None
    if node.additional_properties is False:
        imports.add("from pydantic import ConfigDict")
        config_line = "    model_config = ConfigDict(extra='forbid')"
    elif node.additional_properties is True:
        imports.add("from pydantic import ConfigDict")
        config_line = "    model_config = ConfigDict(extra='allow')"
    # If additional_properties is a schema, we allow extra but don't constrain type
    # (Pydantic doesn't support typed extra fields directly in model_config)
    elif (
        node.additional_properties is not None
        and node.additional_properties is not False
    ):
        imports.add("from pydantic import ConfigDict")
        config_line = "    model_config = ConfigDict(extra='allow')"

    # Build the class definition
    class_lines: list[str] = [f"class {name}(BaseModel):"]

    # Add config if present
    if config_line:
        class_lines.append(config_line)

    # Add fields or pass if empty
    if field_lines:
        class_lines.extend(field_lines)
    elif not config_line:
        class_lines.append("    pass")

    # Combine nested classes and main class
    class_code = "\n".join(class_lines)
    if nested_classes:
        code = "\n\n\n".join(nested_classes) + "\n\n\n" + class_code
    else:
        code = class_code

    return RenderResult(code=code, type_expr=name, imports=imports)


def _to_pascal_case(name: str) -> str:
    """Convert a name to PascalCase for nested class names."""
    # Handle snake_case and camelCase
    parts = name.replace("-", "_").split("_")
    return "".join(part.capitalize() for part in parts if part)


def _is_nullable_type(node: SchemaNode) -> bool:
    """Check if a schema node represents a nullable type (union with null)."""
    if node.kind == "null":
        return True
    if node.kind == "union":
        # Check if any variant is null
        for variant in node.variants:
            if variant.kind == "null":
                return True
    return False


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
