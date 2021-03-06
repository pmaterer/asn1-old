package asn1

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeString(t *testing.T) {
	tests := []struct {
		name        string
		expected    []byte
		value       interface{}
		errExpected bool
	}{
		{
			name:        "Test encode string",
			expected:    []byte{0x66, 0x6f, 0x6f, 0x62, 0x61, 0x72},
			value:       "foobar",
			errExpected: false,
		},
		{
			name:        "Test encode string error",
			expected:    []byte{},
			value:       1582,
			errExpected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := encodeString(reflect.ValueOf(tt.value))
			if tt.errExpected {
				assert.Error(t, err)
			} else {
				assert.Equal(t, tt.expected, b)
			}
		})
	}
}

func TestIsValidPrintableString(t *testing.T) {
	tests := []struct {
		name     string
		expected bool
		value    string
	}{
		{
			name:     "Test encode proper printable string",
			expected: true,
			value:    "print('Woah')",
		},
		{
			name:     "Test encode bad printable string",
			expected: false,
			value:    "nuh uh uh!",
		},
		{
			name:     "Test encode bad printable string",
			expected: false,
			value:    "#asn1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ok := isValidPrintableString(tt.value)
			assert.Equal(t, tt.expected, ok)
		})
	}
}
