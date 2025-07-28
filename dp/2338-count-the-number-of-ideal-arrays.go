package dp

// https://leetcode.com/problems/count-the-number-of-ideal-arrays/

func idealArrays(n int, maxValue int) int {
	const MOD, maxK int = 1e9 + 7, 14

	comb := make([][]int, n)
	for i := 0; i < n; i++ {
		comb[i] = make([]int, maxK+1)
	}

	comb[0][0] = 1
	for i := 1; i < n; i++ {
		comb[i][0] = 1
		for j := 1; j <= maxK && j <= i; j++ {
			comb[i][j] = (comb[i-1][j-1] + comb[i-1][j]) % MOD
		}
	}

	memo := make([][]*int, maxValue+1)
	for i := range memo {
		memo[i] = make([]*int, maxK+1)
	}

	var dp func(num int, k int) int

	dp = func(num int, k int) int {
		if k > maxK {
			return 0
		}

		if memo[num][k] != nil {
			return *memo[num][k]
		}
		// num is in the last pos, coms[k] is the
		// combination of n-1 previous positions
		// k is the number of items to be chosen
		res := comb[n-1][k-1]
		for mult := num * 2; mult <= maxValue; mult += num {
			res = (res + dp(mult, k+1)) % MOD
		}

		memo[num][k] = &res

		return res
	}
	total := 0

	for num := 1; num <= maxValue; num++ {
		total = (total + dp(num, 1)) % MOD
	}

	return total
}
