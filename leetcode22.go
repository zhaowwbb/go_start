package main

// GenerateParenthesis generates all combinations of well-formed parentheses for n pairs.
func GenerateParenthesis(n int) []string {
	var result []string
	if n <= 0 {
		return []string{""}
	}

	// Create a byte slice buffer to build strings efficiently
	backtrack("", 0, 0, n, &result)
	return result
}

func backtrack(currentStr string, openCount int, closeCount int, max int, result *[]string) {
	// Base case: If the current string reaches the maximum length (2 * max)
	if len(currentStr) == 2*max {
		*result = append(*result, currentStr)
		return
	}

	// If we can still add an opening parenthesis, do so
	if openCount < max {
		backtrack(currentStr+"(", openCount+1, closeCount, max, result)
	}

	// If we have more opening than closing parentheses, we can add a closing one
	if closeCount < openCount {
		backtrack(currentStr+")", openCount, closeCount+1, max, result)
	}
}

func GenerateParenthesisV2(n int) []string {
	var result []string
	backtrackV2("", 0, 0, n, &result)

	return result
}

func backtrackV2(s string, openCount int, closeCount int, max int, result *[]string) {
	if len(s) == 2*max {
		*result = append(*result, s)
		return
	}
	if openCount < max {
		backtrackV2(s+"(", openCount+1, closeCount, max, result)
	}
	if closeCount < openCount {
		backtrackV2(s+")", openCount, closeCount+1, max, result)
	}
}

func GenerateParenthesisV3(n int) []string {
	var result []string
	if n <= 0 {
		return []string{""}
	}

	backtrackV3("", 0, 0, n, &result)
	return result
}

func backtrackV3(s string, openCount int, closeCount int, max int, result *[]string) {
	if len(s) == 2*max {
		*result = append(*result, s)
		return
	}
	if openCount < max {
		backtrackV3(s+"(", openCount+1, closeCount, max, result)
	}
	if closeCount < openCount {
		backtrackV3(s+")", openCount, closeCount+1, max, result)
	}
}
