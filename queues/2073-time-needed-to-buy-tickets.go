package queues

// https://leetcode.com/problems/time-needed-to-buy-tickets/

func timeRequiredToBuy(tickets []int, k int) int {
	seconds := 0

	for idx := range tickets {
		if idx <= k {
			seconds += min(tickets[idx], tickets[k])
		} else {
			seconds += min(tickets[idx], tickets[k]-1)
		}
	}

	return seconds
}
