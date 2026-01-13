"""Import collection and deduplication for Pydantic code generation."""

from dataclasses import dataclass, field
from typing import Literal


@dataclass
class ImportCollector:
    """Tracks and deduplicates imports during code generation.

    Imports are grouped by module and formatted as 'from module import Name1, Name2'.
    Supports conditional imports (only added if feature is used).
    """

    _imports: dict[str, set[str]] = field(default_factory=dict)
    """Map from module name to set of imported names."""

    def add(self, import_statement: str) -> None:
        """Add an import statement.

        Accepts statements in format:
        - "from module import Name"
        - "from module import Name1, Name2"
        - "import module"

        Args:
            import_statement: Import statement to add
        """
        if import_statement.startswith("from "):
            # Parse "from module import Name1, Name2"
            parts = import_statement.split(" import ", 1)
            if len(parts) != 2:
                return  # Invalid format, skip
            module = parts[0].replace("from ", "").strip()
            names = parts[1].split(",")

            # Add each name to the module's import set
            if module not in self._imports:
                self._imports[module] = set()
            for name in names:
                self._imports[module].add(name.strip())
        elif import_statement.startswith("import "):
            # Handle "import module" - store as module with empty names
            module = import_statement.replace("import ", "").strip()
            if module not in self._imports:
                self._imports[module] = set()
            # Mark as bare import with special marker
            self._imports[module].add("__bare__")

    def add_from(self, module: str, *names: str) -> None:
        """Add imports from a module.

        Example:
            collector.add_from("typing", "Annotated", "Any")
            # Results in: from typing import Annotated, Any

        Args:
            module: Module to import from
            *names: Names to import from the module
        """
        if module not in self._imports:
            self._imports[module] = set()
        self._imports[module].update(names)

    def merge(self, other: "ImportCollector") -> None:
        """Merge imports from another collector.

        Args:
            other: Another ImportCollector to merge from
        """
        for module, names in other._imports.items():
            if module not in self._imports:
                self._imports[module] = set()
            self._imports[module].update(names)

    def to_list(self) -> list[str]:
        """Convert collected imports to sorted list of import statements.

        Groups imports by module category for organized output:
        1. Standard library (typing, datetime, uuid, re)
        2. Third-party (pydantic, annotated_types)
        3. Other modules

        Returns:
            Sorted list of import statements
        """
        # Group modules by category
        stdlib_modules = {
            "typing",
            "datetime",
            "uuid",
            "re",
        }
        thirdparty_modules = {"pydantic", "annotated_types"}

        stdlib_imports: list[str] = []
        thirdparty_imports: list[str] = []
        other_imports: list[str] = []

        # Sort modules for deterministic output
        for module in sorted(self._imports.keys()):
            names = self._imports[module]

            # Handle bare imports (import module)
            if "__bare__" in names:
                import_stmt = f"import {module}"
            else:
                # Sort names for consistent output
                sorted_names = sorted(names)
                import_stmt = f"from {module} import {', '.join(sorted_names)}"

            # Categorize import
            if module in stdlib_modules or module.split(".")[0] in stdlib_modules:
                stdlib_imports.append(import_stmt)
            elif (
                module in thirdparty_modules
                or module.split(".")[0] in thirdparty_modules
            ):
                thirdparty_imports.append(import_stmt)
            else:
                other_imports.append(import_stmt)

        # Combine all imports in order
        return stdlib_imports + thirdparty_imports + other_imports

    def to_set(self) -> set[str]:
        """Convert collected imports to set of import statements.

        Used for backward compatibility with existing code that expects set[str].

        Returns:
            Set of import statements
        """
        return set(self.to_list())
