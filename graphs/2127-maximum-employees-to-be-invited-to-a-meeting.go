package graphs

import "container/list"

// https://leetcode.com/problems/maximum-employees-to-be-invited-to-a-meeting/

/*
Overview:
1. Every connected component has exactly one cycle.
2. Components with a cycle of length = 2, form non-closed circle.
	They can be connected to other non closed circles.
3. Components with a cycle of length > 2, form closed circles.
	They can't be connected

Solution:
1. Find the longest cycle, aka the longest closed circle.
2. Find all components with length = 2 cycles.
	For each, find the longest path component.
	Sum all of the paths.
3. Return max of step 1 and 2.
*/

// Approach: Topological Sort to Reduce Non-Cyclic Nodes
// Time: O(n)
// Space: O(n)
func maximumInvitations(favorite []int) int {
	var n = len(favorite)
	var indegree = make([]int, n)
	for empl := 0; empl < n; empl++ {
		indegree[favorite[empl]]++
	}
	var queue = list.New()
	for empl := 0; empl < n; empl++ {
		if indegree[empl] == 0 {
			queue.PushBack(empl)
		}
	}
	var depth = make([]int, n)
	var currEmpl, nextEmpl int
	for queue.Len() > 0 {
		currEmpl = queue.Remove(queue.Front()).(int)
		nextEmpl = favorite[currEmpl]
		depth[nextEmpl] = max(depth[nextEmpl], depth[currEmpl]+1)
		indegree[nextEmpl]--
		if indegree[nextEmpl] == 0 {
			queue.PushBack(nextEmpl)
		}
	}
	var longestCycle = 0
	var twoLenCycle = 0
	var cycleLength, tmpEmpl int
	for empl := 0; empl < n; empl++ {
		if indegree[empl] == 0 {
			continue
		}
		cycleLength = 0
		tmpEmpl = empl
		for indegree[tmpEmpl] != 0 {
			indegree[tmpEmpl]--
			cycleLength++
			tmpEmpl = favorite[tmpEmpl]
		}
		if cycleLength == 2 {
			twoLenCycle += depth[empl] + depth[favorite[empl]] + 2
		} else {
			longestCycle = max(longestCycle, cycleLength)
		}
	}
	return max(longestCycle, twoLenCycle)
}
