package greedy

// https://leetcode.com/problems/boats-to-save-people/

// time: O(n)
// space: O(n)
func numRescueBoats(people []int, limit int) int {
	count := make([]int, limit+1)
	for _, weight := range people {
		count[weight]++
	}

	light, heavy := 1, limit
	boats := 0

	for {
		for light <= heavy && count[light] <= 0 {
			light++
		}

		for light <= heavy && count[heavy] <= 0 {
			heavy--
		}

		if light > heavy {
			break
		}

		if light+heavy <= limit {
			count[light]--
		}

		count[heavy]--
		boats++
	}

	return boats
}

// // time: O(n log n)
// // space: O(1)
// func numRescueBoats(people []int, limit int) int {
// 	sort.Ints(people)
// 	var boats = 0
// 	for left, right := 0, len(people)-1; left <= right; {
// 		if people[right]+people[left] <= limit {
// 			left++
// 		}
// 		right--
// 		boats++
// 	}
// 	return boats
// }
