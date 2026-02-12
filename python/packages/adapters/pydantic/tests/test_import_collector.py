"""Tests for ImportCollector."""

import pytest

from xschema_pydantic.import_collector import ImportCollector


def test_add_single_import():
    """Test adding a single import."""
    collector = ImportCollector()
    collector.add("from typing import Annotated")

    imports = collector.to_list()
    assert imports == ["from typing import Annotated"]


def test_add_multiple_imports_same_module():
    """Test adding multiple imports from the same module."""
    collector = ImportCollector()
    collector.add("from typing import Annotated")
    collector.add("from typing import Any")

    imports = collector.to_list()
    assert imports == ["from typing import Annotated, Any"]


def test_deduplication():
    """Test that duplicate imports are deduplicated."""
    collector = ImportCollector()
    collector.add("from typing import Annotated")
    collector.add("from typing import Annotated")
    collector.add("from typing import Annotated")

    imports = collector.to_list()
    assert imports == ["from typing import Annotated"]


def test_add_from_method():
    """Test add_from method for programmatic import addition."""
    collector = ImportCollector()
    collector.add_from("typing", "Annotated", "Any", "Union")

    imports = collector.to_list()
    assert imports == ["from typing import Annotated, Any, Union"]


def test_multiple_modules():
    """Test imports from multiple modules are handled correctly."""
    collector = ImportCollector()
    collector.add("from typing import Annotated")
    collector.add("from pydantic import BaseModel")
    collector.add("from datetime import date")

    imports = collector.to_list()
    # stdlib first (typing, datetime), then third-party (pydantic)
    assert "from typing import Annotated" in imports
    assert "from pydantic import BaseModel" in imports
    assert "from datetime import date" in imports


def test_grouping_order():
    """Test that imports are grouped correctly: stdlib, third-party, other."""
    collector = ImportCollector()
    collector.add("from pydantic import BaseModel")
    collector.add("from typing import Annotated")
    collector.add("from datetime import datetime")
    collector.add("from annotated_types import Ge")

    imports = collector.to_list()

    # Find indices
    typing_idx = next(i for i, imp in enumerate(imports) if "from typing import" in imp)
    datetime_idx = next(
        i for i, imp in enumerate(imports) if "from datetime import" in imp
    )
    pydantic_idx = next(
        i for i, imp in enumerate(imports) if "from pydantic import" in imp
    )
    annotated_idx = next(
        i for i, imp in enumerate(imports) if "from annotated_types import" in imp
    )

    # stdlib (typing, datetime) should come before third-party (pydantic, annotated_types)
    assert typing_idx < pydantic_idx
    assert datetime_idx < pydantic_idx
    assert typing_idx < annotated_idx
    assert datetime_idx < annotated_idx


def test_merge_collectors():
    """Test merging two ImportCollectors."""
    collector1 = ImportCollector()
    collector1.add("from typing import Annotated")
    collector1.add("from pydantic import BaseModel")

    collector2 = ImportCollector()
    collector2.add("from typing import Any")
    collector2.add("from datetime import datetime")

    collector1.merge(collector2)

    imports = collector1.to_list()
    assert "from typing import Annotated, Any" in imports
    assert "from pydantic import BaseModel" in imports
    assert "from datetime import datetime" in imports


def test_comma_separated_imports():
    """Test parsing comma-separated imports."""
    collector = ImportCollector()
    collector.add("from typing import Annotated, Any, Union")

    imports = collector.to_list()
    assert imports == ["from typing import Annotated, Any, Union"]


def test_sorted_names_within_module():
    """Test that import names are sorted alphabetically within each module."""
    collector = ImportCollector()
    collector.add("from typing import Union")
    collector.add("from typing import Annotated")
    collector.add("from typing import Any")

    imports = collector.to_list()
    assert imports == ["from typing import Annotated, Any, Union"]


def test_bare_import():
    """Test bare import statements (import module)."""
    collector = ImportCollector()
    collector.add("import re")

    imports = collector.to_list()
    assert imports == ["import re"]


def test_mixed_import_styles():
    """Test mixing bare imports and from imports."""
    collector = ImportCollector()
    collector.add("import re")
    collector.add("from typing import Annotated")

    imports = collector.to_list()
    assert "import re" in imports
    assert "from typing import Annotated" in imports


def test_complex_scenario():
    """Test a complex scenario with many imports."""
    collector = ImportCollector()

    # Add imports as they might appear in renderer code
    collector.add("from typing import Annotated")
    collector.add("from pydantic import BaseModel")
    collector.add("from typing import Any")
    collector.add("from pydantic import Field")
    collector.add("from datetime import datetime")
    collector.add("from uuid import UUID")
    collector.add("from annotated_types import Ge")
    collector.add("from annotated_types import Le")
    collector.add("from pydantic import EmailStr")

    imports = collector.to_list()

    # Check all imports present
    assert "from typing import Annotated, Any" in imports
    assert "from pydantic import BaseModel, EmailStr, Field" in imports
    assert "from datetime import datetime" in imports
    assert "from uuid import UUID" in imports
    assert "from annotated_types import Ge, Le" in imports

    # Check grouping order (stdlib before third-party)
    typing_idx = next(i for i, imp in enumerate(imports) if "from typing import" in imp)
    pydantic_idx = next(
        i for i, imp in enumerate(imports) if "from pydantic import" in imp
    )
    assert typing_idx < pydantic_idx


def test_to_set_method():
    """Test to_set method for backward compatibility."""
    collector = ImportCollector()
    collector.add("from typing import Annotated")
    collector.add("from pydantic import BaseModel")

    import_set = collector.to_set()
    assert isinstance(import_set, set)
    assert len(import_set) == 2


def test_no_duplicate_modules_in_output():
    """Test that each module only appears once in final output."""
    collector = ImportCollector()
    collector.add("from typing import Annotated")
    collector.add("from typing import Any")
    collector.add("from typing import Union")

    imports = collector.to_list()

    # Count how many times "from typing import" appears
    typing_count = sum(1 for imp in imports if imp.startswith("from typing import"))
    assert typing_count == 1
    assert imports == ["from typing import Annotated, Any, Union"]
