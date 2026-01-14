"""Generate script wrapper for xschema CLI."""

import subprocess
import sys
from pathlib import Path


def main():
    """Run xschema generate from the CLI directory."""
    cli_dir = Path(__file__).parent.parent.parent / "cli"
    example_dir = Path(__file__).parent

    result = subprocess.run(
        ["go", "run", ".", "generate", "--project", str(example_dir)],
        cwd=cli_dir,
        capture_output=False,
    )

    sys.exit(result.returncode)


if __name__ == "__main__":
    main()
