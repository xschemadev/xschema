"""XSchema CLI - Bring your JSON Schemas to life."""

import platform
import subprocess
import sys
from pathlib import Path

# Platform mapping: (system, machine) -> package name
PLATFORMS: dict[tuple[str, str], str] = {
    ("Darwin", "arm64"): "xschema_cli_darwin_arm64",
    ("Darwin", "x86_64"): "xschema_cli_darwin_x64",
    ("Linux", "aarch64"): "xschema_cli_linux_arm64",
    ("Linux", "x86_64"): "xschema_cli_linux_x64",
    ("Windows", "ARM64"): "xschema_cli_win32_arm64",
    ("Windows", "AMD64"): "xschema_cli_win32_x64",
}


def get_binary_path() -> Path:
    """Find the xschema binary for the current platform."""
    system = platform.system()
    machine = platform.machine()

    pkg_name = PLATFORMS.get((system, machine))
    if not pkg_name:
        raise RuntimeError(
            f"xschema: unsupported platform {system}-{machine}\n"
            "Please open an issue at https://github.com/xschemadev/xschema/issues"
        )

    # Try to import the platform-specific package
    try:
        pkg = __import__(pkg_name)
        pkg_dir = Path(pkg.__file__).parent
    except ImportError:
        raise RuntimeError(
            f"xschema: could not find binary for {system}-{machine}\n"
            f"Expected package: {pkg_name}\n\n"
            "Try reinstalling: pip install xschema-cli\n"
            "Or install with platform extra: pip install 'xschema-cli[darwin-arm64]'"
        )

    # Find binary in package directory
    bin_name = "xschema.exe" if system == "Windows" else "xschema"
    bin_path = pkg_dir / bin_name

    if not bin_path.exists():
        raise RuntimeError(
            f"xschema: binary not found at {bin_path}\n"
            f"Package {pkg_name} may be corrupted. Try reinstalling."
        )

    return bin_path


def main() -> None:
    """Entry point for the xschema CLI."""
    try:
        bin_path = get_binary_path()
    except RuntimeError as e:
        print(str(e), file=sys.stderr)
        sys.exit(1)

    # Execute the binary with all arguments
    try:
        result = subprocess.run(
            [str(bin_path)] + sys.argv[1:],
            stdin=sys.stdin,
            stdout=sys.stdout,
            stderr=sys.stderr,
        )
        sys.exit(result.returncode)
    except KeyboardInterrupt:
        sys.exit(130)
    except Exception as e:
        print(f"xschema: failed to execute binary: {e}", file=sys.stderr)
        sys.exit(1)


if __name__ == "__main__":
    main()
