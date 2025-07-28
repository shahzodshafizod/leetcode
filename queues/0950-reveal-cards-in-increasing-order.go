package queues

import (
	"sort"

	"github.com/shahzodshafizod/leetcode/pkg"
)

// https://leetcode.com/problems/reveal-cards-in-increasing-order/

// time: O(n log n)
// space: O(n)
func deckRevealedIncreasing(deck []int) []int {
	sort.Ints(deck) // O(n log n)
	revealed := make([]int, len(deck))

	queue := pkg.NewQueue[int]()
	for idx := range deck { // O(n)
		queue.Enqueue(idx)
	}

	for idx := range deck { // O(n)
		revealed[queue.Dequeue()] = deck[idx]

		queue.Enqueue(queue.Dequeue())
	}

	return revealed
}
