package websockets

import (
	"crypto/ed25519"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/url"
	"sort"

	jsoniter "github.com/json-iterator/go"
)

var utils Utils

type Utils struct{}

func (*Utils) RemoveDuplicates(input []string) []string {
	seen := make(map[string]bool)
	result := make([]string, 0, len(input))

	for _, str := range input {
		if _, exists := seen[str]; !exists {
			seen[str] = true
			result = append(result, str)
		}
	}

	return result
}

// createQueryString transforms a map[string]interface{} into a query string
func (*Utils) CreateQueryString(params map[string]interface{}, sorted bool) string {
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
			jsonValue, err := jsoniter.Marshal(v)
			if err != nil {
				Logger.ERROR(fmt.Sprintf("[VERBOSE] Error marshaling slice for key %s: %v\n", key, err))
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
			Logger.ERROR(fmt.Sprintf("[VERBOSE] Error adding parameter: invalid type detected, received %v", v))
		}
	}

	// Process each key-value pair
	for _, key := range keys {
		addToQuery(key, params[key])
	}

	return query.Encode()
}

// SignEd25519 signs a message with a private Ed25519 key.
func (*Utils) SignEd25519(message []byte, privateKey ed25519.PrivateKey) []byte {
	return ed25519.Sign(privateKey, message)
}

func (*Utils) CreateHMACSignature(value string, privatekey string) (string, error) {
	h := hmac.New(sha256.New, []byte(privatekey))
	_, err := h.Write([]byte(value))
	if err != nil {
		return "", err
	}

	signature := hex.EncodeToString(h.Sum(nil))
	return signature, nil
}
