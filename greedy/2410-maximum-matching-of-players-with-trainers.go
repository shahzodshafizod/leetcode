package greedy

import "sort"

// https://leetcode.com/problems/maximum-matching-of-players-with-trainers/

// Approach: Two-Pointers
// Time: O(p log p + t log t), p=len(players), t=len(trainers)
// Space: O(p + t)
func matchPlayersAndTrainers(players []int, trainers []int) int {
	sort.Ints(players)
	sort.Ints(trainers)

	count := 0

	pi, pn := 0, len(players)
	for ti, tn := 0, len(trainers); ti < tn && pi < pn; ti++ {
		if players[pi] <= trainers[ti] {
			count++
			pi++
		}
	}

	return count
}
