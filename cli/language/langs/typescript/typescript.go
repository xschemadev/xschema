package typescript

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/xschemadev/xschema/language"
)

const adapterBinPrefix = "xschema-"

func init() {
	if err := language.Register(Language()); err != nil {
		panic(fmt.Sprintf("language typescript register failed: %v", err))
	}
}

func Language() language.Language {
	return language.Language{
		Name:                 "typescript",
		Extensions:           []string{".ts", ".tsx", ".js", ".jsx"},
		SchemaURL:            language.XSchemaBaseURL + "typescript.jsonc",
		SchemaExt:            "typescript.jsonc",
		AdapterBinPrefix:     adapterBinPrefix,
		DetectRunner:         detectRunner,
		AdapterInvoker:       adapterInvoker{},
		BuildSchemasImport:   buildSchemasImport,
		ImportPattern:        `(?m)^import\s+.*$`,
		InjectSchemasKey:     injectSchemasKeyBrace,
		ClientFactoryPattern: `createXSchemaClient\s*\(\s*(\{[^}]*\})\s*\)`,
		OutputFile:           "xschema.gen.ts",
		Template:             outputTemplate,
		MergeImports:         MergeImports,
		BuildVarName:         buildVarNameUnderscore,
		IgnoreDirs:           []string{"node_modules", "dist", "build"},
		DetectHarnessRunner:  detectHarnessRunner,
		GetPackageName:       getPackageName,
		AdapterCLIPath:       getAdapterCLIPath,
		HarnessExtension:     ".ts",
		HarnessTemplate:      harnessTemplate,
	}
}

type adapterInvoker struct{}

func (adapterInvoker) BuildAdapterCommand(ctx context.Context, input language.AdapterCommandInput) (language.CommandSpec, error) {
	_ = ctx

	projectRoot := strings.TrimSpace(input.ProjectRoot)
	if projectRoot == "" {
		return language.CommandSpec{}, fmt.Errorf("project root is required")
	}

	adapterRef := strings.TrimSpace(input.AdapterRef)
	if adapterRef == "" {
		return language.CommandSpec{}, fmt.Errorf("adapter ref is required")
	}

	// Migration help for legacy adapter names like "zod".
	if !strings.HasPrefix(adapterRef, "@") && !strings.Contains(adapterRef, "/") {
		return language.CommandSpec{}, fmt.Errorf(
			"invalid typescript adapter ref %q: expected scoped npm package ref like %q (migration: change adapter to %q)",
			adapterRef,
			"@xschemadev/zod",
			"@xschemadev/"+adapterRef,
		)
	}

	if !strings.HasPrefix(adapterRef, "@xschemadev/") {
		return language.CommandSpec{}, fmt.Errorf(
			"invalid typescript adapter ref %q: expected %q scope (example: %q)",
			adapterRef,
			"@xschemadev",
			"@xschemadev/zod",
		)
	}

	parts := strings.Split(adapterRef, "/")
	if len(parts) != 2 || strings.TrimSpace(parts[1]) == "" {
		return language.CommandSpec{}, fmt.Errorf("invalid typescript adapter ref %q: expected format %q", adapterRef, "@xschemadev/<adapter>")
	}

	pkgName := parts[1]
	binName := adapterBinPrefix + pkgName

	runner, runnerArgs, err := detectRunnerInDir(projectRoot)
	if err != nil {
		return language.CommandSpec{}, fmt.Errorf("failed to detect typescript runner: %w", err)
	}

	return language.CommandSpec{
		Cmd:  runner,
		Args: append(runnerArgs, binName),
		Dir:  projectRoot,
	}, nil
}

func buildSchemasImport(importPath string) string {
	return `import { schemas } from "` + importPath + `";`
}

func getPackageName(dir string) string {
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

func getAdapterCLIPath(adapterPath string) string {
	return filepath.Join(adapterPath, "dist", "cli.js")
}

func buildVarNameUnderscore(namespace, id string) string {
	return namespace + "_" + id
}

func injectSchemasKeyBrace(configContent string) string {
	openIdx := strings.Index(configContent, "{")
	if openIdx == -1 {
		return configContent
	}

	if len(configContent) < openIdx+2 {
		return "{ schemas }"
	}

	inner := configContent[openIdx+1 : len(configContent)-1]
	innerTrimmed := strings.TrimSpace(inner)

	if strings.HasPrefix(innerTrimmed, "schemas") && (len(innerTrimmed) == 7 || strings.HasPrefix(innerTrimmed[7:], ",") || strings.HasPrefix(innerTrimmed[7:], "}")) {
		return configContent
	}

	if strings.HasPrefix(innerTrimmed, "schemas:") {
		return configContent
	}

	if strings.HasPrefix(innerTrimmed, `"schemas":`) || strings.HasPrefix(innerTrimmed, `'schemas':`) {
		return configContent
	}

	if innerTrimmed == "" {
		return "{ schemas }"
	}
	return "{ schemas, " + inner + " }"
}
