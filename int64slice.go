package ollama

import (
	"encoding/binary"
	"errors"
)

func BytesToInt64Slice(payload []byte) ([]int64, error) {
	if payload == nil {
		return nil, nil
	}

	if len(payload)%8 != 0 {
		return nil, errors.New("incomplete bytes for int64 conversion")
	}

	length := len(payload) / 8
	result := make([]int64, length)

	for i := 0; i < length; i++ {
		offset := i * 8
		result[i] = int64(binary.LittleEndian.Uint64(payload[offset : offset+8]))
	}

	return result, nil
}

func Int64SliceToBytes(slice []int64) []byte {
	// Each int64 is 8 bytes, so preallocate the exact size needed.
	result := make([]byte, len(slice)*8)

	for i, s := range slice {
		// Convert each int64 into 8 bytes and put them directly into the result slice.
		binary.LittleEndian.PutUint64(result[i*8:], uint64(s))
	}

	return result
}
