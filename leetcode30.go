package main

// findSubstring finds the starting indices of all concatenated substrings
func findSubstring(s string, words []string) []int {
	result := []int{}
	if len(s) == 0 || len(words) == 0 {
		return result
	}

	wordLen := len(words[0])
	wordCount := len(words)
	totalLen := wordLen * wordCount

	if len(s) < totalLen {
		return result
	}

	// Build the frequency map for the words array
	wordFreq := make(map[string]int)
	for _, word := range words {
		wordFreq[word]++
	}

	// Run the sliding window wordLen times for different offsets
	for i := 0; i < wordLen; i++ {
		left := i
		right := i
		currentFreq := make(map[string]int)
		count := 0 // Number of valid words currently in the window

		for right+wordLen <= len(s) {
			// Extract the next word candidate
			word := s[right : right+wordLen]
			right += wordLen

			if targetCount, exists := wordFreq[word]; exists {
				currentFreq[word]++
				count++

				// If we have more occurrences of 'word' than needed, shrink from the left
				for currentFreq[word] > targetCount {
					leftWord := s[left : left+wordLen]
					currentFreq[leftWord]--
					count--
					left += wordLen
				}

				// If the window size matches the total number of words, we found a match
				if count == wordCount {
					result = append(result, left)
				}
			} else {
				// The word is not part of the dictionary; reset the window
				currentFreq = make(map[string]int)
				count = 0
				left = right
			}
		}
	}

	return result
}

func findSubstringV2(s string, words []string) []int {
	result := []int{}
	if len(s) == 0 || len(words) == 0 {
		return result
	}

	wordLen := len(words[0])
	wordCount := len(words)
	subStrLen := wordLen * wordCount
	if len(s) < subStrLen {
		return result
	}

	wordFreq := make(map[string]int)

	for _, word := range words {
		wordFreq[word]++
	}

	for i := 0; i < wordLen; i++ {
		left, right := i, i
		currentFreq := make(map[string]int)
		count := 0
		for right+wordLen <= len(s) {
			word := s[right : right+wordLen]
			right += wordLen
			if targetCount, exists := wordFreq[word]; exists {
				currentFreq[word]++
				count++

				for currentFreq[word] > targetCount {
					leftWord := s[left : left+wordLen]
					currentFreq[leftWord]--
					count--
					left += wordLen
				}
				if count == wordCount {
					result = append(result, left)
				}

			} else {
				currentFreq = make(map[string]int)
				count = 0
				left = right
			}
		}
	}

	return result
}

func findSubstringV3(s string, words []string) []int {
	result := []int{}
	if len(s) == 0 || len(words) == 0 {
		return result
	}
	wordLen := len(words[0])
	wordCount := len(words)
	totalLen := wordLen * wordCount
	if len(s) < totalLen {
		return result
	}

	wordFreq := make(map[string]int)
	for _, word := range words {
		wordFreq[word]++
	}

	for i := 0; i < wordLen; i++ {
		left := i
		right := i
		count := 0
		currentFreq := make(map[string]int)
		for right+wordLen <= len(s) {
			word := s[right : right+wordLen]
			right += wordLen
			if totalCount, exists := wordFreq[word]; exists {
				currentFreq[word]++
				count++
				for currentFreq[word] > totalCount {
					leftWord := s[left : left+wordLen]
					currentFreq[leftWord]--
					count--
					left += wordLen
				}
				if count == wordCount {
					result = append(result, left)
				}
			} else {
				currentFreq = make(map[string]int)
				count = 0
				left = right
			}
		}
	}

	return result
}
