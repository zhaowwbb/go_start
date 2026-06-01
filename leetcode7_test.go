package main

import (
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "Standard positive integer",
			input:    123,
			expected: 321,
		},
		{
			name:     "Standard negative integer",
			input:    -123,
			expected: -321,
		},
		{
			name:     "Integer ending with zero",
			input:    120,
			expected: 21,
		},
		{
			name:     "Zero test",
			input:    0,
			expected: 0,
		},
		{
			name:     "Overflow upper limit boundary check",
			input:    1534236469, // Reversing this exceeds math.MaxInt32
			expected: 0,
		},
		{
			name:     "Underflow lower limit boundary check",
			input:    -2147483648, // math.MinInt32 underflows when processed
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// actual := Reverse(tt.input)
			// actual := ReverseV2(tt.input)
			actual := ReverseV3(tt.input)
			if actual != tt.expected {
				t.Errorf("Reverse(%d) = %d; want %d", tt.input, actual, tt.expected)
			}
		})
	}
}
