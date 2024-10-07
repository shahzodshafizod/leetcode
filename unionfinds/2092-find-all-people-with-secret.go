package unionfinds

import (
	"sort"

	"github.com/shahzodshafizod/leetcode/design"
)

// https://leetcode.com/problems/find-all-people-with-secret/

func findAllPeople(n int, meetings [][]int, firstPerson int) []int {
	sort.Slice(meetings, func(i, j int) bool { return meetings[i][2] < meetings[j][2] })
	var uf = design.NewDisjointSet(n)
	uf.Union(0, firstPerson)
	var told = make([]bool, n)
	told[0] = true
	told[firstPerson] = true
	var x, y, time int
	for i, len := 0, len(meetings); i < len; {
		time = meetings[i][2]
		for j := i; j < len && meetings[j][2] == time; j++ {
			x = meetings[j][0]
			y = meetings[j][1]
			uf.Union(x, y)
		}
		for ; i < len && meetings[i][2] == time; i++ {
			x = meetings[i][0]
			if uf.Connected(x, 0) {
				told[x] = true
			} else {
				uf.Reset(x)
			}
			y = meetings[i][1]
			if uf.Connected(y, 0) {
				told[y] = true
			} else {
				uf.Reset(y)
			}
		}
	}
	var informed = make([]int, 0)
	for person, aware := range told {
		if aware {
			informed = append(informed, person)
		}
	}
	return informed
}
