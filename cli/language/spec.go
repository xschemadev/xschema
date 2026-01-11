package language

import (
	"context"
	"fmt"
	"strings"
)

// GeneratedFile represents one file emitted by a language.
// Path is always relative to the output root.
// Contents are written verbatim.
type GeneratedFile struct {
	Path     string `json:"path"`
	Contents string `json:"contents"`
}

// CommandSpec is a fully-resolved command invocation.
// It is intentionally minimal so different languages can map to OS processes.
type CommandSpec struct {
	Cmd  string   `json:"cmd"`
	Args []string `json:"args,omitempty"`
	Dir  string   `json:"dir,omitempty"`
	Env  []string `json:"env,omitempty"`
}

func (c CommandSpec) Validate() error {
	if strings.TrimSpace(c.Cmd) == "" {
		return fmt.Errorf("command cmd is required")
	}
	return nil
}

// AdapterCommandInput is the input needed to resolve an adapter invocation.
// ProjectRoot is the directory where the command should be run.
type AdapterCommandInput struct {
	ProjectRoot string
	AdapterRef  string
}

// AdapterInvoker resolves adapter refs into a runnable command.
// The generator owns stdin/stdout wiring; the invoker only describes how to run.
type AdapterInvoker interface {
	BuildAdapterCommand(ctx context.Context, input AdapterCommandInput) (CommandSpec, error)
}

// EmitInput is the fully-prepared data for emitting generated files.
// A language implementation can map this to one or many GeneratedFile outputs.
type EmitInput struct {
	OutDir   string
	Imports  string
	Schemas  []SchemaEntry
	Header   string
	Footer   string
	Filename string
}

// OutputEmitter produces one or more files from generated schemas.
type OutputEmitter interface {
	Emit(ctx context.Context, input EmitInput) ([]GeneratedFile, error)
}
