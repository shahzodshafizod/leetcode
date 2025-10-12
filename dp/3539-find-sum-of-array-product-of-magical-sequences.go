package dp

import "math/bits"

// https://leetcode.com/problems/find-sum-of-array-product-of-magical-sequences/

func magicalSum(m int, k int, nums []int) int {
	const mod = int(1e9) + 7

	n := len(nums)

	pows := make([][]int, n)
	for i := range n {
		pows[i] = make([]int, m+1)

		pows[i][0] = 1
		for j := 1; j <= m; j++ {
			pows[i][j] = pows[i][j-1] * nums[i] % mod
		}
	}

	combs := make([][]int, m+1)
	for i := range combs {
		combs[i] = make([]int, m+1)

		combs[i][i], combs[i][0] = 1, 1
		for j := 1; j < i; j++ {
			combs[i][j] = (combs[i-1][j] + combs[i-1][j-1]) % mod
		}
	}

	var memo [50][31][31][31]*int // [n][m+1][m+1][k+1]

	var dp func(i int, m int, k int, mask int) int

	dp = func(i, m, k, mask int) int {
		if m == 0 && bits.OnesCount(uint(mask)) == k {
			return 1
		}

		if i == n {
			return 0
		}

		if memo[i][m][mask][k] != nil {
			return *memo[i][m][mask][k]
		}

		var total, ways, nmask, nk, nxt int
		for freq := range m + 1 {
			ways = combs[m][freq] * pows[i][freq] % mod
			nmask = mask + freq

			nk = k - (nmask & 1)
			if nk < 0 {
				continue
			}

			nxt = dp(i+1, m-freq, nk, nmask>>1)
			total = (total + ways*nxt) % mod
		}

		memo[i][m][mask][k] = &total

		return total
	}

	return dp(0, m, k, 0)
}
