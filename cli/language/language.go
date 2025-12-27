package language

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/python"
	"github.com/smacker/go-tree-sitter/typescript/typescript"
)

var extMap map[string]*Language

type SourceType string

const (
	SourceURL  SourceType = "url"
	SourceFile SourceType = "file"
)

type Language struct {
	Name          string
	Extensions    []string
	GetSitterLang func() *sitter.Language
	Query         string
	MethodMapping map[string]SourceType // maps method names to source type
	DetectRunner  func() (cmd string, args []string, err error)
}

var Languages = []Language{
	{
		Name:          "typescript",
		Extensions:    []string{".ts", ".tsx", ".js", ".jsx"},
		GetSitterLang: typescript.GetLanguage,
		Query:         tsQuery,
		MethodMapping: map[string]SourceType{
			"fromURL":  SourceURL,
			"fromFile": SourceFile,
		},
		DetectRunner: detectTSRunner,
	},
	{
		Name:          "python",
		Extensions:    []string{".py"},
		GetSitterLang: python.GetLanguage,
		Query:         pyQuery,
		MethodMapping: map[string]SourceType{
			"from_url":  SourceURL,
			"from_file": SourceFile,
		},
		DetectRunner: detectPythonRunner,
	},
}

func detectTSRunner() (string, []string, error) {
	checkCmd := func(cmd string) bool {
		_, err := exec.LookPath(cmd)
		return err == nil
	}

	if _, err := os.Stat(filepath.Join(".", "package.json")); err == nil {
		content, err := os.ReadFile(filepath.Join(".", "package.json"))
		if err == nil {
			pm := detectPackageManager(string(content))
			if pm != "" && checkCmd(pm) {
				switch pm {
				case "bun":
					return "bunx", nil, nil
				case "pnpm":
					return "pnpm", []string{"exec"}, nil
				case "yarn":
					return "yarn", nil, nil
				case "npm":
					return "npx", nil, nil
				}
			}
		}
	}

	lockfileCmds := map[string][]string{
		"bun.lock":          {"bunx"},
		"bun.lockb":         {"bunx"},
		"pnpm-lock.yaml":    {"pnpm", "exec"},
		"yarn.lock":         {"yarn"},
		"package-lock.json": {"npx"},
	}

	for lf, cmd := range lockfileCmds {
		if _, err := os.Stat(filepath.Join(".", lf)); err == nil {
			if checkCmd(cmd[0]) {
				return cmd[0], cmd[1:], nil
			}
		}
	}

	for _, cmd := range []string{"bunx", "pnpm", "yarn", "npx"} {
		if checkCmd(cmd) {
			if cmd == "pnpm" {
				return cmd, []string{"exec"}, nil
			}
			return cmd, nil, nil
		}
	}

	return "npx", nil, nil
}

func detectPackageManager(content string) string {
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

func detectPythonRunner() (string, []string, error) {
	checkCmd := func(cmd string) bool {
		_, err := exec.LookPath(cmd)
		return err == nil
	}

	lockfileCmds := map[string][]string{
		"uv.lock":     {"uv", "run"},
		"poetry.lock": {"poetry", "run"},
		"Pipfile":     {"pipenv", "run"},
	}

	for lf, cmd := range lockfileCmds {
		if _, err := os.Stat(filepath.Join(".", lf)); err == nil {
			if checkCmd(cmd[0]) {
				return cmd[0], cmd[1:], nil
			}
		}
	}

	if _, err := os.Stat(filepath.Join(".", "pyproject.toml")); err == nil {
		content, err := os.ReadFile(filepath.Join(".", "pyproject.toml"))
		if err == nil {
			buildSystem := detectBuildSystem(string(content))
			if buildSystem != "" && checkCmd(buildSystem) {
				return buildSystem, []string{"run"}, nil
			}
		}
	}

	return "python", []string{"-m"}, nil
}

func detectBuildSystem(content string) string {
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

// ExtensionGlobs returns glob patterns for all supported extensions
func ExtensionGlobs() []string {
	var globs []string
	for _, lang := range Languages {
		for _, ext := range lang.Extensions {
			globs = append(globs, "**/*"+ext)
		}
	}
	return globs
}

func init() {
	extMap = make(map[string]*Language)
	for i := range Languages {
		for _, ext := range Languages[i].Extensions {
			extMap[ext] = &Languages[i]
		}
	}
}

// ByExtension returns the language config for a file extension
func ByExtension(ext string) *Language {
	return extMap[ext]
}

// ByName returns the language config by name
func ByName(name string) *Language {
	for i, lang := range Languages {
		if lang.Name == name {
			return &Languages[i]
		}
	}
	return nil
}
