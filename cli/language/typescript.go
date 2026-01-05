package language

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

var typescript = Language{
	Name:                 "typescript",
	Extensions:           []string{".ts", ".tsx", ".js", ".jsx"},
	SchemaURL:            XSchemaBaseURL + "ts.jsonc",
	SchemaExt:            "ts.jsonc",
	AdapterBinPrefix:     "xschema-",
	AdaptersPath:         "typescript/packages/adapters",
	DetectRunner:         detectTSRunner,
	BuildSchemasImport:   buildTSSchemasImport,
	ImportPattern:        `(?m)^import\s+.*$`,
	InjectSchemasKey:     injectSchemasKeyBrace,
	ClientFactoryPattern: `createXSchemaClient\s*\(\s*(\{[^}]*\})\s*\)`,
	OutputFile:           "xschema.gen.ts",
	Template:             TSTemplate,
	MergeImports:         MergeTSImports,
	BuildVarName:         buildVarNameUnderscore,
	IgnoreDirs:           []string{"node_modules", "dist", "build"},
	DetectHarnessRunner:  detectTSHarnessRunner,
	GetPackageName:       getTSPackageName,
	AdapterCLIPath:       getTSAdapterCLIPath,
}

func detectTSRunner() (string, []string, error) {
	return detectTSRunnerInDir(".")
}

func detectTSRunnerInDir(dir string) (string, []string, error) {
	// Check package.json for packageManager field
	if _, err := os.Stat(filepath.Join(dir, "package.json")); err == nil {
		content, err := os.ReadFile(filepath.Join(dir, "package.json"))
		if err == nil {
			pm := detectTSPackageManager(string(content))
			if pm != "" && commandExists(pm) {
				return packageManagerToRunner(pm)
			}
		}
	}

	// Check lockfiles
	lockfileCmds := map[string][]string{
		"bun.lock":          {"bunx"},
		"bun.lockb":         {"bunx"},
		"pnpm-lock.yaml":    {"pnpm", "exec"},
		"yarn.lock":         {"yarn"},
		"package-lock.json": {"npx"},
	}

	for lf, cmd := range lockfileCmds {
		if _, err := os.Stat(filepath.Join(dir, lf)); err == nil {
			if commandExists(cmd[0]) {
				return cmd[0], cmd[1:], nil
			}
		}
	}

	// Fallback: try available commands
	for _, cmd := range []string{"bunx", "pnpm", "yarn", "npx"} {
		if commandExists(cmd) {
			if cmd == "pnpm" {
				return cmd, []string{"exec"}, nil
			}
			return cmd, nil, nil
		}
	}

	return "npx", nil, nil
}

func packageManagerToRunner(pm string) (string, []string, error) {
	switch pm {
	case "bun":
		return "bunx", nil, nil
	case "pnpm":
		return "pnpm", []string{"exec"}, nil
	case "yarn":
		return "yarn", nil, nil
	case "npm":
		return "npx", nil, nil
	default:
		return "npx", nil, nil
	}
}

func detectTSPackageManager(content string) string {
	lines := strings.SplitSeq(content, "\n")
	for line := range lines {
		line = strings.TrimSpace(line)
		if strings.Contains(line, `"packageManager"`) {
			if strings.Contains(line, `"bun@`) {
				return "bun"
			}
			if strings.Contains(line, `"pnpm@`) {
				return "pnpm"
			}
			if strings.Contains(line, `"yarn@`) {
				return "yarn"
			}
			if strings.Contains(line, `"npm@`) {
				return "npm"
			}
		}
	}
	return ""
}

func detectTSHarnessRunner(dir string) (string, []string, error) {
	cmd, _, err := detectTSRunnerInDir(dir)
	if err != nil {
		return "", nil, err
	}

	// Transform package runner to file runner:
	// bunx/bun → bun run (bun can run .ts directly)
	// npx/pnpm/yarn → use tsx (tsx runs .ts files)
	switch cmd {
	case "bunx", "bun":
		return "bun", []string{"run"}, nil
	case "npx":
		return "npx", []string{"tsx"}, nil
	case "pnpm":
		return "pnpm", []string{"dlx", "tsx"}, nil
	case "yarn":
		return "yarn", []string{"dlx", "tsx"}, nil
	default:
		return "npx", []string{"tsx"}, nil
	}
}

func buildTSSchemasImport(importPath string) string {
	return `import { schemas } from "` + importPath + `";`
}

func getTSPackageName(dir string) string {
	pkgPath := filepath.Join(dir, "package.json")
	data, err := os.ReadFile(pkgPath)
	if err != nil {
		return filepath.Base(dir)
	}

	var pkg struct {
		Name string `json:"name"`
	}
	if err := json.Unmarshal(data, &pkg); err != nil {
		return filepath.Base(dir)
	}

	if pkg.Name != "" {
		return pkg.Name
	}
	return filepath.Base(dir)
}

func getTSAdapterCLIPath(adapterPath string) string {
	return filepath.Join(adapterPath, "dist", "cli.js")
}
