package json

import (
	"encoding/json"
)

type RawMessage = json.RawMessage

func Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}
