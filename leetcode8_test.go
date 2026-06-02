package main

import (
	"math"
	"testing"
)

func TestMyAtoi(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "Standard positive integer conversion",
			input:    "42",
			expected: 42,
		},
		{
			name:     "Leading spaces and negative sign conversion",
			input:    "   -42",
			expected: -42,
		},
		{
			name:     "Conversion stops at first non-digit letter",
			input:    "1337c0d3",
			expected: 1337,
		},
		{
			name:     "Conversion stops due to internal non-digit symbols",
			input:    "0-1",
			expected: 0,
		},
		{
			name:     "Starts with words, should return 0",
			input:    "words and 987",
			expected: 0,
		},
		{
			name:     "Overflow upper clamping limit test",
			input:    "2147483648", // MaxInt32 + 1
			expected: math.MaxInt32,
		},
		{
			name:     "Underflow lower clamping limit test",
			input:    "-91283472332",
			expected: math.MinInt32,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// actual := MyAtoi(tt.input)
			// actual := MyAtoiV2(tt.input)
			actual := MyAtoiV3(tt.input)
			if actual != tt.expected {
				t.Errorf("MyAtoi(%q) = %d; want %d", tt.input, actual, tt.expected)
			}
		})
	}
}
