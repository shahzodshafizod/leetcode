package maths

// https://leetcode.com/problems/maximum-manhattan-distance-after-k-changes/

// Approach #3
// Time: O(n)
// Space: O(1)
func maxDistance(s string, k int) int {
	var distance = 0
	var x, y int
	var lat, lon = 0, 0
	for idx, c := range s {
		switch c {
		case 'E':
			lon++
		case 'W':
			lon--
		case 'N':
			lat++
		case 'S':
			lat--
		}
		x, y = lon, lat
		if x < 0 {
			x = -x
		}
		if y < 0 {
			y = -y
		}
		distance = max(distance, min(x+y+k+k, idx+1))
	}
	return distance
}

// // Approach #2
// // Time: O(n)
// // Space: O(1)
// func maxDistance(s string, k int) int {
// 	var cnt = make(map[rune]int)
// 	var distance = 0
// 	var vmax, vmin, hmax, hmin, dist int
// 	for _, c := range s {
// 		cnt[c]++
// 		vmax, vmin = max(cnt['N'], cnt['S']), min(cnt['N'], cnt['S'])
// 		hmax, hmin = max(cnt['E'], cnt['W']), min(cnt['E'], cnt['W'])
// 		dist = vmax + 2*min(k, vmin) - vmin
// 		dist += hmax + 2*min(k-min(k, vmin), hmin) - hmin
// 		distance = max(distance, dist)
// 	}
// 	return distance
// }

// // Approach #1
// // Time: O(n)
// // Space: O(1)
// func maxDistance(s string, k int) int {
// 	var abs = func(x int) int {
// 		if x < 0 {
// 			return -x
// 		}
// 		return x
// 	}
// 	var change = map[rune]rune{'N': 'S', 'S': 'N', 'E': 'W', 'W': 'E'}
// 	var distance = 0
// 	for _, hor := range []rune{'E', 'W'} {
// 		for _, ver := range []rune{'N', 'S'} {
// 			var quota, x, y = k, 0, 0
// 			for _, c := range s {
// 				if quota > 0 && (c == hor || c == ver) {
// 					c = change[c]
// 					quota--
// 				}
// 				switch c {
// 				case 'N':
// 					y++
// 				case 'S':
// 					y--
// 				case 'E':
// 					x++
// 				case 'W':
// 					x--
// 				}
// 				distance = max(distance, abs(x)+abs(y))
// 			}
// 		}
// 	}
// 	return distance
// }
