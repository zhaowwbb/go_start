// leetcode32.go
package main

// LongestValidParentheses calculates the length of the longest well-formed parentheses substring.
// Time Complexity: O(n)
// Space Complexity: O(n)
func LongestValidParentheses(s string) int {
	if len(s) == 0 {
		return 0
	}

	// In Go, we can implement a basic stack using a slice of integers.
	// We seed it with -1 as our initial base boundary marker.
	stack := []int{-1}
	maxLength := 0

	for i := 0; i < len(s); i++ {
		char := s[i]

		if char == '(' {
			// Push the current index onto the stack
			stack = append(stack, i)
		} else {
			// Pop the last element from the stack
			stack = stack[:len(stack)-1]

			if len(stack) == 0 {
				// If the stack is empty, this ')' is unmatched.
				// It becomes the new base boundary.
				stack = append(stack, i)
			} else {
				// Calculate the length of the valid match using the index at the top of the stack
				currentLength := i - stack[len(stack)-1]
				if currentLength > maxLength {
					maxLength = currentLength
				}
			}
		}
	}

	return maxLength
}

func LongestValidParenthesesV2(s string) int {
	if len(s) == 0 {
		return 0
	}
	stack := []int{-1}
	max_len := 0

	for i := 0; i < len(s); i++ {
		char := s[i]
		if char == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				currentLength := i - stack[len(stack)-1]
				if currentLength > max_len {
					max_len = currentLength
				}
			}
		}
	}
	return max_len
}
