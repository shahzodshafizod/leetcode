package strings

// https://leetcode.com/problems/strong-password-checker/

// Approach: Optimized Greedy Algorithm
// We analyze the problem by length categories:
// 1. length < 6: need insertions (can fix types + repeats)
// 2. length > 20: need deletions (strategic to help with repeats)
// 3. 6 <= length <= 20: use replacements
// Time: O(n) - single pass to count repeats, process deletions/replacements
// Space: O(1) - only tracking counts and sequences
func strongPasswordChecker(password string) int {
	n := len(password)

	// Check character type requirements
	hasLower, hasUpper, hasDigit := false, false, false

	for _, c := range password {
		if c >= 'a' && c <= 'z' {
			hasLower = true
		} else if c >= 'A' && c <= 'Z' {
			hasUpper = true
		} else if c >= '0' && c <= '9' {
			hasDigit = true
		}
	}

	missingTypes := 0
	if !hasLower {
		missingTypes++
	}

	if !hasUpper {
		missingTypes++
	}

	if !hasDigit {
		missingTypes++
	}

	// Count repeating sequences
	// We track sequences by their length % 3 (important for optimization)
	totalReplacements := 0 // replacements needed for repeating sequences
	oneMod := 0            // sequences where len % 3 == 0 (one delete helps most)
	twoMod := 0            // sequences where len % 3 == 1 (two deletes help most)

	i := 2
	for i < n {
		if i >= 2 && password[i] == password[i-1] && password[i] == password[i-2] {
			length := 2
			for i < n && password[i] == password[i-1] {
				length++
				i++
			}

			// Each sequence of length L needs L//3 replacements
			totalReplacements += length / 3

			// Track by modulo for deletion optimization
			switch length % 3 {
			case 0:
				oneMod++
			case 1:
				twoMod++
			default:
				// length % 3 == 2, no special tracking needed
			}
		} else {
			i++
		}
	}

	// Case 1: Too short (length < 6)
	if n < 6 {
		// Need to insert characters
		insertionsNeeded := 6 - n
		// Insertions can fix both missing types and some repeats
		return max(missingTypes, insertionsNeeded)
	}

	// Case 2: Too long (length > 20)
	if n > 20 {
		deletionsNeeded := n - 20
		remainingDeletions := deletionsNeeded

		// Optimize deletions to reduce replacements needed
		// Priority: sequences with len%3==0 need only 1 deletion to reduce replacements
		use := min(remainingDeletions, oneMod)
		totalReplacements -= use
		remainingDeletions -= use

		// Next: sequences with len%3==1 need 2 deletions
		use = min(remainingDeletions, twoMod*2)
		totalReplacements -= use / 2
		remainingDeletions -= use

		// Remaining deletions: each 3 deletions saves 1 replacement
		totalReplacements -= remainingDeletions / 3

		return deletionsNeeded + max(missingTypes, totalReplacements)
	}

	// Case 3: Length is fine (6 <= length <= 20)
	// Only need replacements
	// Each replacement can fix one repeat AND potentially one missing type
	return max(missingTypes, totalReplacements)
}
