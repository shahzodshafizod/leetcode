package dp

/*
On avarage, we serve more of soup A than soup B:
	A: (100 + 75 + 50 + 25) / 4 = 62.5
	B: (  0 + 25 + 50 + 75) / 4 = 37.5
As n grows, the probability that A empties first approaches 1.
For n > 4800, the probability is so close to 1 that we can just return 1.0.
*/

// https://leetcode.com/problems/soup-servings/

// Approach: Top-Down Dynamic Programming
// Time: O(192 * 192), 192 = 4800 / 25
// Space: O(192 * 192), 192 = 4800 / 25
func soupServings(n int) float64 {
	if n > 4800 {
		return 1
	}

	if n%25 != 0 {
		n = n/25 + 1
	} else {
		n /= 25
	}

	memo := make([][]*float64, n+1)
	for idx := range n + 1 {
		memo[idx] = make([]*float64, n+1)
	}

	var dp func(a int, b int) float64

	dp = func(a int, b int) float64 {
		if a <= 0 {
			if b <= 0 {
				return 0.5
			}

			return 1
		}

		if b <= 0 {
			return 0
		}

		if memo[a][b] != nil {
			return *memo[a][b]
		}

		memo[a][b] = new(float64)
		*memo[a][b] = 0.25*dp(a-4, b) +
			0.25*dp(a-3, b-1) +
			0.25*dp(a-2, b-2) +
			0.25*dp(a-1, b-3)

		return *memo[a][b]
	}

	return dp(n, n)
}
