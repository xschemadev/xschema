package cmd

import (
	"github.com/xschemadev/xschema/retriever"
	"testing"
	"time"
)

func Test_printSummary(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		schemas       []retriever.RetrievedSchema
		outDir        string
		generatedFile string
		duration      time.Duration
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printSummary(tt.schemas, tt.generatedFile, tt.duration)
		})
	}
}
