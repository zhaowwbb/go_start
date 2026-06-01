package main

import (
	"testing"
)

func TestConvert(t *testing.T) {
	tests := []struct {
		name     string
		inputS   string
		numRows  int
		expected string
	}{
		{
			name:     "Standard zigzag with 3 rows",
			inputS:   "PAYPALISHIRING",
			numRows:  3,
			expected: "PAHNAPLSIIGYIR",
		},
		{
			name:     "Standard zigzag with 4 rows",
			inputS:   "PAYPALISHIRING",
			numRows:  4,
			expected: "PINALSIGYAHRPI",
		},
		{
			name:     "Single character string",
			inputS:   "A",
			numRows:  1,
			expected: "A",
		},
		{
			name:     "Multiple characters with single row",
			inputS:   "AB",
			numRows:  1,
			expected: "AB",
		},
		{
			name:     "numRows greater than string length",
			inputS:   "ABCD",
			numRows:  5,
			expected: "ABCD",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// actual := Convert(tt.inputS, tt.numRows)
			// actual := ConvertV2(tt.inputS, tt.numRows)
			actual := ConvertV3(tt.inputS, tt.numRows)
			if actual != tt.expected {
				t.Errorf("Convert(%q, %d) = %q; want %q", tt.inputS, tt.numRows, actual, tt.expected)
			}
		})
	}
}
