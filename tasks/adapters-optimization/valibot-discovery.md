# Valibot Adapter Discovery

## Baseline Per Draft

| Draft | Passed | Failed | Coverage |
| ----- | ------ | ------ | -------- |
| draft2020-12 | 782 | 138 | 85.0% |
| draft2019-09 | 769 | 146 | 84.0% |
| draft7 | 749 | 95 | 88.7% |
| draft6 | 673 | 95 | 87.6% |
| draft4 | 494 | 74 | 87.0% |
| draft3 | 365 | 42 | 89.7% |

## Failing Tests List

### draft2020-12 (138 failures)

| Keyword | Group | Test |
| ------- | ----- | ---- |
| additionalProperties | additionalProperties being false does not allow other properties | ignores arrays |
| additionalProperties | additionalProperties being false does not allow other properties | ignores other non-objects |
| additionalProperties | additionalProperties being false does not allow other properties | ignores strings |
| additionalProperties | additionalProperties being false does not allow other properties | patternProperties are not additional properties |
| additionalProperties | additionalProperties does not look in applicators | properties defined in allOf are not examined |
| contains | contains keyword validation | not array is valid |
| contains | contains keyword with boolean schema false | non-arrays are valid |
| defs | validate definition against metaschema | invalid definition schema |
| defs | validate definition against metaschema | valid definition schema |
| dependentRequired | empty dependents | non-object is valid |
| dependentRequired | single dependency | ignores other non-objects |
| dependentRequired | single dependency | ignores strings |
| dependentSchemas | dependencies with escaped characters | quoted quote |
| dependentSchemas | dependencies with escaped characters | quoted quote invalid under dependent schema |
| dependentSchemas | single dependency | ignores other non-objects |
| dependentSchemas | single dependency | ignores strings |
| dynamicRef | $dynamicRef points to a boolean schema | follow $dynamicRef to a false schema |
| dynamicRef | $dynamicRef points to a boolean schema | follow $dynamicRef to a true schema |
| dynamicRef | $dynamicRef skips over intermediate resources - direct reference | integer property passes |
| dynamicRef | $dynamicRef skips over intermediate resources - direct reference | string property fails |
| dynamicRef | $ref and $dynamicAnchor are independent of order - $defs first | correct extended schema |
| dynamicRef | $ref and $dynamicAnchor are independent of order - $defs first | incorrect extended schema |
| dynamicRef | $ref and $dynamicAnchor are independent of order - $defs first | incorrect parent schema |
| dynamicRef | $ref and $dynamicAnchor are independent of order - $ref first | correct extended schema |
| dynamicRef | $ref and $dynamicAnchor are independent of order - $ref first | incorrect extended schema |
| dynamicRef | $ref and $dynamicAnchor are independent of order - $ref first | incorrect parent schema |
| dynamicRef | $ref to $dynamicRef finds detached $dynamicAnchor | non-number is invalid |
| dynamicRef | $ref to $dynamicRef finds detached $dynamicAnchor | number is valid |
| dynamicRef | A $dynamicRef that initially resolves to a schema with a matching $dynamicAnchor... | The recursive part is not valid against the root |
| dynamicRef | A $dynamicRef that initially resolves to a schema with a matching $dynamicAnchor... | The recursive part is valid against the root |
| dynamicRef | A $dynamicRef that initially resolves to a schema without a matching $dynamicAnchor... | The recursive part doesn't need to validate against the root |
| dynamicRef | A $dynamicRef to a $dynamicAnchor in the same schema resource... | An array containing non-strings is invalid |
| dynamicRef | A $dynamicRef to a $dynamicAnchor in the same schema resource... | An array of strings is valid |
| dynamicRef | A $dynamicRef to an $anchor in the same schema resource... | An array containing non-strings is invalid |
| dynamicRef | A $dynamicRef to an $anchor in the same schema resource... | An array of strings is valid |
| dynamicRef | A $ref to a $dynamicAnchor in the same schema resource... | An array containing non-strings is invalid |
| dynamicRef | A $ref to a $dynamicAnchor in the same schema resource... | An array of strings is valid |
| dynamicRef | after leaving a dynamic scope, it is not used by a $dynamicRef | /then/$defs/thingy is the final stop for the $dynamicRef |
| dynamicRef | after leaving a dynamic scope, it is not used by a $dynamicRef | first_scope is not in dynamic scope for the $dynamicRef |
| dynamicRef | after leaving a dynamic scope, it is not used by a $dynamicRef | string matches /$defs/thingy, but the $dynamicRef does not stop here |
| dynamicRef | multiple dynamic paths to the $dynamicRef keyword | number list with number values |
| dynamicRef | multiple dynamic paths to the $dynamicRef keyword | number list with string values |
| dynamicRef | multiple dynamic paths to the $dynamicRef keyword | string list with number values |
| dynamicRef | multiple dynamic paths to the $dynamicRef keyword | string list with string values |
| dynamicRef | strict-tree schema, guards against misspelled properties | instance with correct field |
| dynamicRef | strict-tree schema, guards against misspelled properties | instance with misspelled field |
| dynamicRef | tests for implementation dynamic anchor and reference link | correct extended schema |
| dynamicRef | tests for implementation dynamic anchor and reference link | incorrect extended schema |
| dynamicRef | tests for implementation dynamic anchor and reference link | incorrect parent schema |
| enum | heterogeneous enum validation | one of the enum is valid |
| enum | heterogeneous enum validation | valid object matches |
| exclusiveMaximum | exclusiveMaximum validation | ignores non-numbers |
| exclusiveMinimum | exclusiveMinimum validation | ignores non-numbers |
| items | a schema given for items | JavaScript pseudo-array is valid |
| items | a schema given for items | ignores non-arrays |
| items | items and subitems | fewer items is valid |
| items | items and subitems | too many items |
| items | items and subitems | too many sub-items |
| items | items and subitems | wrong sub-item |
| items | items with boolean schema (false) | any non-empty array is invalid |
| items | items with heterogeneous array | heterogeneous invalid instance |
| items | prefixItems with no additional items allowed | additional items are not permitted |
| maxItems | maxItems validation | ignores non-arrays |
| maxLength | maxLength validation | ignores non-strings |
| maxProperties | maxProperties validation | ignores arrays |
| maxProperties | maxProperties validation | ignores other non-objects |
| maxProperties | maxProperties validation | ignores strings |
| maximum | maximum validation | ignores non-numbers |
| minItems | minItems validation | ignores non-arrays |
| minLength | minLength validation | ignores non-strings |
| minProperties | minProperties validation | ignores arrays |
| minProperties | minProperties validation | ignores other non-objects |
| minProperties | minProperties validation | ignores strings |
| minimum | minimum validation | ignores non-numbers |
| minimum | minimum validation with signed integer | ignores non-numbers |
| multipleOf | by int | ignores non-numbers |
| oneOf | oneOf with missing optional property | first oneOf valid |
| oneOf | oneOf with missing optional property | second oneOf valid |
| oneOf | oneOf with required | first valid - valid |
| oneOf | oneOf with required | second valid - valid |
| pattern | pattern validation | ignores arrays |
| pattern | pattern validation | ignores booleans |
| pattern | pattern validation | ignores floats |
| pattern | pattern validation | ignores integers |
| pattern | pattern validation | ignores null |
| pattern | pattern validation | ignores objects |
| patternProperties | patternProperties validates properties matching a regex | ignores other non-objects |
| patternProperties | patternProperties validates properties matching a regex | ignores strings |
| prefixItems | a schema given for prefixItems | JavaScript pseudo-array is valid |
| prefixItems | a schema given for prefixItems | empty array |
| prefixItems | a schema given for prefixItems | incomplete array of items |
| prefixItems | prefixItems with boolean schemas | array with one item is valid |
| prefixItems | prefixItems with boolean schemas | empty array is valid |
| properties | object properties validation | ignores other non-objects |
| properties | properties whose names are Javascript object property names | __proto__ not valid |
| properties | properties whose names are Javascript object property names | all present and valid |
| properties | properties whose names are Javascript object property names | constructor not valid |
| properties | properties whose names are Javascript object property names | ignores arrays |
| properties | properties whose names are Javascript object property names | ignores other non-objects |
| properties | properties whose names are Javascript object property names | none of the properties mentioned |
| properties | properties whose names are Javascript object property names | toString not valid |
| properties | properties, patternProperties, additionalProperties interaction | patternProperty validates nonproperty |
| propertyNames | propertyNames validation | ignores other non-objects |
| propertyNames | propertyNames validation | ignores strings |
| propertyNames | propertyNames validation | some property names invalid |
| propertyNames | propertyNames validation with pattern | non-matching property name is invalid |
| propertyNames | propertyNames with boolean schema false | object with any properties is invalid |
| propertyNames | propertyNames with const | object with any other property is invalid |
| propertyNames | propertyNames with enum | object with any other property is invalid |
| ref | URN ref with nested pointer ref | a non-string is invalid |
| ref | URN ref with nested pointer ref | a string is valid |
| ref | refs with relative uris and defs | invalid on inner field |
| ref | refs with relative uris and defs | invalid on outer field |
| ref | refs with relative uris and defs | valid on both fields |
| ref | relative refs with absolute uris and defs | invalid on inner field |
| ref | relative refs with absolute uris and defs | invalid on outer field |
| ref | relative refs with absolute uris and defs | valid on both fields |
| ref | remote ref, containing refs itself | remote ref invalid |
| ref | remote ref, containing refs itself | remote ref valid |
| refRemote | anchor within remote ref | remote anchor invalid |
| refRemote | anchor within remote ref | remote anchor valid |
| refRemote | retrieved nested refs resolve relative to their URI not $id | number is invalid |
| refRemote | retrieved nested refs resolve relative to their URI not $id | string is valid |
| required | required properties whose names are Javascript object property names | __proto__ present |
| required | required properties whose names are Javascript object property names | all present |
| required | required properties whose names are Javascript object property names | constructor present |
| required | required properties whose names are Javascript object property names | ignores arrays |
| required | required properties whose names are Javascript object property names | ignores other non-objects |
| required | required properties whose names are Javascript object property names | none of the properties mentioned |
| required | required properties whose names are Javascript object property names | toString present |
| required | required validation | ignores other non-objects |
| required | required validation | ignores strings |
| required | required validation | non-present required property is invalid |
| required | required with escaped characters | object with some properties missing is invalid |
| type | object type matches objects | an array is not an object |
| uniqueItems | uniqueItems with an array of items and additionalItems=false | extra items are invalid even if unique |
| uniqueItems | uniqueItems=false with an array of items and additionalItems=false | extra items are invalid even if unique |
| vocabulary | schema that uses custom metaschema with with no validation vocabulary | no validation: invalid number, but it still validates |

