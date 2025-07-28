package slidingwindows

// https://leetcode.com/problems/count-the-hidden-sequences/

func numberOfArrays(differences []int, lower int, upper int) int {
	peak, peakVal := upper, upper
	deep, deepVal := lower, lower

	for idx, n := 0, len(differences); idx < n && peak >= deep; idx++ {
		peakVal += differences[idx]
		peak = min(peak, upper-(peakVal-upper))
		deepVal += differences[idx]
		deep = max(deep, lower+(lower-deepVal))
	}

	return max(0, peak-deep+1)
}
