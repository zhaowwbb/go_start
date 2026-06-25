package main

// RomanToInt converts a Roman numeral string to an integer.
func RomanToInt(s string) int {
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	total := 0
	n := len(s)

	for i := 0; i < n; i++ {
		currentVal := romanMap[s[i]]

		// If this is not the last character and its value is less than
		// the next character's value, it's a subtractive form (e.g., IV, IX)
		if i < n-1 && currentVal < romanMap[s[i+1]] {
			total -= currentVal
		} else {
			total += currentVal
		}
	}

	return total
}

func RomanToIntV2(s string) int {
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	total := 0
	n := len(s)
	for i := 0; i < n; i++ {
		currentVal := romanMap[s[i]]
		if i < n-1 && currentVal < romanMap[s[i+1]] {
			total -= currentVal
		} else {
			total += currentVal
		}
	}
	return total
}

func RomanToIntV3(s string) int {
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	total := 0
	n := len(s)
	for i := 0; i < n; i++ {
		currentVal := romanMap[s[i]]
		if i < n-1 && currentVal < romanMap[s[i+1]] {
			total -= currentVal
		} else {
			total += currentVal
		}
	}
	return total
}
