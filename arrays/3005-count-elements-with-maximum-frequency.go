package arrays

// https://leetcode.com/problems/count-elements-with-maximum-frequency/

func maxFrequencyElements(nums []int) int {
	var count [101]int

	total, maxCount := 0, 0

	for _, num := range nums {
		count[num]++

		switch {
		case count[num] > maxCount:
			maxCount = count[num]
			total = 1
		case count[num] == maxCount:
			total++
		}
	}

	return total * maxCount
}