### draft2019-09 (146 failures)

| Keyword | Group | Test |
| ------- | ----- | ---- |
| additionalItems | additionalItems as false without items | ignores non-arrays |
| additionalItems | additionalItems with heterogeneous array | heterogeneous invalid instance |
| additionalItems | array of items with no additionalItems permitted | additional items are not permitted |
| additionalProperties | additionalProperties being false does not allow other properties | ignores arrays |
| additionalProperties | additionalProperties being false does not allow other properties | ignores other non-objects |
| additionalProperties | additionalProperties being false does not allow other properties | ignores strings |
| additionalProperties | additionalProperties being false does not allow other properties | patternProperties are not additional properties |
| additionalProperties | additionalProperties does not look in applicators | properties defined in allOf are not examined |
| contains | contains keyword validation | not array is valid |
| contains | contains keyword with boolean schema false | non-arrays are valid |
| defs | validate definition against metaschema | invalid definition schema |
| defs | validate definition against metaschema | valid definition schema |
| dependentRequired | empty dependents | non-object is valid |
| dependentRequired | single dependency | ignores other non-objects |
| dependentRequired | single dependency | ignores strings |
| dependentSchemas | dependencies with escaped characters | quoted quote |
| dependentSchemas | dependencies with escaped characters | quoted quote invalid under dependent schema |
| dependentSchemas | single dependency | ignores other non-objects |
| dependentSchemas | single dependency | ignores strings |
| enum | heterogeneous enum validation | one of the enum is valid |
| enum | heterogeneous enum validation | valid object matches |
| exclusiveMaximum | exclusiveMaximum validation | ignores non-numbers |
| exclusiveMinimum | exclusiveMinimum validation | ignores non-numbers |
| items | a schema given for items | JavaScript pseudo-array is valid |
| items | a schema given for items | ignores non-arrays |
| items | an array of schemas for items | JavaScript pseudo-array is valid |
| items | an array of schemas for items | empty array |
| items | an array of schemas for items | incomplete array of items |
| items | items and subitems | fewer items is valid |
| items | items and subitems | too many items |
| items | items and subitems | too many sub-items |
| items | items and subitems | wrong sub-item |
| items | items with boolean schema (false) | any non-empty array is invalid |
| items | items with boolean schemas | array with one item is valid |
| items | items with boolean schemas | empty array is valid |
| maxItems | maxItems validation | ignores non-arrays |
| maxLength | maxLength validation | ignores non-strings |
| maxProperties | maxProperties validation | ignores arrays |
| maxProperties | maxProperties validation | ignores other non-objects |
| maxProperties | maxProperties validation | ignores strings |
| maximum | maximum validation | ignores non-numbers |
| minItems | minItems validation | ignores non-arrays |
| minLength | minLength validation | ignores non-strings |
| minProperties | minProperties validation | ignores arrays |
| minProperties | minProperties validation | ignores other non-objects |
| minProperties | minProperties validation | ignores strings |
| minimum | minimum validation | ignores non-numbers |
| minimum | minimum validation with signed integer | ignores non-numbers |
| multipleOf | by int | ignores non-numbers |
| oneOf | oneOf with missing optional property | first oneOf valid |
| oneOf | oneOf with missing optional property | second oneOf valid |
| oneOf | oneOf with required | first valid - valid |
| oneOf | oneOf with required | second valid - valid |
| pattern | pattern validation | ignores arrays |
| pattern | pattern validation | ignores booleans |
| pattern | pattern validation | ignores floats |
| pattern | pattern validation | ignores integers |
| pattern | pattern validation | ignores null |
| pattern | pattern validation | ignores objects |
| patternProperties | patternProperties validates properties matching a regex | ignores other non-objects |
| patternProperties | patternProperties validates properties matching a regex | ignores strings |
| properties | object properties validation | ignores other non-objects |
| properties | properties whose names are Javascript object property names | __proto__ not valid |
| properties | properties whose names are Javascript object property names | all present and valid |
| properties | properties whose names are Javascript object property names | constructor not valid |
| properties | properties whose names are Javascript object property names | ignores arrays |
| properties | properties whose names are Javascript object property names | ignores other non-objects |
| properties | properties whose names are Javascript object property names | none of the properties mentioned |
| properties | properties whose names are Javascript object property names | toString not valid |
| properties | properties, patternProperties, additionalProperties interaction | patternProperty validates nonproperty |
| propertyNames | propertyNames validation | ignores other non-objects |
| propertyNames | propertyNames validation | ignores strings |
| propertyNames | propertyNames validation | some property names invalid |
| propertyNames | propertyNames validation with pattern | non-matching property name is invalid |
| propertyNames | propertyNames with boolean schema false | object with any properties is invalid |
| propertyNames | propertyNames with const | object with any other property is invalid |
| propertyNames | propertyNames with enum | object with any other property is invalid |
| recursiveRef | $recursiveRef with $recursiveAnchor: false works like $ref | integer does not match as a property value |
| recursiveRef | $recursiveRef with $recursiveAnchor: false works like $ref | integer matches at the outer level |
| recursiveRef | $recursiveRef with $recursiveAnchor: false works like $ref | single level match |
| recursiveRef | $recursiveRef with $recursiveAnchor: false works like $ref | two levels, integer does not match as a property value |
| recursiveRef | $recursiveRef with $recursiveAnchor: false works like $ref | two levels, properties match with inner definition |
| recursiveRef | $recursiveRef with nesting | integer matches at the outer level |
| recursiveRef | $recursiveRef with nesting | integer now matches as a property value |
| recursiveRef | $recursiveRef with nesting | single level match |
| recursiveRef | $recursiveRef with nesting | two levels, properties match with $recursiveRef |
| recursiveRef | $recursiveRef with nesting | two levels, properties match with inner definition |
| recursiveRef | $recursiveRef with no $recursiveAnchor in the initial target schema resource | leaf node does not match: recursion uses the inner schema |
| recursiveRef | $recursiveRef with no $recursiveAnchor in the initial target schema resource | leaf node does not match; no recursion |
| recursiveRef | $recursiveRef with no $recursiveAnchor in the initial target schema resource | leaf node matches: recursion uses the inner schema |
| recursiveRef | $recursiveRef with no $recursiveAnchor in the outer schema resource | leaf node does not match: recursion only uses inner schema |
| recursiveRef | $recursiveRef with no $recursiveAnchor in the outer schema resource | leaf node does not match; no recursion |
| recursiveRef | $recursiveRef with no $recursiveAnchor in the outer schema resource | leaf node matches: recursion only uses inner schema |
| recursiveRef | $recursiveRef with no $recursiveAnchor works like $ref | integer does not match as a property value |
| recursiveRef | $recursiveRef with no $recursiveAnchor works like $ref | integer matches at the outer level |
| recursiveRef | $recursiveRef with no $recursiveAnchor works like $ref | single level match |
| recursiveRef | $recursiveRef with no $recursiveAnchor works like $ref | two levels, integer does not match as a property value |
| recursiveRef | $recursiveRef with no $recursiveAnchor works like $ref | two levels, properties match with inner definition |
| recursiveRef | $recursiveRef without $recursiveAnchor works like $ref | match |
| recursiveRef | $recursiveRef without $recursiveAnchor works like $ref | mismatch |
| recursiveRef | $recursiveRef without $recursiveAnchor works like $ref | recursive match |
| recursiveRef | $recursiveRef without $recursiveAnchor works like $ref | recursive mismatch |
| recursiveRef | $recursiveRef without using nesting | integer does not match as a property value |
| recursiveRef | $recursiveRef without using nesting | integer matches at the outer level |
| recursiveRef | $recursiveRef without using nesting | single level match |
| recursiveRef | $recursiveRef without using nesting | two levels, no match |
| recursiveRef | $recursiveRef without using nesting | two levels, properties match with inner definition |
| recursiveRef | dynamic $recursiveRef destination (not predictable at schema compile time) | integer node |
| recursiveRef | dynamic $recursiveRef destination (not predictable at schema compile time) | numeric node |
| recursiveRef | multiple dynamic paths to the $recursiveRef keyword | recurse to anyLeafNode - floats are allowed |
| recursiveRef | multiple dynamic paths to the $recursiveRef keyword | recurse to integerNode - floats are not allowed |
| ref | $ref with $recursiveAnchor | extra items allowed for inner arrays |
| ref | $ref with $recursiveAnchor | extra items disallowed for root |
| ref | URN ref with nested pointer ref | a non-string is invalid |
| ref | URN ref with nested pointer ref | a string is valid |
| ref | refs with relative uris and defs | invalid on inner field |
| ref | refs with relative uris and defs | invalid on outer field |
| ref | refs with relative uris and defs | valid on both fields |
| ref | relative refs with absolute uris and defs | invalid on inner field |
| ref | relative refs with absolute uris and defs | invalid on outer field |
| ref | relative refs with absolute uris and defs | valid on both fields |
| ref | remote ref, containing refs itself | remote ref invalid |
| ref | remote ref, containing refs itself | remote ref valid |
| refRemote | anchor within remote ref | remote anchor invalid |
| refRemote | anchor within remote ref | remote anchor valid |
| refRemote | retrieved nested refs resolve relative to their URI not $id | number is invalid |
| refRemote | retrieved nested refs resolve relative to their URI not $id | string is valid |
| required | required properties whose names are Javascript object property names | __proto__ present |
| required | required properties whose names are Javascript object property names | all present |
| required | required properties whose names are Javascript object property names | constructor present |
| required | required properties whose names are Javascript object property names | ignores arrays |
| required | required properties whose names are Javascript object property names | ignores other non-objects |
| required | required properties whose names are Javascript object property names | none of the properties mentioned |
| required | required properties whose names are Javascript object property names | toString present |
| required | required validation | ignores other non-objects |
| required | required validation | ignores strings |
| required | required validation | non-present required property is invalid |
| required | required with escaped characters | object with some properties missing is invalid |
| type | object type matches objects | an array is not an object |
| unevaluatedItems | unevaluatedItems with $recursiveRef | with no unevaluated items |
| unevaluatedItems | unevaluatedItems with $recursiveRef | with unevaluated items |
| unevaluatedProperties | unevaluatedProperties with $recursiveRef | with no unevaluated properties |
| unevaluatedProperties | unevaluatedProperties with $recursiveRef | with unevaluated properties |
| uniqueItems | uniqueItems with an array of items and additionalItems=false | extra items are invalid even if unique |
| uniqueItems | uniqueItems=false with an array of items and additionalItems=false | extra items are invalid even if unique |
| vocabulary | schema that uses custom metaschema with with no validation vocabulary | no validation: invalid number, but it still validates |

