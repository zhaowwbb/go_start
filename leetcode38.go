package main

import (
	"strconv"
	"strings"
)

// Package-level cache initialization. Constraints state 1 <= n <= 30.
var memo [31]string

// CountAndSay generates the nth element of the count-and-say sequence.
func CountAndSay(n int) string {
	// Base case initialization
	if memo[1] == "" {
		memo[1] = "1"
	}

	// Fill the cache up to n if not already computed
	for i := 2; i <= n; i++ {
		if memo[i] == "" {
			memo[i] = getNextSequence(memo[i-1])
		}
	}

	return memo[n]
}

// getNextSequence performs Run-Length Encoding (RLE) on a string
func getNextSequence(s string) string {
	var sb strings.Builder
	length := len(s)

	i := 0
	for i < length {
		currentChar := s[i]
		count := 0

		// Count consecutive identical characters
		for i < length && s[i] == currentChar {
			count++
			i++
		}

		// Write frequency followed by the character itself
		sb.WriteString(strconv.Itoa(count))
		sb.WriteByte(currentChar)
	}

	return sb.String()
}

func CountAndSayV2(n int) string {
	if n == 1 {
		return "1"
	}
	memo[1] = "1"
	for i := 2; i <= n; i++ {
		memo[i] = getNextSequenceV2(memo[i-1])
	}
	return memo[n]
}

func getNextSequenceV2(s string) string {
	var sb strings.Builder
	i := 0
	for i < len(s) {
		count := 0
		currentChar := s[i]
		for i < len(s) && currentChar == s[i] {
			count++
			i++
		}
		sb.WriteString(strconv.Itoa(count))
		sb.WriteByte(currentChar)
	}
	return sb.String()
}

func CountAndSayV3(n int) string {
	if n == 1 {
		return "1"
	}
	memo[1] = "1"
	for i := 2; i <= n; i++ {
		memo[i] = getNextSequenceV3(memo[i-1])
	}
	return memo[n]
}

func getNextSequenceV3(s string) string {
	i := 0
	length := len(s)
	var sb strings.Builder
	for i < length {
		currentChar := s[i]
		count := 0
		for i < length && currentChar == s[i] {
			count++
			i++
		}
		sb.WriteString(strconv.Itoa(count))
		sb.WriteByte(currentChar)
	}
	return sb.String()
}
