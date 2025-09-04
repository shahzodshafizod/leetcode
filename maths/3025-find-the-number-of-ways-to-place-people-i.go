package maths

// https://leetcode.com/problems/find-the-number-of-ways-to-place-people-i/

func numberOfPairs(points [][]int) int {
	count, n := 0, len(points)

	var xi, yi, xj, yj, xk, yk int
	for i := range n {
		xi, yi = points[i][0], points[i][1]
		for j := range n {
			xj, yj = points[j][0], points[j][1]
			// we ensure, points[i] is on the upper left side of points[j]
			if i == j || xi > xj || yi < yj {
				continue
			}

			count++

			for k := range n {
				if k == i || k == j {
					continue
				}

				xk, yk = points[k][0], points[k][1]
				if xi <= xk && xk <= xj && yi >= yk && yk >= yj {
					count--

					break
				}
			}
		}
	}

	return count
}
