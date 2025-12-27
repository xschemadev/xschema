package injector

import (
	"context"

	"github.com/xschema/cli/generator"
)

type InjectInput struct {
	Language string                     `json:"language"` // typescript, python, go
	Outputs  []generator.GenerateOutput `json:"outputs"`
	OutDir   string                     `json:"outDir"` // default .xschema
}

// Inject writes generated code to .xschema/ directory
func Inject(ctx context.Context, input InjectInput) error {
	// TODO: implement
	_ = ctx
	return nil
}
