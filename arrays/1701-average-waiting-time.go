package arrays

// https://leetcode.com/problems/average-waiting-time/

func averageWaitingTime(customers [][]int) float64 {
	var nextIdleTime = 0
	var netWaitingTime = 0
	for _, customer := range customers {
		nextIdleTime = max(nextIdleTime, customer[0]) + customer[1]
		netWaitingTime += nextIdleTime - customer[0]
	}
	return float64(netWaitingTime) / float64(len(customers))
}
