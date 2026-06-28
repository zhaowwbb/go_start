package main

// IsValid determines if an input string containing brackets is valid.
// It uses a stack slice to track opening brackets in O(N) time and O(N) space.
func IsValid(s string) bool {
	// Map closing brackets to their corresponding opening brackets
	bracketMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	// Use a slice as a stack
	var stack []rune

	for _, char := range s {
		// If it's a closing bracket
		if opening, exists := bracketMap[char]; exists {
			// If stack is empty or top element doesn't match the expected opening bracket
			if len(stack) == 0 || stack[len(stack)-1] != opening {
				return false
			}
			// Pop the top element
			stack = stack[:len(stack)-1]
		} else {
			// It's an opening bracket, push it onto the stack
			stack = append(stack, char)
		}
	}

	// If the stack is empty, all brackets were matched correctly
	return len(stack) == 0
}

func IsValidV2(s string) bool {
	bracketMap := map[rune]rune{')': '(', ']': '[', '}': '{'}
	var stack []rune
	for _, c := range s {
		if value, exist := bracketMap[c]; exist {
			if len(stack) == 0 || stack[len(stack)-1] != value {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, c)
		}
	}
	return len(stack) == 0
}

func IsValidV3(s string) bool {
	bracketMap := map[rune]rune{')': '(', ']': '[', '}': '{'}
	var stack []rune

	for _, c := range s {
		if value, exist := bracketMap[c]; exist {
			if len(stack) == 0 || stack[len(stack)-1] != value {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, c)
		}
	}

	return len(stack) == 0
}
