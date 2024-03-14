package ollama

import (
	"encoding/binary"
	"errors"
)

func BytesSliceToBytes(c [][]byte) []byte {
	out := make([]byte, 0)

	for _, s := range c {
		out = binary.LittleEndian.AppendUint64(out, uint64(len(s)))
		out = append(out, s...)
	}

	return out
}

func BytesToBytesSlice(ds []byte) ([][]byte, error) {
	result := make([][]byte, 0)

	for idx := 0; idx < len(ds); {
		if idx+8 >= len(ds) {
			return nil, errors.New("header out of bound")
		}
		size := int(binary.LittleEndian.Uint64(ds[idx : idx+8]))
		idx += 8
		if idx+size > len(ds) {
			return nil, errors.New("payload out of bound")
		}
		result = append(result, ds[idx:idx+size])
		idx += size
	}
	return result, nil
}
