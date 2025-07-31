package stacks

import (
	"sort"
)

// https://leetcode.com/problems/robot-collisions/

// time: O(N Log N)
// space: O(N)
func survivedRobotsHealths(positions []int, healths []int, directions string) []int {
	posid := make([][2]int, len(positions))
	for idx, position := range positions {
		posid[idx][0] = position
		posid[idx][1] = idx
	}

	sort.Slice(posid, // O(N Log N)
		func(i, j int) bool { return posid[i][0] < posid[j][0] },
	)

	type node struct {
		index int
		next  *node
	}

	var stack *node

	var newid, topid int
	for _, posid := range posid {
		newid = posid[1]
		if directions[newid] == 'R' {
			stack = &node{index: newid, next: stack}
		} else {
			for stack != nil && healths[newid] > 0 {
				topid = stack.index
				stack = stack.next

				if healths[newid] > healths[topid] {
					healths[topid] = 0
					healths[newid]--
				} else if healths[newid] < healths[topid] {
					healths[newid] = 0
					healths[topid]--
					stack = &node{index: topid, next: stack}
				} else {
					healths[newid] = 0
					healths[topid] = 0
				}
			}
		}
	}

	survivors := make([]int, 0)

	for _, health := range healths {
		if health > 0 {
			survivors = append(survivors, health)
		}
	}

	return survivors
}
