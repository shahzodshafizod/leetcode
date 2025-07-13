package slidingwindows

// https://leetcode.com/problems/reschedule-meetings-for-maximum-free-time-i/

func maxFreeTime(eventTime int, k int, startTime []int, endTime []int) int {
    var free, busy, n = 0, 0, len(startTime)
	var left, right int
	for i := 0; i < n; i++ {
		busy += endTime[i] - startTime[i]
		left = 0
		if i+1 > k {
			left = endTime[i-k]
		}
		right = eventTime
		if i+1 < n {
			right = startTime[i+1]
		}
		free = max(free, right - left - busy)
		if i+1 >= k {
			busy -= endTime[i+1-k] - startTime[i+1-k]
		}
	}
	return free
}