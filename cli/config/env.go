// Package config handles configuration loading for the CLI.
package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/xschemadev/xschema/ui"
)

// LoadEnvFile loads environment variables from a .env file.
// If envFile is empty, attempts to load .env from projectRoot.
// Returns nil if no env file is specified and default doesn't exist.
func LoadEnvFile(envFile, projectRoot string) error {
	if envFile != "" {
		ui.Verbosef("loading env file: %s", envFile)
		if err := godotenv.Load(envFile); err != nil {
			return fmt.Errorf("failed to load env file %s: %w", envFile, err)
		}
		return nil
	}

	// Try to load .env from project root if it exists
	defaultEnvPath := filepath.Join(projectRoot, ".env")
	if _, err := os.Stat(defaultEnvPath); err == nil {
		ui.Verbosef("loading default .env file: %s", defaultEnvPath)
		_ = godotenv.Load(defaultEnvPath) // ignore error - default env is optional
	}
	return nil
}
