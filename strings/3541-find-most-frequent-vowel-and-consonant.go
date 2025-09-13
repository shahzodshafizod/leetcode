package strings

// https://leetcode.com/problems/find-most-frequent-vowel-and-consonant/

func maxFreqSum(s string) int {
	var freq [26]int
	for _, c := range s {
		freq[int(c-'a')]++
	}

	vowels := map[rune]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}

	var vmx, cmx int

	for i, f := range freq {
		if vowels[rune('a'+i)] {
			vmx = max(vmx, f)
		} else {
			cmx = max(cmx, f)
		}
	}

	return vmx + cmx
}
