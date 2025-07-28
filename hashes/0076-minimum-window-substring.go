package hashes

// https://leetcode.com/problems/minimum-window-substring/

func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}

	hash := make(map[rune]int)

	table := make(map[rune]int)
	for _, r := range t {
		table[r]++
	}

	count := len(table)
	minLength := 0

	var window string

	leftIndex := 0

	for index, r := range s {
		if miqdor := table[r]; miqdor > 0 {
			hash[r]++
			if hash[r] == miqdor {
				count--
			}

			if count == 0 {
				rightIndex := index
				for ; leftIndex < rightIndex; leftIndex++ {
					key := rune(s[leftIndex])
					if count, exists := hash[key]; exists {
						if count == table[key] {
							break
						}

						hash[key]--
					}
				}

				if length := rightIndex - leftIndex + 1; minLength == 0 || length < minLength {
					minLength = length
					window = s[leftIndex : rightIndex+1]
				}

				hash[rune(s[leftIndex])]--
				leftIndex++
				count = 1
			}
		}
	}

	return window
}
