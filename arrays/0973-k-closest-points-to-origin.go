package arrays

import (
	"math"

	"github.com/shahzodshafizod/alkhwarizmi/design"
)

// https://leetcode.com/problems/k-closest-points-to-origin/

// Priority Queue
func kClosest(points [][]int, k int) [][]int {
	var distance = func(x int, y int) float64 { return math.Sqrt(float64(x*x + y*y)) }
	var minHeap = design.NewPQ[[]int](
		make([][]int, 0),
		func(x, y []int) bool { return distance(x[0], x[1]) > distance(y[0], y[1]) },
	)
	for _, point := range points {
		minHeap.Push(point)
	}
	var closestPoints = make([][]int, 0, k)
	for k > 0 {
		k--
		point := minHeap.Pop()
		closestPoints = append(closestPoints, []int{point[0], point[1]})
	}
	return closestPoints
}

// // Quick Select (TLE)
// func kClosest(points [][]int, k int) [][]int {
// 	var len = len(points)
// 	if len == k {
// 		return points
// 	}
// 	var distance = func(p []int) float64 { return math.Sqrt(float64(p[0]*p[0] + p[1]*p[1])) }
// 	var swap = func(i int, j int) { points[i], points[j] = points[j], points[i] }
// 	var left, right = 0, len - 1
// 	for left < right {
// 		var pivotDistance = distance(points[left])
// 		var mid = left + 1
// 		for i := mid + 1; i < len; i++ {
// 			if distance(points[i]) < pivotDistance {
// 				swap(mid, i)
// 				mid++
// 			}
// 		}
// 		swap(mid, left)
// 		if mid+1 < k {
// 			left = mid + 1
// 		} else if mid+1 > k {
// 			right = mid - 1
// 		}
// 	}
// 	return points[:k]
// }