package language

// TypeScript/TSX query for xschema.fromURL/fromFile calls
// - Matches string/template for name
// - Matches string/template for URL/file
// - Matches identifier for adapter
// - Rejects template interpolation via predicate
var tsQuery = `
(call_expression
  function: (member_expression
    object: (identifier) @_obj
    property: (property_identifier) @method)
  arguments: (arguments 
    . [(string) (template_string)] @name 
    . [(string) (template_string)] @source 
    . (identifier) @adapter .)
  (#eq? @_obj "xschema")
  (#not-match? @name "\\$\\{")
  (#not-match? @source "\\$\\{"))
`

// Python query for xschema.from_url/from_file calls
// - Matches string for source, identifier for adapter
// - Skips f-strings via predicate
// - Uses . anchor at end to ensure adapter is last arg
var pyQuery = `
(call
  function: (attribute
    object: (identifier) @_obj
    attribute: (identifier) @method)
  arguments: (argument_list
    (string) @name
    (string) @source
    (identifier) @adapter .)
  (#eq? @_obj "xschema")
  (#not-match? @name "^f[\"']")
  (#not-match? @source "^f[\"']"))
`

// TypeScript/TSX import query for adapter packages
// - Matches import { identifier } from "package" statements
var tsImportQuery = `
(import_statement
  source: (string) @package
  (import_specifier
    name: (identifier) @imported_name))
`

// Python import query for adapter packages
// - Matches from package import identifier statements
var pyImportQuery = `
(import_from_statement
  module_name: (dotted_name) @package
  (dotted_name) @imported_name)
`
