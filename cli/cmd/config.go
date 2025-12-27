package cmd

import "time"

type Config struct {
	// Directories
	InputDir  string
	OutputDir string

	// HTTP/Retriever
	Concurrency int
	HTTPTimeout time.Duration
	Retries     int
	NoCache     bool

	// Filtering
	Include string // regex for files to include
	Exclude string // regex for files to exclude
	Adapter string // filter by adapter name

	// Output behavior
	DryRun  bool
	Force   bool
	Verbose bool
}
