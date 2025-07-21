package hashes

// https://leetcode.com/problems/find-the-number-of-distinct-colors-among-the-balls/

func queryResults(_ int, queries [][]int) []int {
	balls := make(map[int]int)  // ball -> color
	colors := make(map[int]int) // color -> count
	result := make([]int, len(queries))
	var x, y, color int
	var found bool
	count := 0
	for idx := range queries {
		x, y = queries[idx][0], queries[idx][1]
		color, found = balls[x]
		if found {
			colors[color]--
			if colors[color] == 0 {
				count--
			}
		}
		balls[x] = y
		colors[y]++
		if colors[y] == 1 {
			count++
		}
		result[idx] = count
	}
	return result
}
