package strings

// https://leetcode.com/problems/largest-3-same-digit-number-in-string/

func largestGoodInteger(num string) string {
	good := ""

	for i, n := 2, len(num); i < n; i++ {
		if num[i-2] == num[i-1] && num[i-1] == num[i] {
			good = max(good, num[i-2:i+1])
		}
	}

	return good
}
