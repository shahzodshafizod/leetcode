package slidingwindows

// https://leetcode.com/problems/grumpy-bookstore-owner/

// time: O(n)
// space: O(1)
func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	var start = 0
	var window, maxWindow = 0, 0
	var satisfied = 0
	for end := range grumpy {
		if grumpy[end] == 1 {
			window += customers[end]
		} else {
			satisfied += customers[end]
		}
		if end-start >= minutes {
			window -= grumpy[start] * customers[start]
			start++
		}
		maxWindow = max(maxWindow, window)
	}
	return satisfied + maxWindow
}