### draft7 (95 failures)

| Keyword | Group | Test |
| ------- | ----- | ---- |
| additionalItems | additionalItems as false without items | ignores non-arrays |
| additionalItems | additionalItems with heterogeneous array | heterogeneous invalid instance |
| additionalItems | array of items with no additionalItems permitted | additional items are not permitted |
| additionalProperties | additionalProperties being false does not allow other properties | ignores arrays |
| additionalProperties | additionalProperties being false does not allow other properties | ignores other non-objects |
| additionalProperties | additionalProperties being false does not allow other properties | ignores strings |
| additionalProperties | additionalProperties being false does not allow other properties | patternProperties are not additional properties |
| additionalProperties | additionalProperties does not look in applicators | properties defined in allOf are not examined |
| contains | contains keyword validation | not array is valid |
| contains | contains keyword with boolean schema false | non-arrays are valid |
| definitions | validate definition against metaschema | invalid definition schema |
| dependencies | dependencies | ignores other non-objects |
| dependencies | dependencies | ignores strings |
| dependencies | dependencies with empty array | non-object is valid |
| dependencies | dependencies with escaped characters | invalid object 3 |
| enum | heterogeneous enum validation | one of the enum is valid |
| enum | heterogeneous enum validation | valid object matches |
| exclusiveMaximum | exclusiveMaximum validation | ignores non-numbers |
| exclusiveMinimum | exclusiveMinimum validation | ignores non-numbers |
| items | a schema given for items | JavaScript pseudo-array is valid |
| items | a schema given for items | ignores non-arrays |
| items | an array of schemas for items | JavaScript pseudo-array is valid |
| items | an array of schemas for items | empty array |
| items | an array of schemas for items | incomplete array of items |
| items | items and subitems | fewer items is valid |
| items | items and subitems | too many items |
| items | items and subitems | too many sub-items |
| items | items and subitems | wrong sub-item |
| items | items with boolean schema (false) | any non-empty array is invalid |
| items | items with boolean schemas | array with one item is valid |
| items | items with boolean schemas | empty array is valid |
| maxItems | maxItems validation | ignores non-arrays |
| maxLength | maxLength validation | ignores non-strings |
| maxProperties | maxProperties validation | ignores arrays |
| maxProperties | maxProperties validation | ignores other non-objects |
| maxProperties | maxProperties validation | ignores strings |
| maximum | maximum validation | ignores non-numbers |
| minItems | minItems validation | ignores non-arrays |
| minLength | minLength validation | ignores non-strings |
| minProperties | minProperties validation | ignores arrays |
| minProperties | minProperties validation | ignores other non-objects |
| minProperties | minProperties validation | ignores strings |
| minimum | minimum validation | ignores non-numbers |
| minimum | minimum validation with signed integer | ignores non-numbers |
| multipleOf | by int | ignores non-numbers |
| oneOf | oneOf with missing optional property | first oneOf valid |
| oneOf | oneOf with missing optional property | second oneOf valid |
| oneOf | oneOf with required | first valid - valid |
| oneOf | oneOf with required | second valid - valid |
| pattern | pattern validation | ignores arrays |
| pattern | pattern validation | ignores booleans |
| pattern | pattern validation | ignores floats |
| pattern | pattern validation | ignores integers |
| pattern | pattern validation | ignores null |
| pattern | pattern validation | ignores objects |
| patternProperties | patternProperties validates properties matching a regex | ignores other non-objects |
| patternProperties | patternProperties validates properties matching a regex | ignores strings |
| properties | object properties validation | ignores other non-objects |
| properties | properties whose names are Javascript object property names | __proto__ not valid |
| properties | properties whose names are Javascript object property names | all present and valid |
| properties | properties whose names are Javascript object property names | constructor not valid |
| properties | properties whose names are Javascript object property names | ignores arrays |
| properties | properties whose names are Javascript object property names | ignores other non-objects |
| properties | properties whose names are Javascript object property names | none of the properties mentioned |
| properties | properties whose names are Javascript object property names | toString not valid |
| properties | properties, patternProperties, additionalProperties interaction | patternProperty validates nonproperty |
| propertyNames | propertyNames validation | ignores other non-objects |
| propertyNames | propertyNames validation | ignores strings |
| propertyNames | propertyNames validation | some property names invalid |
| propertyNames | propertyNames validation with pattern | non-matching property name is invalid |
| propertyNames | propertyNames with boolean schema false | object with any properties is invalid |
| propertyNames | propertyNames with const | object with any other property is invalid |
| propertyNames | propertyNames with enum | object with any other property is invalid |
| ref | refs with relative uris and defs | invalid on inner field |
| ref | refs with relative uris and defs | invalid on outer field |
| ref | refs with relative uris and defs | valid on both fields |
| ref | relative refs with absolute uris and defs | invalid on inner field |
| ref | relative refs with absolute uris and defs | invalid on outer field |
| ref | relative refs with absolute uris and defs | valid on both fields |
| refRemote | retrieved nested refs resolve relative to their URI not $id | number is invalid |
| refRemote | retrieved nested refs resolve relative to their URI not $id | string is valid |
| required | required properties whose names are Javascript object property names | __proto__ present |
| required | required properties whose names are Javascript object property names | all present |
| required | required properties whose names are Javascript object property names | constructor present |
| required | required properties whose names are Javascript object property names | ignores arrays |
| required | required properties whose names are Javascript object property names | ignores other non-objects |
| required | required properties whose names are Javascript object property names | none of the properties mentioned |
| required | required properties whose names are Javascript object property names | toString present |
| required | required validation | ignores other non-objects |
| required | required validation | ignores strings |
| required | required validation | non-present required property is invalid |
| required | required with escaped characters | object with some properties missing is invalid |
| type | object type matches objects | an array is not an object |
| uniqueItems | uniqueItems with an array of items and additionalItems=false | extra items are invalid even if unique |
| uniqueItems | uniqueItems=false with an array of items and additionalItems=false | extra items are invalid even if unique |

