package dp

// https://leetcode.com/problems/count-number-of-balanced-permutations/

func countBalancedPermutations(num string) int {
	var cnt [10]int
	sum, maxFreq := 0, 0
	var digit int
	for _, c := range num {
		digit = int(c - '0')
		cnt[digit]++
		maxFreq = max(maxFreq, cnt[digit])
		sum += digit
	}
	if sum&1 == 1 {
		return 0
	}

	var postCnt [11]int
	for i := 9; i >= 0; i-- {
		postCnt[i] = cnt[i] + postCnt[i+1]
	}

	const MOD int = 1e9 + 7
	n := len(num)
	oddCnt := (n + 1) / 2
	halfSum := sum >> 1

	// var comb = func(n int, k int) int {
	// 	if k > n {
	// 		return 0
	// 	}
	// 	res := 1
	// 	for i := 0; i < k; i++ {
	// 		res = res * (n - i) / (i + 1)
	// 	}
	// 	return res % MOD
	// }
	comb := make([][]int, oddCnt+1)
	for i := 0; i <= oddCnt; i++ {
		comb[i] = make([]int, maxFreq+1)
	}
	comb[0][0] = 1
	for i := 1; i <= oddCnt; i++ {
		comb[i][0] = 1
		for j := 1; j <= maxFreq; j++ {
			comb[i][j] = (comb[i-1][j-1] + comb[i-1][j]) % MOD
		}
	}

	memo := make([][][10]*int, oddCnt+1)
	for idx := 0; idx <= oddCnt; idx++ {
		memo[idx] = make([][10]*int, halfSum+1)
	}

	var dfs func(digit int, ocnt int, balance int) int
	dfs = func(digit int, ocnt int, balance int) int {
		ecnt := n - ocnt - postCnt[digit+1]
		if ocnt == 0 && ecnt == 0 && balance == 0 {
			return 1
		}
		if digit < 0 || ocnt < 0 || ecnt < 0 || balance < 0 {
			return 0
		}
		if memo[ocnt][balance][digit] != nil {
			return *memo[ocnt][balance][digit]
		}
		res := 0
		var nxt, ways int
		for freq := 0; freq <= cnt[digit]; freq++ {
			nxt = dfs(digit-1, ocnt-freq, balance-digit*freq)
			ways = comb[ocnt][freq] * comb[ecnt][cnt[digit]-freq] % MOD
			res = (res + ways*nxt) % MOD
		}
		memo[ocnt][balance][digit] = &res
		return res
	}

	return dfs(9, oddCnt, halfSum)
}
