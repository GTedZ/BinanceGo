package serializer

import (
	"encoding/binary"
	"fmt"
	"math"
)

type bigEndian struct{}

var BigEndian = bigEndian{}

// Serialization

func (b bigEndian) SerializeInt64(i int64) []byte {
	// Serialize into 8 bytes using big-endian format
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))

	return buf
}

func (b bigEndian) SerializeFloat64(f float64) []byte {
	// Serialize into 8 bytes using big-endian format
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, math.Float64bits(f))

	return buf
}

// Deserialization

func (b bigEndian) DeserializeInt64(data []byte) (int64, error) {
	if len(data) != 8 {
		return 0, fmt.Errorf("invalid data length for int64 deserialization")
	}

	return int64(binary.BigEndian.Uint64(data)), nil
}

func (b bigEndian) DeserializeFloat64(data []byte) (float64, error) {
	if len(data) != 8 {
		return 0, fmt.Errorf("invalid data length for float64 deserialization")
	}

	return math.Float64frombits(binary.BigEndian.Uint64(data)), nil
}
