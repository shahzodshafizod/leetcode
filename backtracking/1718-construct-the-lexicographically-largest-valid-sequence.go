package backtracking

// https://leetcode.com/problems/construct-the-lexicographically-largest-valid-sequence/

// Approach: Backtracking
// Time: O(n!)
// Space: O(n)
func constructDistancedSequence(n int) []int {
	size := 2*n - 1
	seq := make([]int, size)
	used := make([]bool, n+1)

	var backtrack func(int) bool

	backtrack = func(idx int) bool {
		if idx == size {
			return true
		}

		if seq[idx] != 0 {
			return backtrack(idx + 1)
		}

		for num := n; num > 0; num-- {
			if used[num] {
				continue
			}

			seq[idx] = num
			used[num] = true

			if num == 1 {
				if backtrack(idx + 1) {
					return true
				}
			} else if idx+num < size && seq[idx+num] == 0 {
				seq[idx+num] = num

				if backtrack(idx + 1) {
					return true
				}

				seq[idx+num] = 0
			}

			seq[idx] = 0
			used[num] = false
		}

		return false
	}
	backtrack(0)

	return seq
}
