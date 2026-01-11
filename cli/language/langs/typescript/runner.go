package typescript

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func detectRunner() (string, []string, error) {
	return detectRunnerInDir(".")
}

func detectRunnerInDir(dir string) (string, []string, error) {
	// Check package.json for packageManager field.
	if _, err := os.Stat(filepath.Join(dir, "package.json")); err == nil {
		content, err := os.ReadFile(filepath.Join(dir, "package.json"))
		if err == nil {
			pm := detectPackageManager(string(content))
			if pm != "" && commandExists(pm) {
				return packageManagerToRunner(pm)
			}
		}
	}

	// Check lockfiles.
	lockfileCmds := map[string][]string{
		"bun.lock":          {"bunx"},
		"bun.lockb":         {"bunx"},
		"pnpm-lock.yaml":    {"pnpm", "exec"},
		"yarn.lock":         {"yarn"},
		"package-lock.json": {"npx"},
	}

	for lockfile, cmd := range lockfileCmds {
		if _, err := os.Stat(filepath.Join(dir, lockfile)); err == nil {
			if commandExists(cmd[0]) {
				return cmd[0], cmd[1:], nil
			}
		}
	}

	// Fallback: try available commands.
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

func detectHarnessRunner(dir string) (string, []string, error) {
	cmd, _, err := detectRunnerInDir(dir)
	if err != nil {
		return "", nil, err
	}

	// Transform package runner to file runner:
	// bunx/bun → bun (bun can run .ts directly)
	// npx/pnpm/yarn → use tsx (tsx runs .ts files)
	switch cmd {
	case "bunx", "bun":
		return "bun", nil, nil
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

func commandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}