### draft6 (95 failures)

| Keyword | Group | Test |
| ------- | ----- | ---- |
| additionalItems | additionalItems as false without items | ignores non-arrays |
| additionalItems | additionalItems with heterogeneous array | heterogeneous invalid instance |
| additionalItems | array of items with no additionalItems permitted | additional items are not permitted |
| additionalProperties | additionalProperties being false does not allow other properties | ignores arrays |
| additionalProperties | additionalProperties being false does not allow other properties | ignores other non-objects |
| additionalProperties | additionalProperties being false does not allow other properties | ignores strings |
| additionalProperties | additionalProperties being false does not allow other properties | patternProperties are not additional properties |
| additionalProperties | additionalProperties does not look in applicators | properties defined in allOf are not examined |
| contains | contains keyword validation | not array is valid |
| contains | contains keyword with boolean schema false | non-arrays are valid |
| definitions | validate definition against metaschema | invalid definition schema |
| dependencies | dependencies | ignores other non-objects |
| dependencies | dependencies | ignores strings |
| dependencies | dependencies with empty array | non-object is valid |
| dependencies | dependencies with escaped characters | invalid object 3 |
| enum | heterogeneous enum validation | one of the enum is valid |
| enum | heterogeneous enum validation | valid object matches |
| exclusiveMaximum | exclusiveMaximum validation | ignores non-numbers |
| exclusiveMinimum | exclusiveMinimum validation | ignores non-numbers |
| items | a schema given for items | JavaScript pseudo-array is valid |
| items | a schema given for items | ignores non-arrays |
| items | an array of schemas for items | JavaScript pseudo-array is valid |
| items | an array of schemas for items | empty array |
| items | an array of schemas for items | incomplete array of items |
| items | items and subitems | fewer items is valid |
| items | items and subitems | too many items |
| items | items and subitems | too many sub-items |
| items | items and subitems | wrong sub-item |
| items | items with boolean schema (false) | any non-empty array is invalid |
| items | items with boolean schemas | array with one item is valid |
| items | items with boolean schemas | empty array is valid |
| maxItems | maxItems validation | ignores non-arrays |
| maxLength | maxLength validation | ignores non-strings |
| maxProperties | maxProperties validation | ignores arrays |
| maxProperties | maxProperties validation | ignores other non-objects |
| maxProperties | maxProperties validation | ignores strings |
| maximum | maximum validation | ignores non-numbers |
| minItems | minItems validation | ignores non-arrays |
| minLength | minLength validation | ignores non-strings |
| minProperties | minProperties validation | ignores arrays |
| minProperties | minProperties validation | ignores other non-objects |
| minProperties | minProperties validation | ignores strings |
| minimum | minimum validation | ignores non-numbers |
| minimum | minimum validation with signed integer | ignores non-numbers |
| multipleOf | by int | ignores non-numbers |
| oneOf | oneOf with missing optional property | first oneOf valid |
| oneOf | oneOf with missing optional property | second oneOf valid |
| oneOf | oneOf with required | first valid - valid |
| oneOf | oneOf with required | second valid - valid |
| pattern | pattern validation | ignores arrays |
| pattern | pattern validation | ignores booleans |
| pattern | pattern validation | ignores floats |
| pattern | pattern validation | ignores integers |
| pattern | pattern validation | ignores null |
| pattern | pattern validation | ignores objects |
| patternProperties | patternProperties validates properties matching a regex | ignores other non-objects |
| patternProperties | patternProperties validates properties matching a regex | ignores strings |
| properties | object properties validation | ignores other non-objects |
| properties | properties whose names are Javascript object property names | __proto__ not valid |
| properties | properties whose names are Javascript object property names | all present and valid |
| properties | properties whose names are Javascript object property names | constructor not valid |
| properties | properties whose names are Javascript object property names | ignores arrays |
| properties | properties whose names are Javascript object property names | ignores other non-objects |
| properties | properties whose names are Javascript object property names | none of the properties mentioned |
| properties | properties whose names are Javascript object property names | toString not valid |
| properties | properties, patternProperties, additionalProperties interaction | patternProperty validates nonproperty |
| propertyNames | propertyNames validation | ignores other non-objects |
| propertyNames | propertyNames validation | ignores strings |
| propertyNames | propertyNames validation | some property names invalid |
| propertyNames | propertyNames validation with pattern | non-matching property name is invalid |
| propertyNames | propertyNames with boolean schema false | object with any properties is invalid |
| propertyNames | propertyNames with const | object with any other property is invalid |
| propertyNames | propertyNames with enum | object with any other property is invalid |
| ref | refs with relative uris and defs | invalid on inner field |
| ref | refs with relative uris and defs | invalid on outer field |
| ref | refs with relative uris and defs | valid on both fields |
| ref | relative refs with absolute uris and defs | invalid on inner field |
| ref | relative refs with absolute uris and defs | invalid on outer field |
| ref | relative refs with absolute uris and defs | valid on both fields |
| refRemote | retrieved nested refs resolve relative to their URI not $id | number is invalid |
| refRemote | retrieved nested refs resolve relative to their URI not $id | string is valid |
| required | required properties whose names are Javascript object property names | __proto__ present |
| required | required properties whose names are Javascript object property names | all present |
| required | required properties whose names are Javascript object property names | constructor present |
| required | required properties whose names are Javascript object property names | ignores arrays |
| required | required properties whose names are Javascript object property names | ignores other non-objects |
| required | required properties whose names are Javascript object property names | none of the properties mentioned |
| required | required properties whose names are Javascript object property names | toString present |
| required | required validation | ignores other non-objects |
| required | required validation | ignores strings |
| required | required validation | non-present required property is invalid |
| required | required with escaped characters | object with some properties missing is invalid |
| type | object type matches objects | an array is not an object |
| uniqueItems | uniqueItems with an array of items and additionalItems=false | extra items are invalid even if unique |
| uniqueItems | uniqueItems=false with an array of items and additionalItems=false | extra items are invalid even if unique |

