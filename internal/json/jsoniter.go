package json

import jsoniter "github.com/json-iterator/go"

// This is a faster JSON parser
var jsoniter_instance = jsoniter.Config{
	EscapeHTML:              false, // Avoids escaping HTML (faster for most JSON use cases)
	SortMapKeys:             false, // Disables map key sorting (reduces overhead)
	MarshalFloatWith6Digits: true,  // Optimizes float marshaling
	CaseSensitive:           true,  // Whether its case-insensitive matching
	OnlyTaggedField:         false, // Ignores untagged struct fields (if you don't need them)
	// ValidateJsonRawMessage:  false, // Skips validation of raw messages
	// ObjectFieldMustBeSimpleString: true,  // Optimizes string field parsing
}.Froze()
