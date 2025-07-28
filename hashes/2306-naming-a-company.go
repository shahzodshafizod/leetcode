package hashes

// https://leetcode.com/problems/naming-a-company/

// Approach: Hash Set
// Time: O(N)
// Space: O(N)
func distinctNames(ideas []string) int64 {
	suffixes := make(map[string][]int)

	var prefix int

	var suffix string

	var counts [26]int64

	var same [26][26]int64

	for idx := range ideas {
		prefix = int(ideas[idx][0] - 'a')
		suffix = ideas[idx][1:]
		counts[prefix]++

		for _, c := range suffixes[suffix] {
			same[prefix][c]++
			same[c][prefix]++
		}

		suffixes[suffix] = append(suffixes[suffix], prefix)
	}

	var count int64 = 0

	for a := 0; a < 26; a++ {
		for b := a + 1; b < 26; b++ {
			count += (counts[a] - same[a][b]) * (counts[b] - same[b][a])
		}
	}

	return count * 2
}

// // Approach: Tries
// // Time: O(NxW), N=len(ideas), W=max(len(ideas[i]))
// // Space: O(N)
// func distinctNames(ideas []string) int64 {
// 	type Node struct {
// 		children map[int]*Node
// 		start    bool
// 	}
// 	var counts [26]int64
// 	var same [26][26]int64
// 	var key int
// 	var root = &Node{children: make(map[int]*Node)}
// 	for idx := range ideas {
// 		curr := root
// 		prev := curr
// 		for c := len(ideas[idx]) - 1; c >= 0; c-- {
// 			key = int(ideas[idx][c] - 'a')
// 			if curr.children[key] == nil {
// 				curr.children[key] = &Node{children: make(map[int]*Node)}
// 			}
// 			prev = curr
// 			curr = curr.children[key]
// 		}
// 		counts[key]++
// 		for c, next := range prev.children {
// 			if next != nil && next.start {
// 				same[key][c]++
// 				same[c][key]++
// 			}
// 		}
// 		curr.start = true
// 	}
// 	var count int64 = 0
// 	for a := 0; a < 26; a++ {
// 		for b := a + 1; b < 26; b++ {
// 			count += (counts[a] - same[a][b]) * (counts[b] - same[b][a]) // * 2
// 		}
// 	}
// 	return count * 2
// }
