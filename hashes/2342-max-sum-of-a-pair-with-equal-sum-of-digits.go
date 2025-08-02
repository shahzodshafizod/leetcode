package hashes

// https://leetcode.com/problems/max-sum-of-a-pair-with-equal-sum-of-digits/

func maximumSum(nums []int) int {
	pairs := make(map[int][2]int)

	var tmp, sum int

	maxSum := -1

	for _, num := range nums {
		tmp, sum = num, 0
		for tmp > 0 {
			sum += tmp % 10
			tmp /= 10
		}

		if _, ok := pairs[sum]; !ok {
			pairs[sum] = [2]int{num, 0}

			continue
		}

		if num > pairs[sum][0] {
			pairs[sum] = [2]int{num, pairs[sum][0]}
		} else if num > pairs[sum][1] {
			pairs[sum] = [2]int{pairs[sum][0], num}
		}

		maxSum = max(maxSum, pairs[sum][0]+pairs[sum][1])
	}

	return maxSum
}
