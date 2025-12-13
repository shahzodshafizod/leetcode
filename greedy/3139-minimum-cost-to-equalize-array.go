package greedy

// https://leetcode.com/problems/minimum-cost-to-equalize-array/

// Approach: Greedy with 4 Cases Analysis (Optimized)
// Based on lee215's solution: analyze 4 distinct cases
// Case 1: cost1*2 <= cost2 or n <= 2 → always use cost1
// Case 2: Use cost2 as much as possible up to max_val
// Case 3: Go beyond max_val to reduce op1 usage
// Case 4: Go even further to eliminate op1 completely
// Time: O(n), Space: O(1)
func minCostToEqualizeArray(nums []int, cost1 int, cost2 int) int {
	mod := int(1e9 + 7)
	n := len(nums)

	if n == 1 {
		return 0
	}

	ma, mi := nums[0], nums[0]
	sum := 0

	for _, num := range nums {
		if num > ma {
			ma = num
		}

		if num < mi {
			mi = num
		}

		sum += num
	}

	total := ma*n - sum // Total increments needed to reach max

	// Case 1: cost1 is always better, or array too small for pairing
	if cost1*2 <= cost2 || n <= 2 {
		return (total * cost1) % mod
	}

	// Case 2: Increment to max(nums), use cost2 as much as possible
	op1 := (ma-mi)*2 - total
	if op1 < 0 {
		op1 = 0
	}

	op2 := total - op1
	res := (op1+op2%2)*cost1 + op2/2*cost2

	// Case 3: Increment beyond max to reduce op1
	// Each round increases total by n, but reduces op1 by (n-2)
	if op1 > 0 {
		rounds := op1 / (n - 2)
		total += rounds * n
		op1 %= (n - 2)
		op2 = total - op1

		cost := (op1+op2%2)*cost1 + op2/2*cost2
		if cost < res {
			res = cost
		}
	}

	// Case 4: Increment even further to make total even (eliminate all op1)
	for range 2 {
		total += n

		cost := total%2*cost1 + total/2*cost2
		if cost < res {
			res = cost
		}
	}

	return res % mod
}
