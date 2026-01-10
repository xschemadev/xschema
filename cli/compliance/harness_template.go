package compliance

// HarnessTemplateData contains the data for rendering a harness template
type HarnessTemplateData struct {
	Imports        string               // merged import block (already formatted)
	Schemas        []HarnessSchemaEntry // schema entries
	TestDataString string               // JSON string of test data [{groupId, tests}, ...]
}

// HarnessSchemaEntry represents a single schema entry
type HarnessSchemaEntry struct {
	GroupID    string // unique identifier for this group
	Schema     string // generated schema code
	Type       string // generated type expression
	Validate   string // validation function
	IsTypeOnly bool   // true when Validate is empty
}
