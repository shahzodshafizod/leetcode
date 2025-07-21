package greedy

import "slices"

// https://leetcode.com/problems/minimum-number-of-moves-to-seat-everyone/

// Approach #2: Bucket (Count) Sorting
// n = len(seats) = len(students)
// m = max(maxseat, maxstudent)
// time: O(4n + m) = O(n + m)
// space: O(m)
func minMovesToSeat(seats []int, students []int) int {
	maxseat := slices.Max(seats) // O(n)
	countseat := make([]int, maxseat+1)
	for _, seat := range seats { // O(n)
		countseat[seat]++
	}
	maxstudent := slices.Max(students) // O(n)
	countstudent := make([]int, maxstudent+1)
	for _, student := range students { // O(n)
		countstudent[student]++
	}
	moves := 0
	seat, student := 0, 0
	for { // O(m)
		for seat <= maxseat && countseat[seat] == 0 {
			seat++
		}
		for student <= maxstudent && countstudent[student] == 0 {
			student++
		}
		if seat > maxseat || student > maxstudent {
			break
		}
		if student > seat {
			moves += student - seat
		} else {
			moves += seat - student
		}
		countseat[seat]--
		countstudent[student]--
	}
	return moves
}

// // Approach #1: Sorting (Greedy)
// // time: O(2n log n) = O(n log n)
// // space: O(1)
// func minMovesToSeat(seats []int, students []int) int {
// 	sort.Ints(seats)
// 	sort.Ints(students)
// 	var moves = 0
// 	for idx := range seats {
// 		if seats[idx] > students[idx] {
// 			moves += seats[idx] - students[idx]
// 		} else {
// 			moves += students[idx] - seats[idx]
// 		}
// 	}
// 	return moves
// }
