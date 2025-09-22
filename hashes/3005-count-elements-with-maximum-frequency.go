package hashes

// https://leetcode.com/problems/count-elements-with-maximum-frequency/

func maxFrequencyElements(nums []int) int {
	freq := make(map[int]int)

	var total, maxFreq int

	for _, num := range nums {
		freq[num]++
		if freq[num] > maxFreq {
			total = freq[num]
			maxFreq = freq[num]
		} else if freq[num] == maxFreq {
			total += freq[num]
		}
	}

	return total
}
