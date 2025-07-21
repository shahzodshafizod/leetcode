package bits

// https://leetcode.com/problems/count-the-number-of-consistent-strings/

func countConsistentStrings(allowed string, words []string) int {
	mask := 0
	for _, c := range allowed {
		mask |= 1 << int(c-'a')
	}
	count := 0
	for _, word := range words {
		count++
		for _, c := range word {
			if (1<<int(c-'a'))&mask == 0 {
				count--
				break
			}
		}
	}
	return count
}
