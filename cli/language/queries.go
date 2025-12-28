package language

// TypeScript/TSX query for xschema.fromURL/fromFile calls
// - Captures @obj to filter by client name in Go code
// - Matches string/template for name
// - Matches string/template for URL/file
// - Matches identifier for adapter
// - Rejects template interpolation via predicate
var tsQuery = `
(call_expression
  function: (member_expression
    object: (identifier) @obj
    property: (property_identifier) @method)
  arguments: (arguments 
    . [(string) (template_string)] @name 
    . [(string) (template_string)] @source 
    . (identifier) @adapter .)
  (#not-match? @name "\\$\\{")
  (#not-match? @source "\\$\\{"))
`

// TypeScript query for createXSchemaClient calls
// Captures the variable name assigned to the client
var tsClientQuery = `
(lexical_declaration
  (variable_declarator
    name: (identifier) @client_name
    value: (call_expression
      function: (identifier) @_fn
      (#eq? @_fn "createXSchemaClient"))))
`

// TypeScript query for config object in createXSchemaClient call
// Extracts key-value pairs from the first argument (config object)
var tsConfigQuery = `
(call_expression
  function: (identifier) @_fn
  arguments: (arguments
    (object
      (pair
        key: (property_identifier) @config_key
        value: (_) @config_value)))
  (#eq? @_fn "createXSchemaClient"))
`

// TypeScript query to find config object for injection
// Captures the full config object (detection of existing schemas done in Go code)
var tsClientCallQuery = `
(call_expression
  function: (identifier) @_fn
  arguments: (arguments
    (object) @config)
  (#eq? @_fn "createXSchemaClient"))
`

// Python query for xschema.from_url/from_file calls
// - Captures @obj to filter by client name in Go code
// - Matches string for source, identifier for adapter
// - Skips f-strings via predicate
// - Uses . anchor at end to ensure adapter is last arg
var pyQuery = `
(call
  function: (attribute
    object: (identifier) @obj
    attribute: (identifier) @method)
  arguments: (argument_list
    (string) @name
    (string) @source
    (identifier) @adapter .)
  (#not-match? @name "^f[\"']")
  (#not-match? @source "^f[\"']"))
`

// Python query for create_xschema_client calls
// Captures the variable name assigned to the client
var pyClientQuery = `
(assignment
  left: (identifier) @client_name
  right: (call
    function: (identifier) @_fn
    (#eq? @_fn "create_xschema_client")))
`

// Python query for config dict in create_xschema_client call
// Extracts key-value pairs from the first argument (config dict)
var pyConfigQuery = `
(call
  function: (identifier) @_fn
  arguments: (argument_list
    (dictionary
      (pair
        key: (string) @config_key
        value: (_) @config_value)))
  (#eq? @_fn "create_xschema_client"))
`

// Python query to find config dict for injection
var pyClientCallQuery = `
(call
  function: (identifier) @_fn
  arguments: (argument_list
    (dictionary) @config
    (pair
      key: (string) @_key
      (#match? @_key "schemas"))? @schemas_key)
  (#eq? @_fn "create_xschema_client"))
`

// TypeScript/TSX import query for adapter packages
// - Matches import { identifier } from "package" statements
var tsImportQuery = `
(import_statement
  (import_clause
    (named_imports
      (import_specifier
        name: (identifier) @imported_name)))
  source: (string) @package)
`

// Python import query for adapter packages
// - Matches from package import identifier statements
var pyImportQuery = `
(import_from_statement
  module_name: (dotted_name) @package
  (dotted_name) @imported_name)
`
