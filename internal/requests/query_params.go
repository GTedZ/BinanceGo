package requests

import (
	"encoding/json"
	"fmt"
	"net/url"
	"sort"
)

// CreateQueryString transforms a map[string]interface{} into a query string
func CreateQueryString(params map[string]interface{}, sorted bool) string {
	if params == nil {
		return ""
	}

	// Extract keys to sort them if `sorted` is true
	keys := make([]string, 0, len(params))
	for key := range params {
		keys = append(keys, key)
	}

	if sorted {
		sort.Strings(keys)
	}

	query := url.Values{}

	// Helper function to process values
	var addToQuery func(key string, value interface{})
	addToQuery = func(key string, value interface{}) {
		switch v := value.(type) {
		case string:
			query.Add(key, v)
		case []string:
			// Encode slices as JSON arrays
			jsonValue, err := json.Marshal(v)
			if err != nil {
				fmt.Printf("[VERBOSE] Error marshaling slice for key %s: %v\n", key, err)
				return
			}
			query.Add(key, string(jsonValue)) // Add JSON-encoded array
		case []interface{}:
			for _, item := range v {
				addToQuery(key, item) // Recursively handle each item
			}
		case map[string]interface{}:
			// Handle nested maps with dot notation
			for subKey, subValue := range v {
				addToQuery(key+"."+subKey, subValue)
			}
		case int, int64, float64, bool: // Convert basic types to string
			query.Add(key, fmt.Sprintf("%v", v))
		default:
			fmt.Printf("[VERBOSE] Error adding parameter: invalid type detected, received %v", v)
		}
	}

	// Process each key-value pair
	for _, key := range keys {
		addToQuery(key, params[key])
	}

	return query.Encode()
}
