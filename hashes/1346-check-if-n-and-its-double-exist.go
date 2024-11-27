package hashes

// https://leetcode.com/problems/check-if-n-and-its-double-exist/

func checkIfExist(arr []int) bool {
	var seen = make(map[int]bool)
	for _, num := range arr {
		if seen[2*num] || num&1 == 0 && seen[num/2] {
			return true
		}
		seen[num] = true
	}
	return false
}