### draft4 (74 failures)

| Keyword | Group | Test |
| ------- | ----- | ---- |
| additionalItems | additionalItems as false without items | ignores non-arrays |
| additionalItems | additionalItems with heterogeneous array | heterogeneous invalid instance |
| additionalItems | array of items with no additionalItems permitted | additional items are not permitted |
| additionalProperties | additionalProperties being false does not allow other properties | ignores arrays |
| additionalProperties | additionalProperties being false does not allow other properties | ignores other non-objects |
| additionalProperties | additionalProperties being false does not allow other properties | ignores strings |
| additionalProperties | additionalProperties being false does not allow other properties | patternProperties are not additional properties |
| additionalProperties | additionalProperties does not look in applicators | properties defined in allOf are not examined |
| definitions | validate definition against metaschema | invalid definition schema |
| dependencies | dependencies | ignores other non-objects |
| dependencies | dependencies | ignores strings |
| dependencies | dependencies with escaped characters | invalid object 3 |
| enum | heterogeneous enum validation | one of the enum is valid |
| enum | heterogeneous enum validation | valid object matches |
| items | a schema given for items | JavaScript pseudo-array is valid |
| items | a schema given for items | ignores non-arrays |
| items | an array of schemas for items | JavaScript pseudo-array is valid |
| items | an array of schemas for items | empty array |
| items | an array of schemas for items | incomplete array of items |
| items | items and subitems | fewer items is valid |
| items | items and subitems | too many items |
| items | items and subitems | too many sub-items |
| items | items and subitems | wrong sub-item |
| maxItems | maxItems validation | ignores non-arrays |
| maxLength | maxLength validation | ignores non-strings |
| maxProperties | maxProperties validation | ignores arrays |
| maxProperties | maxProperties validation | ignores other non-objects |
| maxProperties | maxProperties validation | ignores strings |
| maximum | maximum validation | ignores non-numbers |
| maximum | maximum validation (explicit false exclusivity) | ignores non-numbers |
| minItems | minItems validation | ignores non-arrays |
| minLength | minLength validation | ignores non-strings |
| minProperties | minProperties validation | ignores arrays |
| minProperties | minProperties validation | ignores other non-objects |
| minProperties | minProperties validation | ignores strings |
| minimum | minimum validation | ignores non-numbers |
| minimum | minimum validation (explicit false exclusivity) | ignores non-numbers |
| minimum | minimum validation with signed integer | ignores non-numbers |
| multipleOf | by int | ignores non-numbers |
| oneOf | oneOf with missing optional property | first oneOf valid |
| oneOf | oneOf with missing optional property | second oneOf valid |
| oneOf | oneOf with required | first valid - valid |
| oneOf | oneOf with required | second valid - valid |
| pattern | pattern validation | ignores arrays |
| pattern | pattern validation | ignores booleans |
| pattern | pattern validation | ignores floats |
| pattern | pattern validation | ignores integers |
| pattern | pattern validation | ignores null |
| pattern | pattern validation | ignores objects |
| patternProperties | patternProperties validates properties matching a regex | ignores other non-objects |
| patternProperties | patternProperties validates properties matching a regex | ignores strings |
| properties | object properties validation | ignores other non-objects |
| properties | properties whose names are Javascript object property names | __proto__ not valid |
| properties | properties whose names are Javascript object property names | all present and valid |
| properties | properties whose names are Javascript object property names | constructor not valid |
| properties | properties whose names are Javascript object property names | ignores arrays |
| properties | properties whose names are Javascript object property names | ignores other non-objects |
| properties | properties whose names are Javascript object property names | none of the properties mentioned |
| properties | properties whose names are Javascript object property names | toString not valid |
| properties | properties, patternProperties, additionalProperties interaction | patternProperty validates nonproperty |
| required | required properties whose names are Javascript object property names | __proto__ present |
| required | required properties whose names are Javascript object property names | all present |
| required | required properties whose names are Javascript object property names | constructor present |
| required | required properties whose names are Javascript object property names | ignores arrays |
| required | required properties whose names are Javascript object property names | ignores other non-objects |
| required | required properties whose names are Javascript object property names | none of the properties mentioned |
| required | required properties whose names are Javascript object property names | toString present |
| required | required validation | ignores other non-objects |
| required | required validation | ignores strings |
| required | required validation | non-present required property is invalid |
| required | required with escaped characters | object with some properties missing is invalid |
| type | object type matches objects | an array is not an object |
| uniqueItems | uniqueItems with an array of items and additionalItems=false | extra items are invalid even if unique |
| uniqueItems | uniqueItems=false with an array of items and additionalItems=false | extra items are invalid even if unique |

