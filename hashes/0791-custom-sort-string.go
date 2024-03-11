package hashes

// https://leetcode.com/problems/custom-sort-string/

// time: O(n)
// space: O(n)
func customSortString(order string, s string) string {
	var hash = make(map[rune]int)
	for _, char := range s {
		hash[char]++
	}
	var permuted = make([]rune, 0, len(s))
	for _, char := range order {
		if count, exists := hash[char]; exists {
			for count > 0 {
				count--
				permuted = append(permuted, char)
			}
			delete(hash, char)
		}
	}
	for char, count := range hash {
		for count > 0 {
			count--
			permuted = append(permuted, rune(char))
		}
	}
	return string(permuted)
}
