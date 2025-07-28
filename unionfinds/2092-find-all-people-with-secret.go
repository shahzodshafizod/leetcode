package unionfinds

import (
	"sort"

	"github.com/shahzodshafizod/leetcode/pkg"
)

// https://leetcode.com/problems/find-all-people-with-secret/

func findAllPeople(n int, meetings [][]int, firstPerson int) []int {
	sort.Slice(meetings, func(i, j int) bool { return meetings[i][2] < meetings[j][2] })

	uf := pkg.NewUnionFind(n)
	uf.Union(0, firstPerson)

	told := make([]bool, n)
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
			if uf.Find(x) == uf.Find(0) {
				told[x] = true
			} else {
				uf.Reset(x)
			}

			y = meetings[i][1]
			if uf.Find(y) == uf.Find(0) {
				told[y] = true
			} else {
				uf.Reset(y)
			}
		}
	}

	informed := make([]int, 0)

	for person, aware := range told {
		if aware {
			informed = append(informed, person)
		}
	}

	return informed
}
