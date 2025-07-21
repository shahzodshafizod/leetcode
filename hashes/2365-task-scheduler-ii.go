package hashes

// https://leetcode.com/problems/task-scheduler-ii/

func taskSchedulerII(tasks []int, space int) int64 {
	var days int64 = 0
	nextday := make(map[int]int64)
	for _, task := range tasks {
		days = max(days, nextday[task])
		days++
		nextday[task] = days + int64(space)
	}
	return days
}
