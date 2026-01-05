package language

import (
	"os"
	"path/filepath"
	"strings"
)

var python = Language{
	Name:                 "python",
	Extensions:           []string{".py"},
	SchemaURL:            XSchemaBaseURL + "py.jsonc",
	SchemaExt:            "py.jsonc",
	AdapterBinPrefix:     "xschema-",
	DetectRunner:         detectPythonRunner,
	BuildSchemasImport:   buildPySchemasImport,
	ImportPattern:        `(?m)^(?:import\s+|from\s+).*$`,
	InjectSchemasKey:     injectSchemasKeyBrace,
	ClientFactoryPattern: `create_xschema_client\s*\(\s*(\{[^}]*\})\s*\)`,
	OutputFile:           "__init__.py",
	Template:             PyTemplate,
	MergeImports:         MergePyImports,
	BuildFooter:          BuildPythonFooter,
	BuildVarName:         buildVarNameUnderscore,
	IgnoreDirs:           []string{"__pycache__", ".venv", "venv"},
	DetectHarnessRunner:  detectPyHarnessRunner,
	GetPackageName:       getPyPackageName,
	AdapterCLIPath:       getPyAdapterCLIPath,
}

func detectPythonRunner() (string, []string, error) {
	return detectPythonRunnerInDir(".")
}

func detectPythonRunnerInDir(dir string) (string, []string, error) {
	// Check lockfiles
	lockfileCmds := map[string][]string{
		"uv.lock":     {"uv", "run"},
		"poetry.lock": {"poetry", "run"},
		"Pipfile":     {"pipenv", "run"},
	}

	for lf, cmd := range lockfileCmds {
		if _, err := os.Stat(filepath.Join(dir, lf)); err == nil {
			if commandExists(cmd[0]) {
				return cmd[0], cmd[1:], nil
			}
		}
	}

	// Check pyproject.toml for build system
	if _, err := os.Stat(filepath.Join(dir, "pyproject.toml")); err == nil {
		content, err := os.ReadFile(filepath.Join(dir, "pyproject.toml"))
		if err == nil {
			buildSystem := detectPyBuildSystem(string(content))
			if buildSystem != "" && commandExists(buildSystem) {
				return buildSystem, []string{"run"}, nil
			}
		}
	}

	return "python", []string{"-m"}, nil
}

func detectPyBuildSystem(content string) string {
	lines := strings.SplitSeq(content, "\n")
	for line := range lines {
		line = strings.TrimSpace(line)
		if strings.Contains(line, "build-backend") && strings.Contains(line, "uv") {
			return "uv"
		}
		if strings.Contains(line, "requires") && strings.Contains(line, "poetry-core") {
			return "poetry"
		}
		if strings.Contains(line, "requires") && strings.Contains(line, "flit") {
			return "flit"
		}
		if strings.Contains(line, "requires") && strings.Contains(line, "setuptools") {
			return ""
		}
	}
	return ""
}

func detectPyHarnessRunner(dir string) (string, []string, error) {
	cmd, args, err := detectPythonRunnerInDir(dir)
	if err != nil {
		return "", nil, err
	}

	// Strip -m flag - only used for module execution, not direct files
	if len(args) == 1 && args[0] == "-m" {
		return cmd, nil, nil
	}
	return cmd, args, nil
}

func buildPySchemasImport(importPath string) string {
	// Convert path to module notation: ./.xschema/xschema -> .xschema.xschema
	module := strings.ReplaceAll(importPath, "/", ".")
	module = strings.TrimPrefix(module, ".")
	return "from " + module + " import schemas"
}

func getPyPackageName(dir string) string {
	pkgPath := filepath.Join(dir, "pyproject.toml")
	data, err := os.ReadFile(pkgPath)
	if err != nil {
		return filepath.Base(dir)
	}

	// Simple TOML parsing - look for name = "..." under [project] or [tool.poetry]
	content := string(data)
	lines := strings.Split(content, "\n")
	inProject := false

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "[project]" || line == "[tool.poetry]" {
			inProject = true
			continue
		}
		if strings.HasPrefix(line, "[") {
			inProject = false
			continue
		}
		if inProject && strings.HasPrefix(line, "name") {
			// Extract value: name = "foo"
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				val := strings.TrimSpace(parts[1])
				val = strings.Trim(val, `"'`)
				if val != "" {
					return val
				}
			}
		}
	}

	return filepath.Base(dir)
}

func getPyAdapterCLIPath(adapterPath string) string {
	return filepath.Join(adapterPath, "src", "cli.py")
}
