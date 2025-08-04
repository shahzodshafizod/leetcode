package slidingwindows

// https://leetcode.com/problems/fruit-into-baskets/

func totalFruit(fruits []int) int {
	a, b := -1, -1
	totallen, lenb := 0, 0
	res := 0

	for _, c := range fruits {
		if c == a || c == b {
			totallen++
		} else {
			totallen = lenb + 1
		}

		if c != b {
			a, b = b, c
			lenb = 0
		}

		lenb++

		res = max(res, totallen)
	}

	return res
}
