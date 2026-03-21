package json

import (
	stdjson "encoding/json"
)

var api = jsoniter_instance

type RawMessage = stdjson.RawMessage

func Marshal(v any) ([]byte, error) {
	return api.Marshal(v)
}

func Unmarshal(data []byte, v any) error {
	return api.Unmarshal(data, v)
}
