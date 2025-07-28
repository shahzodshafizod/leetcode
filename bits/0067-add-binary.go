package bits

import "slices"

// https://leetcode.com/problems/add-binary/

func addBinary(a string, b string) string {
	aid := len(a) - 1
	bid := len(b) - 1
	carry := 0
	bisum := make([]byte, 0)

	for aid >= 0 || bid >= 0 {
		if aid >= 0 {
			carry += int(a[aid] - '0')
			aid--
		}

		if bid >= 0 {
			carry += int(b[bid] - '0')
			bid--
		}

		bisum = append(bisum, byte('0'+carry&1))
		carry >>= 1
	}

	if carry != 0 {
		bisum = append(bisum, byte('0'+carry))
	}

	slices.Reverse(bisum)

	return string(bisum)
}