### draft3 (42 failures)

| Keyword | Group | Test |
| ------- | ----- | ---- |
| additionalItems | additionalItems as false without items | ignores non-arrays |
| additionalItems | additionalItems with heterogeneous array | heterogeneous invalid instance |
| additionalItems | array of items with no additionalItems permitted | additional items are not permitted |
| additionalProperties | additionalProperties being false does not allow other properties | ignores arrays |
| additionalProperties | additionalProperties being false does not allow other properties | ignores other non-objects |
| additionalProperties | additionalProperties being false does not allow other properties | ignores strings |
| additionalProperties | additionalProperties being false does not allow other properties | patternProperties are not additional properties |
| additionalProperties | additionalProperties does not look in applicators | properties defined in extends are not examined |
| dependencies | dependencies | ignores other non-objects |
| dependencies | dependencies | ignores strings |
| dependencies | dependencies | missing dependency |
| divisibleBy | by int | ignores non-numbers |
| enum | heterogeneous enum validation | one of the enum is valid |
| items | a schema given for items | ignores non-arrays |
| maxItems | maxItems validation | ignores non-arrays |
| maxLength | maxLength validation | ignores non-strings |
| maximum | maximum validation | ignores non-numbers |
| maximum | maximum validation (explicit false exclusivity) | ignores non-numbers |
| minItems | minItems validation | ignores non-arrays |
| minLength | minLength validation | ignores non-strings |
| minimum | minimum validation | ignores non-numbers |
| minimum | minimum validation with signed integer | ignores non-numbers |
| pattern | pattern validation | ignores arrays |
| pattern | pattern validation | ignores booleans |
| pattern | pattern validation | ignores floats |
| pattern | pattern validation | ignores integers |
| pattern | pattern validation | ignores null |
| pattern | pattern validation | ignores objects |
| patternProperties | patternProperties validates properties matching a regex | ignores other non-objects |
| properties | object properties validation | ignores other non-objects |
| properties | properties, patternProperties, additionalProperties interaction | patternProperty validates nonproperty |
| required | required validation | present required property is valid |
| type | applies a nested schema | an object is invalid otherwise |
| type | object type matches objects | an array is not an object |
| type | types can include schemas | a boolean is invalid |
| type | types can include schemas | a float is invalid |
| type | types can include schemas | a string is invalid |
| type | types can include schemas | an integer is invalid |
| type | types can include schemas | null is invalid |
| type | types from separate schemas are merged | an integer is invalid |
| uniqueItems | uniqueItems with an array of items and additionalItems=false | extra items are invalid even if unique |
| uniqueItems | uniqueItems=false with an array of items and additionalItems=false | extra items are invalid even if unique |

