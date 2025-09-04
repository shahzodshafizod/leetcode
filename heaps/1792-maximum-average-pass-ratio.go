package heaps

import (
	"container/heap"

	"github.com/shahzodshafizod/leetcode/pkg"
)

// https://leetcode.com/problems/maximum-average-pass-ratio/

func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	var (
		pq = pkg.NewHeap([][3]float64{}, func(x, y [3]float64) bool {
			return x[0] > y[0]
		})
		p, t float64
	)

	var ratio float64
	for idx := range classes {
		p, t = float64(classes[idx][0]), float64(classes[idx][1])
		heap.Push(pq, [3]float64{(p+1)/(t+1) - p/t, p, t})
		ratio += p / t
	}

	for ; extraStudents > 0; extraStudents-- {
		item, ok := heap.Pop(pq).([3]float64)
		_ = ok
		p, t = item[1], item[2]
		ratio -= p / t
		p, t = p+1, t+1
		ratio += p / t
		heap.Push(pq, [3]float64{(p+1)/(t+1) - p/t, p, t})
	}

	return ratio / float64(len(classes))
}
