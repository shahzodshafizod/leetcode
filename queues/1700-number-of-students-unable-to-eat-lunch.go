package queues

// https://leetcode.com/problems/number-of-students-unable-to-eat-lunch/

// Counting Approach
func countStudents(students []int, sandwiches []int) int {
	var requirement [2]int
	for _, student := range students {
		requirement[student]++
	}
	for _, sandwich := range sandwiches {
		if requirement[sandwich] == 0 {
			return requirement[0] + requirement[1]
		}
		requirement[sandwich]--
	}
	return 0
}

// func countStudents(students []int, sandwiches []int) int {
// 	studentsCount, sandwichesCount := len(students), len(sandwiches)
// 	var eatenCount int
// 	for {
// 		eatenCount = 0
// 		for i := 0; i < studentsCount && sandwichesCount > 0; i++ {
// 			if students[i] == sandwiches[0] {
// 				eatenCount++
// 				students = append(students[:i], students[i+1:]...)
// 				studentsCount--
// 				sandwiches = sandwiches[1:]
// 				sandwichesCount--
// 				i--
// 			}
// 		}
// 		if eatenCount == 0 {
// 			break
// 		}
// 	}
// 	return studentsCount
// }
