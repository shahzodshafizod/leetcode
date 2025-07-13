package greedy

// https://leetcode.com/problems/reschedule-meetings-for-maximum-free-time-ii/

// Approach: Greedy
// Time: O(n)
// Space: O(1)
func maxFreeTime(eventTime int, startTime []int, endTime []int) int {
    var free, n = 0, len(startTime)
    // left to right
    var slot, left = 0, 0
    var right, meeting int
    for i := 0; i < n; i++ {
        meeting = endTime[i] - startTime[i]
        if slot >= meeting {
            meeting = 0
        }
        right = eventTime
        if i+1 < n {
            right = startTime[i+1]
        }
        free = max(free, right - left - meeting)
        slot = max(slot, startTime[i] - left)
        left = endTime[i]
    }
    // right to left
    slot, right = 0, eventTime
    for i := n-1; i >= 0; i-- {
        meeting = endTime[i] - startTime[i]
        if slot >= meeting {
            meeting = 0
        }
        left = 0
        if i > 0 {
            left = endTime[i-1]
        }
        free = max(free, right - left - meeting)
        slot = max(slot, right - endTime[i])
        right = startTime[i]
    }
    return free
}