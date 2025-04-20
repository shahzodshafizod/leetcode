package greedy

// https://leetcode.com/problems/rabbits-in-forest/

func numRabbits(answers []int) int {
	var count = make(map[int]int)
	for _, others := range answers {
		count[others]++
	}
	var result = 0
	for others, total := range count {
		result += (total + others) / (others + 1) * (others + 1)
	}
	return result
}
