package arrays

// https://leetcode.com/problems/fruits-into-baskets-ii/

func numOfUnplacedFruits(fruits []int, baskets []int) int {
	unplaced := 0
	for _, fruit := range fruits {
		unplaced++

		for idx := range baskets {
			if fruit <= baskets[idx] {
				baskets[idx] = 0
				unplaced--

				break
			}
		}
	}

	return unplaced
}
