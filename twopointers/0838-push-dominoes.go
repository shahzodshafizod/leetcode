package twopointers

import "strings"

/*
'R' pushes right, 'L' pushes left.
If both hit the same dot, it stays upright.

'R......R' => 'RRRRRRR'
'R......L' => 'RRRRLLLL' or 'RRR.LLL'
'L......R' => 'L.....R'
'L......L' => 'LLLLLLLL'
*/

// https://leetcode.com/problems/push-dominoes/

// Approach: Calculate Force
// Time: O(n)
// Space: O(n)
func pushDominoes(dominoes string) string {
	n := len(dominoes)
	force := make([]int, n)
	f := 0

	for idx := n - 1; idx >= 0; idx-- {
		switch dominoes[idx] {
		case 'L':
			f = n
		case 'R':
			f = 0
		default:
			f = max(f-1, 0)
		}

		force[idx] -= f
	}

	var (
		sb  strings.Builder
		err error
	)

	f = 0

	for idx := range n {
		switch dominoes[idx] {
		case 'R':
			f = n
		case 'L':
			f = 0
		default:
			f = max(f-1, 0)
		}

		force[idx] += f
		if force[idx] > 0 {
			err = sb.WriteByte('R')
		} else if force[idx] < 0 {
			err = sb.WriteByte('L')
		} else {
			err = sb.WriteByte('.')
		}

		_ = err
	}

	return sb.String()
}
