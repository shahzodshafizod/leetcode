package tries

// https://leetcode.com/problems/find-unique-binary-string/

// Approach #2: Cantor's Diagonal Argument
// Time: O(n)
// Space: O(1)
func findDifferentBinaryString(nums []string) string {
	n := len(nums)
	num := make([]byte, n)

	for idx := 0; idx < n; idx++ {
		if nums[idx][idx] == '0' {
			num[idx] = '1'
		} else {
			num[idx] = '0'
		}
	}

	return string(num)
}

// // Approach #1: Trie
// // Time: O(nn)
// // Space: O(nn)
// func findDifferentBinaryString(nums []string) string {
// 	type TrieNode struct {
// 		children map[rune]*TrieNode
// 		count    int
// 	}

// 	var NewTrieNode = func() *TrieNode {
// 		return &TrieNode{
// 			children: make(map[rune]*TrieNode),
// 			count:    0,
// 		}
// 	}
// 	var root = NewTrieNode()
// 	var insert = func(s string) {
// 		var curr = root
// 		for _, c := range s {
// 			if curr.children[c] == nil {
// 				curr.children[c] = NewTrieNode()
// 			}
// 			curr = curr.children[c]
// 			curr.count++
// 		}
// 	}
// 	var n = len(nums)
// 	for idx := 0; idx < n; idx++ {
// 		insert(nums[idx])
// 	}
// 	var num = make([]rune, n)
// 	var template = "01"
// 	var partitionLen = 1 << n
// 	var curr = root
// 	for idx := 0; idx < n; idx++ {
// 		partitionLen /= 2
// 		for _, c := range template {
// 			if curr.children[c] == nil {
// 				curr.children[c] = NewTrieNode()
// 			}
// 			if curr.children[c].count < partitionLen {
// 				curr = curr.children[c]
// 				num[idx] = c
// 				break
// 			}
// 		}
// 	}
// 	return string(num)
// }
