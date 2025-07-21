package hashes

// https://leetcode.com/problems/tuple-with-same-product/

func tupleSameProduct(nums []int) int {
	squares := make(map[int]int)
	for i, n := 0, len(nums); i < n; i++ {
		for j := i + 1; j < n; j++ {
			squares[nums[i]*nums[j]]++
		}
	}
	count := 0
	for _, cnt := range squares {
		count += 8 * cnt * (cnt - 1)
	}
	return count / 2
}

/*
1: a * b = c * d
2: b * a = c * d
3: a * b = d * c
4: b * a = d * c
5: c * d = a * b
6: c * d = b * a
7: d * c = a * b
8: d * c = b * a
*/
