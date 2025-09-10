package greedy

// https://leetcode.com/problems/minimum-number-of-people-to-teach/

func minimumTeachings(n int, languages [][]int, friendships [][]int) int {
	unconnected := make(map[int]struct{})

	var (
		fr1, fr2 int
		canComm  bool
	)

	for idx := range friendships {
		fr1 = friendships[idx][0] - 1
		fr2 = friendships[idx][1] - 1

		slangs := make(map[int]bool)
		for _, lang := range languages[fr1] {
			slangs[lang] = true
		}

		canComm = false

		for _, lang := range languages[fr2] {
			if slangs[lang] {
				canComm = true

				break
			}
		}

		if !canComm {
			unconnected[fr1] = struct{}{}
			unconnected[fr2] = struct{}{}
		}
	}

	freq := make([]int, n+1)

	var maxCommon int

	for friend := range unconnected {
		for _, lang := range languages[friend] {
			freq[lang]++
			maxCommon = max(maxCommon, freq[lang])
		}
	}

	return len(unconnected) - maxCommon
}
