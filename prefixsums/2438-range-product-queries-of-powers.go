package prefixsums

// https://leetcode.com/problems/range-product-queries-of-powers/

func productQueries(n int, queries [][]int) []int {
	const mod = int(1e9) + 7

	var powers []int

	for i := range 32 {
		if (n>>i)&1 == 1 {
			powers = append(powers, 1<<i)
		}
	}

	m := len(powers)
	prepro := make([][]int, m)

	var pro int

	for i := range m {
		prepro[i] = make([]int, m)

		pro = 1
		for j := i; j < m; j++ {
			pro = pro * powers[j] % mod
			prepro[i][j] = pro
		}
	}

	m = len(queries)
	answers := make([]int, m)

	var left, right int
	for i := range m {
		left, right = queries[i][0], queries[i][1]
		answers[i] = prepro[left][right]
	}

	return answers
}
