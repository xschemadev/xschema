package compliance

// HarnessTemplateData contains the data for rendering a harness template
type HarnessTemplateData struct {
	Schema          string // generated schema code
	Type            string // generated type expression (type-only adapters)
	Imports         string // merged import block (already formatted)
	Validate        string // validation function (empty = type-only)
	TestCasesString string // JSON string of test cases (already escaped for JS string literal)
	IsTypeOnly      bool   // true when Validate is empty
}