## Categorized Failures

### not-supported (bundler architectural limits)

These failures are caused by bundler limitations that require significant architectural changes. Not planned for implementation.

| Category | Tests | Description |
| -------- | ----- | ----------- |
| $dynamicRef/$dynamicAnchor | 33 | draft2020-12 dynamic reference semantics require runtime evaluation context |
| $recursiveRef/$recursiveAnchor | 36 | draft2019-09 recursive reference semantics require runtime evaluation context |
| Relative URI with $id scoping | ~12 | bundler flattens schemas, losing `$id`-scoped relative refs |
| URN refs with nested pointers | ~4 | bundler doesn't resolve URN-based refs |
| Remote refs containing dynamic keywords | ~6 | fetching JSON Schema metaschema fails due to dynamic keywords |
| Nested $defs after remote fetch | ~8 | bundler's key-sanitization creates invalid lookup paths |

### adapter-native (fixable in valibot adapter)

| Category | Tests | Description |
| -------- | ----- | ----------- |
| "ignores non-X" type coercion failures | ~100 | valibot validators reject wrong types (e.g., maxLength rejects non-strings instead of ignoring). JSON Schema says keyword-specific validators should only apply to matching types. **Requires passthrough pipe for type mismatches.** |
| type: "object" accepts arrays | 6 | valibot `v.object()` accepts arrays, but JSON Schema "object" excludes arrays |
| propertyNames validation | 5 | propertyNames schema not being enforced (always passes) |
| required validation | ~4 | required properties not being checked correctly |
| tuple/prefixItems validation | ~12 | tuple schemas (items/prefixItems as array) have issues with partial arrays and additionalItems |
| additionalItems enforcement | ~6 | additionalItems: false not blocking extra items |
| JS prototype property names | ~14 | properties named `__proto__`, `constructor`, `toString` cause runtime errors in valibot |
| escaped characters in property keys | ~4 | properties with escaped quotes in keys not handled correctly |
| heterogeneous enum with object | ~4 | enum containing object values fails to match |
| oneOf/required interaction | ~8 | oneOf branches with required fields incorrectly fail |
| patternProperties + additionalProperties interaction | ~6 | patternProperties not properly excluding from additionalProperties check |

