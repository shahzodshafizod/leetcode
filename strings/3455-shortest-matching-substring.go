package strings

import (
	"math"
	"sort"
	"strings"
)

// https://leetcode.com/problems/shortest-matching-substring/

// Approach: Binary Search with Pattern Matching
// Find all occurrences of prefix, middle, suffix, then use binary search
// to find the shortest valid combination
// Time: O(n log n) where n = len(s)
// Space: O(n) for storing occurrences
//
//nolint:gocyclo,revive,cyclop
func minValidSubstring(s string, p string) int {
	// Split pattern by '*'
	parts := strings.Split(p, "*")
	if len(parts) != 3 {
		return -1
	}

	prefix, middle, suffix := parts[0], parts[1], parts[2]

	// Edge case: empty pattern parts
	if prefix == "" && middle == "" && suffix == "" {
		return 0
	}

	lenPre := len(prefix)
	lenMid := len(middle)
	lenSuf := len(suffix)

	// Find all occurrences of a substring
	findAllOccurrences := func(part string) []int {
		if part == "" {
			return []int{}
		}

		positions := []int{}
		pos := 0

		for {
			idx := strings.Index(s[pos:], part)
			if idx == -1 {
				break
			}

			positions = append(positions, pos+idx)
			pos = pos + idx + 1
		}

		return positions
	}

	listP := findAllOccurrences(prefix)
	listM := findAllOccurrences(middle)
	listS := findAllOccurrences(suffix)

	minLen := math.MaxInt32

	// Handle all 8 cases based on which parts are empty
	if prefix != "" && middle != "" && suffix != "" {
		// All three parts present
		for _, pStart := range listP {
			pEnd := pStart + lenPre
			// Find first middle at or after prefix ends
			mIdx := sort.SearchInts(listM, pEnd)
			if mIdx >= len(listM) {
				continue
			}

			mStart := listM[mIdx]
			mEnd := mStart + lenMid
			// Find first suffix at or after middle ends
			sIdx := sort.SearchInts(listS, mEnd)
			if sIdx >= len(listS) {
				continue
			}

			sStart := listS[sIdx]
			sEnd := sStart + lenSuf

			length := sEnd - pStart
			if length < minLen {
				minLen = length
			}
		}
	} else if prefix == "" && middle != "" && suffix != "" {
		// Only middle and suffix
		for _, mStart := range listM {
			mEnd := mStart + lenMid

			sIdx := sort.SearchInts(listS, mEnd)
			if sIdx >= len(listS) {
				continue
			}

			sStart := listS[sIdx]
			sEnd := sStart + lenSuf

			length := sEnd - mStart
			if length < minLen {
				minLen = length
			}
		}
	} else if prefix != "" && middle == "" && suffix != "" {
		// Only prefix and suffix
		for _, pStart := range listP {
			pEnd := pStart + lenPre

			sIdx := sort.SearchInts(listS, pEnd)
			if sIdx >= len(listS) {
				continue
			}

			sStart := listS[sIdx]
			sEnd := sStart + lenSuf

			length := sEnd - pStart
			if length < minLen {
				minLen = length
			}
		}
	} else if prefix != "" && middle != "" && suffix == "" {
		// Only prefix and middle
		for _, pStart := range listP {
			pEnd := pStart + lenPre

			mIdx := sort.SearchInts(listM, pEnd)
			if mIdx >= len(listM) {
				continue
			}

			mStart := listM[mIdx]
			mEnd := mStart + lenMid

			length := mEnd - pStart
			if length < minLen {
				minLen = length
			}
		}
	} else if prefix != "" && middle == "" && suffix == "" {
		// Only prefix (pattern: "prefix**")
		if len(listP) > 0 {
			minLen = lenPre
		}
	} else if prefix == "" && middle != "" && suffix == "" {
		// Only middle (pattern: "*middle*")
		if len(listM) > 0 {
			minLen = lenMid
		}
	} else if prefix == "" && middle == "" && suffix != "" {
		// Only suffix (pattern: "**suffix")
		if len(listS) > 0 {
			minLen = lenSuf
		}
	}

	if minLen == math.MaxInt32 {
		return -1
	}

	return minLen
}

// // Approach #2: Pattern matching without binary search (TLE)
// // Time: O(n^2) where n = len(s)
// // Space: O(1)
// func minValidSubstring(s string, p string) int {
// 	indexFrom := func(s, substr string, start int) int {
// 		if start >= len(s) {
// 			return -1
// 		}
//
// 		for i := start; i <= len(s)-len(substr); i++ {
// 			if s[i:i+len(substr)] == substr {
// 				return i
// 			}
// 		}
//
// 		return -1
// 	}
//
// 	// Split pattern by '*'
// 	parts := make([]string, 0, 3)
// 	current := ""
//
// 	for _, ch := range p {
// 		if ch == '*' {
// 			parts = append(parts, current)
// 			current = ""
// 		} else {
// 			current += string(ch)
// 		}
// 	}
//
// 	parts = append(parts, current)
//
// 	if len(parts) != 3 {
// 		return -1
// 	}
//
// 	prefix, middle, suffix := parts[0], parts[1], parts[2]
//
// 	// Edge case: empty pattern parts (like "**")
// 	if prefix == "" && middle == "" && suffix == "" {
// 		return 0
// 	}
//
// 	n := len(s)
// 	minLen := math.MaxInt32
//
// 	// Try all possible starting positions
// 	for start := range n {
// 		// Match prefix
// 		pos := start
// 		if prefix != "" {
// 			if start+len(prefix) > n {
// 				continue
// 			}
//
// 			if s[start:start+len(prefix)] != prefix {
// 				continue
// 			}
//
// 			pos = start + len(prefix)
// 		}
//
// 		// Try to find middle starting from pos
// 		if middle != "" {
// 			midPos := indexFrom(s, middle, pos)
// 			if midPos == -1 {
// 				continue
// 			}
//
// 			pos = midPos + len(middle)
// 		}
//
// 		// Try to find suffix starting from pos
// 		var end int
//
// 		if suffix != "" {
// 			sufPos := indexFrom(s, suffix, pos)
// 			if sufPos == -1 {
// 				continue
// 			}
//
// 			end = sufPos + len(suffix)
// 		} else {
// 			end = pos
// 		}
//
// 		// Found a match
// 		length := end - start
// 		if length < minLen {
// 			minLen = length
// 		}
// 	}
//
// 	if minLen == math.MaxInt32 {
// 		return -1
// 	}
//
// 	return minLen
// }
