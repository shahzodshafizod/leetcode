package strings

// https://leetcode.com/problems/custom-sort-string/

// time: O(2n) = O(n)
// space: O(26+n) = O(n)
func customSortString(order string, s string) string {
	var counts [26]int
	var slen = len(s)
	for idx := 0; idx < slen; idx++ {
		counts[s[idx]-'a']++
	}
	var permuted = make([]byte, 0, slen)
	var count, index int
	for idx, olen := 0, len(order); idx < olen; idx++ {
		index = int(order[idx] - 'a')
		count = counts[index]
		for count > 0 {
			count--
			permuted = append(permuted, order[idx])
		}
		counts[index] = 0
	}
	var char byte
	for idx, count := range counts {
		char = byte(idx + 'a')
		for count > 0 {
			count--
			permuted = append(permuted, char)
		}
	}
	return string(permuted)
}
