package queues

// https://leetcode.com/problems/number-of-students-unable-to-eat-lunch/

func countStudents(students []int, sandwiches []int) int {
	studentsCount, sandwichesCount := len(students), len(sandwiches)
	var eatenCount int
	for {
		eatenCount = 0
		for i := 0; i < studentsCount && sandwichesCount > 0; i++ {
			if students[i] == sandwiches[0] {
				eatenCount++
				students = append(students[:i], students[i+1:]...)
				studentsCount--
				sandwiches = sandwiches[1:]
				sandwichesCount--
				i--
			}
		}
		if eatenCount == 0 {
			break
		}
	}
	return studentsCount
}
