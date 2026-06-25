package main

// RomanMapping associates an integer value with its corresponding Roman numeral token.
type RomanMapping struct {
	Value  int
	Symbol string
}

// IntToRoman converts an integer to a Roman numeral string.
func IntToRoman(num int) string {
	mappings := []RomanMapping{
		{1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"},
		{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
		{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
	}

	var result string

	for _, mapping := range mappings {
		if num == 0 {
			break
		}
		count := num / mapping.Value
		if count > 0 {
			for i := 0; i < count; i++ {
				result += mapping.Symbol
			}
			num %= mapping.Value
		}
	}

	return result
}

func IntToRomanV2(num int) string {
	mappings := []RomanMapping{
		{1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"},
		{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
		{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
	}

	var result string
	for _, mapping := range mappings {
		if num == 0 {
			break
		}

		count := num / mapping.Value
		if count > 0 {
			for i := 0; i < count; i++ {
				result += mapping.Symbol
			}
			num %= mapping.Value
		}
	}

	return result
}

func IntToRomanV3(num int) string {
	mappings := []RomanMapping{
		{1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"},
		{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
		{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
	}

	var result string
	for _, mapping := range mappings {
		if num == 0 {
			break
		}
		count := num / mapping.Value
		if count > 0 {
			for i := 0; i < count; i++ {
				result += mapping.Symbol
			}
			num %= mapping.Value
		}
	}

	return result
}
