package trees

// https://leetcode.com/problems/number-of-ways-to-reorder-array-to-get-same-bst/

// Approach #2: Divide and Conquer with Combinatorics
// Time: O(n^2) - computing binomial coefficients
// Space: O(n^2) - storing Pascal's triangle
func numOfWays(nums []int) int {
	const mod int64 = 1e9 + 7

	n := len(nums)

	// Precompute Pascal's triangle for binomial coefficients
	// pascal[i][j] = C(i, j)
	pascal := make([][]int64, n+1)
	for i := range pascal {
		pascal[i] = make([]int64, n+1)

		pascal[i][0] = 1
		for j := 1; j <= i; j++ {
			pascal[i][j] = (pascal[i-1][j-1] + pascal[i-1][j]) % mod
		}
	}

	var countWays func(arr []int) int64

	countWays = func(arr []int) int64 {
		if len(arr) <= 1 {
			return 1
		}

		root := arr[0]
		left := []int{}
		right := []int{}

		for i := 1; i < len(arr); i++ {
			if arr[i] < root {
				left = append(left, arr[i])
			} else {
				right = append(right, arr[i])
			}
		}

		leftWays := countWays(left)
		rightWays := countWays(right)

		// Number of ways to interleave left and right
		// C(len(left) + len(right), len(left))
		ways := pascal[len(left)+len(right)][len(left)]

		return (ways * leftWays % mod) * rightWays % mod
	}

	// Subtract 1 because the original ordering is not counted as a reordering
	result := (countWays(nums) - 1 + mod) % mod

	return int(result)
}
