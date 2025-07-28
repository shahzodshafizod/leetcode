package arrays

// https://leetcode.com/problems/find-if-digit-game-can-be-won/

func canAliceWin(nums []int) bool {
	singleSum, doubleSum := 0, 0

	for _, num := range nums {
		if num < 10 {
			singleSum += num
		} else {
			doubleSum += num
		}
	}

	return singleSum != doubleSum
}
