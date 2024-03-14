package ollama_test

import (
	"testing"

	"github.com/samyfodil/go-ollama-tau-sdk"

	"gotest.tools/v3/assert"
)

func TestBytesSliceToBytes(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]byte
		expected []byte
	}{
		{
			name:     "empty input",
			input:    [][]byte{},
			expected: []byte{},
		},
		{
			name:     "single element",
			input:    [][]byte{{0x01, 0x02, 0x03}},
			expected: append([]byte{0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, 0x01, 0x02, 0x03),
		},
		{
			name: "multiple elements",
			input: [][]byte{
				{0x01, 0x02, 0x03},
				{0x04, 0x05},
			},
			expected: append(
				append([]byte{0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, 0x01, 0x02, 0x03),
				append([]byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, 0x04, 0x05)...,
			),
		},
		{
			name:     "nil element in slice",
			input:    [][]byte{nil},
			expected: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := ollama.BytesSliceToBytes(tc.input)
			assert.DeepEqual(t, tc.expected, result)
		})
	}
}

func TestBytesToBytesSlice(t *testing.T) {
	tests := []struct {
		name        string
		input       []byte
		expected    [][]byte
		expectError bool
	}{
		{
			name:        "empty input",
			input:       []byte{},
			expected:    [][]byte{},
			expectError: false,
		},
		{
			name:  "valid single element",
			input: append([]byte{0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, 0x01, 0x02, 0x03),
			expected: [][]byte{
				{0x01, 0x02, 0x03},
			},
			expectError: false,
		},
		{
			name: "valid multiple elements",
			input: append(
				append([]byte{0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, 0x01, 0x02, 0x03),
				append([]byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, 0x04, 0x05)...,
			),
			expected: [][]byte{
				{0x01, 0x02, 0x03},
				{0x04, 0x05},
			},
			expectError: false,
		},
		{
			name:        "header out of bound",
			input:       []byte{0x01},
			expected:    nil,
			expectError: true,
		},
		{
			name:        "payload out of bound",
			input:       append([]byte{0x05, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, 0x01, 0x02),
			expected:    nil,
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := ollama.BytesToBytesSlice(tc.input)
			if tc.expectError {
				assert.ErrorContains(t, err, "")
			} else {
				assert.NilError(t, err)
				assert.DeepEqual(t, tc.expected, result)
			}
		})
	}
}
