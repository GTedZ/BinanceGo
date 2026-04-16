package serializer

import (
	"encoding/binary"
	"fmt"
	"math"
)

type littleEndian struct{}

var LittleEndian = littleEndian{}

// Serialization

func (b littleEndian) SerializeInt64(i int64) []byte {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, uint64(i))
	return buf
}

func (b littleEndian) SerializeFloat64(f float64) []byte {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, math.Float64bits(f))
	return buf
}

// Deserialization

func (b littleEndian) DeserializeInt64(data []byte) (int64, error) {
	if len(data) != 8 {
		return 0, fmt.Errorf("invalid data length for int64 deserialization")
	}

	return int64(binary.LittleEndian.Uint64(data)), nil
}

func (b littleEndian) DeserializeFloat64(data []byte) (float64, error) {
	if len(data) != 8 {
		return 0, fmt.Errorf("invalid data length for float64 deserialization")
	}

	return math.Float64frombits(binary.LittleEndian.Uint64(data)), nil
}
