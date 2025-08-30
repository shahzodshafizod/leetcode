package strings

// https://leetcode.com/problems/find-the-k-th-character-in-string-game-i/

func kthCharacter(k int) byte {
	word := make([]int, 0, k*2)
	word = append(word, 0)

	length := 1
	for ; length < k; length *= 2 {
		for i := range length {
			word = append(word, (word[i]+1)%26)
		}
	}

	return byte('a' + word[k-1])
}
