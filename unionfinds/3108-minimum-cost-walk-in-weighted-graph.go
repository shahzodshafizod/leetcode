package unionfinds

import (
	"github.com/shahzodshafizod/leetcode/pkg"
)

// https://leetcode.com/problems/minimum-cost-walk-in-weighted-graph/

// Time: O(v+e+q), v=n, e=len(edges), q=len(query)
// Space: O(v)
func minimumCost(n int, edges [][]int, query [][]int) []int {
	uf := pkg.NewUnionFindRanked(n)
	for _, edge := range edges {
		uf.Union(edge[0], edge[1])
	}
	walks := make([]*int, n)
	var root int
	for _, edge := range edges {
		root = uf.Find(edge[0])
		if walks[root] == nil {
			walks[root] = new(int)
			*walks[root] = -1
		}
		*walks[root] &= edge[2]
	}
	answer := make([]int, len(query))
	var rs, rt int
	for idx := range query {
		rs = uf.Find(query[idx][0])
		rt = uf.Find(query[idx][1])
		if rs == rt {
			answer[idx] = *walks[rs]
		} else {
			answer[idx] = -1
		}
	}
	return answer
}