### forced-emulation (non-idiomatic, not recommended)

| Category | Tests | Description | Reason to skip |
| -------- | ----- | ----------- | -------------- |
| additionalProperties + applicators | 6 | requires tracking evaluated properties across allOf/extends | needs unevaluatedProperties semantics |
| validate definition against metaschema | 5 | requires validating schema definitions against metaschema | meta-validation out of scope |
| custom vocabulary (no validation) | 2 | custom metaschema disabling validation vocabulary | edge case, no real use |
| draft3 type with schemas | 7 | draft3 allows `type: [{...schema...}]` | legacy draft3 feature, rarely used |

### core-missing (blocked by core parser/bundler)

None - all bundler issues have been addressed in US-006/007/008.

## Expected Regressions

None expected. All failures are either:
1. Bundler limitations (not-supported)
2. Adapter implementation gaps (adapter-native)
3. Non-idiomatic edge cases we intentionally don't support

## Observations

1. **Lower baseline than Zod**: valibot at 85-89% vs zod at 93-97% - significant gap
2. **Massive "ignores non-X" failures**: ~100 tests fail because valibot validators reject wrong types instead of ignoring them. This is the single biggest category.
3. **JS prototype property bug**: Properties named `__proto__`, `constructor`, `toString` cause valibot runtime errors (`this.entries[key]._run is not a function`). This is a valibot library bug, not adapter issue.
4. **Tuple handling broken**: items/prefixItems as arrays have multiple issues (empty array handling, partial arrays, additionalItems enforcement)
5. **propertyNames not enforced**: All propertyNames validation tests fail - implementation may be missing
6. **Required validation issues**: basic required validation failing in some cases
7. **Dynamic/recursive refs**: Same bundler limitations as zod (~67 tests)

## Recommended Implementation Stories

Based on this discovery, the following implementation stories would provide the most value:

1. **valibot adapter: type-coercion passthrough** - would fix ~100 "ignores non-X" tests. Keywords like maxLength, pattern, minimum etc should pass (not validate) when the value isn't the expected type.

2. **valibot adapter: fix type:object to reject arrays** - would fix 6 tests. Currently `v.object()` accepts arrays.

3. **valibot adapter: fix propertyNames validation** - would fix 5 tests. propertyNames schema needs to be enforced.

Note: JS prototype property bug is upstream valibot issue. Tuple handling may require significant refactor.
