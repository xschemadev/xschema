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

	// Check for uv.lock (indicates uv project)
	if _, err := os.Stat(filepath.Join(dir, "uv.lock")); err == nil {
		if commandExists("uv") {
			return "uv", []string{"tool", "run"}, nil
		}
	}

	// Check for pyproject.toml
	pyprojectPath := filepath.Join(dir, "pyproject.toml")
	if _, err := os.Stat(pyprojectPath); err == nil {
		content, err := os.ReadFile(pyprojectPath)
		if err == nil {
			contentStr := string(content)

			// Check for [tool.uv] section
			if strings.Contains(contentStr, "[tool.uv]") {
				if commandExists("uv") {
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

	// Transform tool runner to pytest runner
	switch runner {
	case "uv":
		return "uv", []string{"run", "pytest"}, nil
	case "poetry":
		return "poetry", []string{"run", "pytest"}, nil
	case "pipenv":
		return "pipenv", []string{"run", "pytest"}, nil
	case "pipx":
		return "python", []string{"-m", "pytest"}, nil
	default:
		return "python", []string{"-m", "pytest"}, nil
	}
}

func commandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}
