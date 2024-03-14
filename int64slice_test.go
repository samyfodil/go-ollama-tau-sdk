package ollama_test

import (
	"testing"

	"github.com/samyfodil/go-ollama-tau-sdk"

	"gotest.tools/v3/assert"
)

func TestBytesToInt64Slice(t *testing.T) {
	tests := []struct {
		name        string
		input       []byte
		expected    []int64
		expectError bool
	}{
		{
			name:        "empty input",
			input:       []byte{},
			expected:    []int64{},
			expectError: false,
		},
		{
			name:        "valid conversion",
			input:       []byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			expected:    []int64{1, 2},
			expectError: false,
		},
		{
			name:        "incomplete bytes",
			input:       []byte{0x01, 0x00, 0x00},
			expected:    nil,
			expectError: true, // Expecting error due to binary.Read failure on incomplete data
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := ollama.BytesToInt64Slice(tc.input)
			if tc.expectError {
				assert.ErrorContains(t, err, "")
			} else {
				assert.NilError(t, err)
				assert.DeepEqual(t, tc.expected, result)
			}
		})
	}
}

func TestInt64SliceToBytes(t *testing.T) {
	tests := []struct {
		name     string
		input    []int64
		expected []byte
	}{
		{
			name:     "empty input",
			input:    []int64{},
			expected: []byte{},
		},
		{
			name:     "valid conversion",
			input:    []int64{1, 2},
			expected: []byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := ollama.Int64SliceToBytes(tc.input)

			if len(result) == 0 && len(tc.expected) == 0 {
				// Pass the test if both are empty, regardless of nil or not
				assert.Assert(t, result == nil || len(result) == 0)
			} else {
				assert.DeepEqual(t, tc.expected, result)
			}
		})
	}
}
