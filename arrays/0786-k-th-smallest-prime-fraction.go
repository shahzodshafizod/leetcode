package arrays

// https://leetcode.com/problems/k-th-smallest-prime-fraction/

// Approach: Binary Search
// time: O(n log m^2); m - max array value
// space: O(1)
func kthSmallestPrimeFraction(arr []int, k int) []int {
	n := len(arr)
	var left, right float64 = 0, 1
	var mid, fraction, maxfraction float64
	var smallfractions, numerator, denominator, j int
	for left < right {
		mid = (left + right) / 2
		maxfraction = 0.0
		smallfractions = 0
		numerator, denominator = 0, 0
		j = 1
		for i := 0; i < n; i++ {
			for j < n && float64(arr[i]) >= mid*float64(arr[j]) {
				j++
			}
			if j == n {
				break
			}
			smallfractions += n - j
			fraction = float64(arr[i]) / float64(arr[j])
			if fraction > maxfraction {
				maxfraction = fraction
				numerator, denominator = i, j
			}
		}
		if smallfractions == k {
			break
		} else if smallfractions > k {
			right = mid
		} else {
			left = mid
		}
	}
	return []int{arr[numerator], arr[denominator]}
}

// // Approach: Priority Queue
// // time: O((n+k) log n)
// // space: O(n)
// func kthSmallestPrimeFraction(arr []int, k int) []int {
// 	type fraction struct {
// 		fraction       float64
// 		numeratorIdx   int
// 		denominatorIdx int
// 	}
// 	var pq = design.NewPQ(
// 		make([]*fraction, 0),
// 		func(x, y *fraction) bool { return x.fraction > y.fraction },
// 	)
// 	var n = len(arr)
// 	for i := 0; i < n-1; i++ {
// 		pq.Push(&fraction{
// 			fraction:       float64(arr[i]) / float64(arr[n-1]),
// 			numeratorIdx:   i,
// 			denominatorIdx: n - 1,
// 		})
// 	}
// 	// the top element of the queue is the smallest
// 	// fraction among all fractions formed
// 	// by dividing each element by the last element.

// 	// Essentially, we replace the smallest fraction
// 	// with the next smallest fraction having
// 	// the same numerator. Repeating this k - 1 times
// 	// leaves the kth smallest fraction
// 	// at the top of the priority queue.
// 	var top *fraction
// 	for k > 1 {
// 		k--
// 		top = pq.Pop()
// 		top.denominatorIdx--
// 		if top.numeratorIdx < top.denominatorIdx {
// 			pq.Push(&fraction{
// 				fraction:       float64(arr[top.numeratorIdx]) / float64(arr[top.denominatorIdx]),
// 				numeratorIdx:   top.numeratorIdx,
// 				denominatorIdx: top.denominatorIdx,
// 			})
// 		}
// 	}
// 	return []int{arr[pq.Peek().numeratorIdx], arr[pq.Peek().denominatorIdx]}
// }

// // Approach: brute Force
// // time: O(n x n)
// // space: O(k)
// func kthSmallestPrimeFraction(arr []int, k int) []int {
// 	type fraction struct {
// 		fraction    float32
// 		numerator   int
// 		denominator int
// 	}
// 	var pq = design.NewPQ(
// 		make([]fraction, 0),
// 		func(x, y fraction) bool { return x.fraction < y.fraction },
// 	)
// 	var n = len(arr)
// 	var currfrac float32
// 	for i := 0; i < n; i++ {
// 		for j := i + 1; j < n; j++ {
// 			currfrac = float32(arr[i]) / float32(arr[j])
// 			if pq.Len() < k {
// 				pq.Push(fraction{
// 					currfrac,
// 					arr[i],
// 					arr[j],
// 				})
// 			} else if currfrac < pq.Peek().fraction {
// 				pq.Pop()
// 				pq.Push(fraction{currfrac, arr[i], arr[j]})
// 			}
// 		}
// 	}
// 	return []int{pq.Peek().numerator, pq.Peek().denominator}
// }
