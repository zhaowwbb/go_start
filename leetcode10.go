package main

// IsMatch implements regular expression matching with support for '.' and '*' using Dynamic Programming.
func IsMatch(s string, p string) bool {
	m := len(s)
	n := len(p)

	// dp[i][j] will be true if s[0..i-1] matches p[0..j-1]
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}

	// Base case: empty string matches empty pattern
	dp[0][0] = true

	// Deals with patterns like a*, a*b*, a*b*c* matching an empty string s
	for j := 1; j <= n; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}

	// Fill the DP table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			currentP := p[j-1]
			currentS := s[i-1]

			if currentP == currentS || currentP == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else if currentP == '*' {
				// Case 1: Count '*' as 0 occurrences of the preceding element
				dp[i][j] = dp[i][j-2]

				// Case 2: Count '*' as 1 or more occurrences
				// This is only valid if the preceding character in p matches currentS or is '.'
				predecessorP := p[j-2]
				if predecessorP == currentS || predecessorP == '.' {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			}
		}
	}

	return dp[m][n]
}

func IsMatchV2(s string, p string) bool {
	m := len(s)
	n := len(p)
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for j := 2; j <= n; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			curS := s[i-1]
			curP := p[j-1]
			if curS == curP || curP == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else if curP == '*' {
				dp[i][j] = dp[i][j-2]

				preP := p[j-2]
				if preP == curS || preP == '.' {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			}
		}
	}

	return dp[m][n]
}

func IsMatchV3(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for j := 2; j <= n; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			curS := s[i-1]
			curP := p[j-1]
			if curS == curP || curP == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else if curP == '*' {
				dp[i][j] = dp[i][j-2]
				preP := p[j-2]
				if preP == curS || preP == '.' {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			}
		}
	}

	return dp[m][n]
}
