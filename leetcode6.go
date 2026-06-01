package main

import "strings"

// Convert transforms a string into a zigzag pattern using a fixed number of rows,
// then reads the characters row by row.
// Approach: Simulate Row-by-Row Traversal using String Builders
// Time Complexity: O(n), Space Complexity: O(n)
func Convert(s string, numRows int) string {
	// Base Case: If there's only 1 row or the rows exceed string length,
	// the zigzag layout remains identical to the original sequence.
	if numRows == 1 || numRows >= len(s) {
		return s
	}

	// Initialize a slice of Builders for each row bucket
	rows := make([]strings.Builder, numRows)
	currRow := 0
	goingDown := false

	// Distribute characters into their corresponding row bucket
	for i := 0; i < len(s); i++ {
		rows[currRow].WriteByte(s[i])

		// Switch direction whenever we reach the top (0) or bottom (numRows - 1)
		if currRow == 0 || currRow == numRows-1 {
			goingDown = !goingDown
		}

		// Move index up or down depending on the direction
		if goingDown {
			currRow++
		} else {
			currRow--
		}
	}

	// Combine all rows into a single string builder
	var result strings.Builder
	for i := 0; i < numRows; i++ {
		result.WriteString(rows[i].String())
	}

	return result.String()
}

func ConvertV2(s string, numRows int) string {
	if len(s) <= numRows || numRows == 1 {
		return s
	}
	currentRow := 0
	goDown := false
	rows := make([]strings.Builder, numRows)

	for i := 0; i < len(s); i++ {
		rows[currentRow].WriteByte(s[i])

		if currentRow == 0 || currentRow == numRows-1 {
			goDown = !goDown
		}

		if goDown {
			currentRow++
		} else {
			currentRow--
		}
	}

	var result strings.Builder
	for i := 0; i < numRows; i++ {
		result.WriteString(rows[i].String())
	}
	return result.String()
}

func ConvertV3(s string, numRows int) string {
	if len(s) <= numRows || numRows == 1 {
		return s
	}

	currentRow := 0
	goDown := false
	rows := make([]strings.Builder, numRows)

	for i := 0; i < len(s); i++ {
		rows[currentRow].WriteByte(s[i])
		if currentRow == 0 || currentRow == numRows-1 {
			goDown = !goDown
		}

		if goDown {
			currentRow++
		} else {
			currentRow--
		}
	}

	var result strings.Builder
	for i := 0; i < numRows; i++ {
		result.WriteString(rows[i].String())
	}
	return result.String()
}
