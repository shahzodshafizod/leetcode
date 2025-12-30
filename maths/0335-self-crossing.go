package maths

// https://leetcode.com/problems/self-crossing/

// Approach: Pattern Recognition and Geometry
// Key insight: Due to the counter-clockwise spiral pattern, the i-th line can only
// cross with lines that are 3, 4, or 5 steps before it.
//
// Three crossing patterns:
//
//  1. Fourth line crosses first line (i crosses i-3):
//     Happens when: distance[i] >= distance[i-2] and distance[i-1] <= distance[i-3]
//
//  2. Fifth line crosses second line (i crosses i-4):
//     Happens when: distance[i-1] == distance[i-3] and
//     distance[i] + distance[i-4] >= distance[i-2]
//
//  3. Sixth line crosses third line (i crosses i-5):
//     Happens when: distance[i-2] >= distance[i-4] and
//     distance[i-3] >= distance[i-1] and
//     distance[i-1] + distance[i-5] >= distance[i-3] and
//     distance[i] + distance[i-4] >= distance[i-2]
//
// Visual patterns:
// Pattern 1 (4th crosses 1st):    Pattern 2 (5th crosses 2nd):    Pattern 3 (6th crosses 3rd):
//
//	0                              0──1                             0──1
//	│                              │  │                             │  │
//
// 3───┼──1                       4───┼──┼──2                      5───┼──2
//
//	│  │                           │  │                         │   │  │
//	└──2                           └──3                         └───┼──3
//	                                                                │
//	                                                                4
//
// Time: O(n) where n is the length of distance array
// Space: O(1) - only using a few variables
func isSelfCrossing(distance []int) bool {
	n := len(distance)
	if n < 4 {
		return false
	}

	for i := 3; i < n; i++ {
		// Pattern 1: Fourth line crosses first line (i crosses i-3)
		if distance[i] >= distance[i-2] && distance[i-1] <= distance[i-3] {
			return true
		}

		// Pattern 2: Fifth line crosses second line (i crosses i-4)
		if i >= 4 {
			if distance[i-1] == distance[i-3] && distance[i]+distance[i-4] >= distance[i-2] {
				return true
			}
		}

		// Pattern 3: Sixth line crosses third line (i crosses i-5)
		if i >= 5 {
			if distance[i-2] >= distance[i-4] &&
				distance[i-3] >= distance[i-1] &&
				distance[i-1]+distance[i-5] >= distance[i-3] &&
				distance[i]+distance[i-4] >= distance[i-2] {
				return true
			}
		}
	}

	return false
}
