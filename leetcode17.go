package main

// Global phone keypad mapping
var phoneMap = map[string]string{
	"2": "abc", "3": "def", "4": "ghi", "5": "jkl",
	"6": "mno", "7": "pqrs", "8": "tuv", "9": "wxyz",
}

// LetterCombinations returns all possible letter combinations that the digits could represent.
// Time Complexity: O(4^N * N) | Space Complexity: O(N) internal recursion stack
func LetterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	var combinations []string
	backtrack(0, "", digits, &combinations)
	return combinations
}

func backtrack(index int, path string, digits string, result *[]string) {
	// Base Case: If the current string matches the input length, save it
	if len(path) == len(digits) {
		*result = append(*result, path)
		return
	}

	// Get letters associated with the current digit
	currentDigit := string(digits[index])
	letters := phoneMap[currentDigit]

	for i := 0; i < len(letters); i++ {
		// Choose a letter, move to the next index, and implicitly backtrack via string copying
		backtrack(index+1, path+string(letters[i]), digits, result)
	}
}

func LetterCombinationsV2(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	var combinations []string
	backtrackV2(0, "", digits, &combinations)
	return combinations
}

func backtrackV2(index int, path string, digits string, result *[]string) {
	if len(path) == len(digits) {
		*result = append(*result, path)
		return
	}
	currentDigit := string(digits[index])
	possibleLetters := phoneMap[currentDigit]
	for i := 0; i < len(possibleLetters); i++ {
		backtrackV2(index+1, path+string(possibleLetters[i]), digits, result)
	}
}

func LetterCombinationsV3(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	var combinations []string
	backtrackV3(0, "", digits, &combinations)
	return combinations
}

func backtrackV3(index int, path string, digits string, result *[]string) {
	if len(path) == len(digits) {
		*result = append(*result, path)
		return
	}

	letterIndex := string(digits[index])
	possibleLetters := phoneMap[letterIndex]
	for i := 0; i < len(possibleLetters); i++ {
		backtrackV3(index+1, path+string(possibleLetters[i]), digits, result)
	}
}
