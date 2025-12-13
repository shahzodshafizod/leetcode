package binarysearch

// https://leetcode.com/problems/kth-smallest-amount-with-single-denomination-combination/

// Approach: Binary Search + Inclusion-Exclusion Principle
// Binary search on the answer. For a given value x, count how many amounts <= x
// can be formed using the coins (using inclusion-exclusion with LCM).
// Time: O(log(k*max_coin) * 2^n) where n = len(coins)
// Space: O(n) for recursion
func findKthSmallest(coins []int, k int) int64 {
	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}

		return a
	}

	lcm := func(a, b int) int {
		return a * b / gcd(a, b)
	}

	// Remove coins that are multiples of other coins (optimization)
	coinsMap := make(map[int]bool)
	for _, coin := range coins {
		coinsMap[coin] = true
	}

	filteredCoins := []int{}

	sortedCoins := make([]int, 0, len(coinsMap))
	for coin := range coinsMap {
		sortedCoins = append(sortedCoins, coin)
	}

	// Simple bubble sort for small arrays
	for i := 0; i < len(sortedCoins); i++ {
		for j := i + 1; j < len(sortedCoins); j++ {
			if sortedCoins[i] > sortedCoins[j] {
				sortedCoins[i], sortedCoins[j] = sortedCoins[j], sortedCoins[i]
			}
		}
	}

	for _, coin := range sortedCoins {
		isMultiple := false

		for _, other := range filteredCoins {
			if coin%other == 0 {
				isMultiple = true

				break
			}
		}

		if !isMultiple {
			filteredCoins = append(filteredCoins, coin)
		}
	}

	coins = filteredCoins
	n := len(coins)

	countAmounts := func(x int) int {
		total := 0

		// Iterate through all subsets
		for mask := 1; mask < (1 << n); mask++ {
			subsetLcm := 1
			bits := 0

			for i := range n {
				if mask&(1<<i) != 0 {
					subsetLcm = lcm(subsetLcm, coins[i])
					bits++
					// Optimization: if LCM becomes too large, skip
					if subsetLcm > x {
						break
					}
				}
			}

			if subsetLcm <= x {
				// Inclusion-exclusion: add if odd number of coins, subtract if even
				if bits%2 == 1 {
					total += x / subsetLcm
				} else {
					total -= x / subsetLcm
				}
			}
		}

		return total
	}

	// Binary search for the kth smallest amount
	maxCoin := coins[0]
	for _, coin := range coins {
		if coin > maxCoin {
			maxCoin = coin
		}
	}

	left, right := 1, k*maxCoin

	for left < right {
		mid := (left + right) / 2
		if countAmounts(mid) < k {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return int64(left)
}
