package queues

// https://leetcode.com/problems/find-the-winner-of-the-circular-game/

// Approach #2: Math
func findTheWinner(n int, k int) int {
	winner := 0
	for p := 1; p <= n; p++ {
		winner = (winner + k) % p
	}
	return winner + 1
}

// // Approach #1: Queue
// func findTheWinner(n int, k int) int {
// 	type node struct {
// 		val  int
// 		next *node
// 	}
// 	var head = &node{}
// 	var tail = head
// 	for i := 1; i <= n; i++ {
// 		tail.next = &node{val: i}
// 		tail = tail.next
// 	}
// 	head = head.next
// 	for head != tail {
// 		for i := 1; i < k; i++ {
// 			tail.next = head
// 			tail = head
// 			head = head.next
// 		}
// 		head = head.next
// 	}
// 	return head.val
// }
