package python

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
	// Priority order: uv > poetry > pipenv > pip

	pyprojectPath := filepath.Join(dir, "pyproject.toml")

	// Check for uv.lock (indicates standalone uv project, not workspace member)
	hasLocalLock := false
	if _, err := os.Stat(filepath.Join(dir, "uv.lock")); err == nil {
		hasLocalLock = true
		if commandExists("uv") {
			return "uv", []string{"tool", "run"}, nil
		}
	}

	// Check for pyproject.toml
	if _, err := os.Stat(pyprojectPath); err == nil {
		content, err := os.ReadFile(pyprojectPath)
		if err == nil {
			contentStr := string(content)

			// Check for uv workspace member (has workspace sources but no local lock)
			isWorkspaceMember := !hasLocalLock && strings.Contains(contentStr, "workspace = true")

			// Check for [tool.uv] section
			if strings.Contains(contentStr, "[tool.uv]") || isWorkspaceMember {
				if commandExists("uv") {
					// If workspace member, use "uv run" instead of "uv tool run"
					if isWorkspaceMember {
						return "uv", []string{"run"}, nil
					}
					return "uv", []string{"tool", "run"}, nil
				}
			}

			// Check for [tool.poetry] section
			if strings.Contains(contentStr, "[tool.poetry]") {
				if commandExists("poetry") {
					return "poetry", []string{"run"}, nil
				}
			}
		}
	}

	// Check for Pipfile (pipenv)
	if _, err := os.Stat(filepath.Join(dir, "Pipfile")); err == nil {
		if commandExists("pipenv") {
			return "pipenv", []string{"run"}, nil
		}
	}

	// Fallback: try available commands in order of preference
	if commandExists("uv") {
		return "uv", []string{"tool", "run"}, nil
	}
	if commandExists("poetry") {
		return "poetry", []string{"run"}, nil
	}
	if commandExists("pipenv") {
		return "pipenv", []string{"run"}, nil
	}

	// Final fallback: use pipx if available, otherwise python -m
	if commandExists("pipx") {
		return "pipx", []string{"run"}, nil
	}

	return "python", []string{"-m"}, nil
}

func detectHarnessRunner(dir string) (string, []string, error) {
	runner, _, err := detectRunnerInDir(dir)
	if err != nil {
		return "", nil, err
	}

	// Transform tool runner to python runner for executing harness scripts
	switch runner {
	case "uv":
		return "uv", []string{"run", "python"}, nil
	case "poetry":
		return "poetry", []string{"run", "python"}, nil
	case "pipenv":
		return "pipenv", []string{"run", "python"}, nil
	default:
		return "python", nil, nil
	}
}

func commandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}
