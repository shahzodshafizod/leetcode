package backtracking

// https://leetcode.com/problems/letter-tile-possibilities/

func numTilePossibilities(tiles string) int {
	count := make(map[byte]int)
	for idx := range tiles {
		count[tiles[idx]]++
	}

	var backtrack func() int

	backtrack = func() int {
		res := 0

		for key, cnt := range count {
			if cnt == 0 {
				continue
			}

			count[key]--
			res += 1 + backtrack()
			count[key]++
		}

		return res
	}

	return backtrack()
}
