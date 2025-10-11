package dp

import "math/bits"

// https://leetcode.com/problems/smallest-sufficient-team/

// Approach #2: Bottom-Up Dynamic Programming + Bit Manipulation
// Time: O(2^sn * pn), sn=len(reqSkills), pn=len(people)
// Space: O(2^sn)
func smallestSufficientTeam(reqSkills []string, people [][]string) []int {
	indices := make(map[string]int)
	for idx, skill := range reqSkills {
		indices[skill] = idx
	}

	n := len(people)
	skills := make([]int, n)

	for person := range people {
		for _, skill := range people[person] {
			skills[person] |= 1 << indices[skill]
		}
	}

	maximal := (1 << len(reqSkills)) - 1
	dp := make([]int, maximal+1)
	dp[maximal] = 0

	for smask := maximal - 1; smask >= 0; smask-- { // smask: skills mask
		dp[smask] = (1 << n) - 1

		var pmask int

		for person := range n {
			if smask|skills[person] != smask {
				pmask = dp[smask|skills[person]] | (1 << person)
				if bits.OnesCount(uint(pmask)) < bits.OnesCount(uint(dp[smask])) {
					dp[smask] = pmask
				}
			}
		}
	}

	mask := dp[0]

	var team []int

	for pid := range n {
		if mask&(1<<pid) != 0 {
			team = append(team, pid)
		}
	}

	return team
}

// // Approach #1: Top-Down Dynamic Programming + Bit Manipulation
// // Time: O(2^sn * pn), sn=len(reqSkills), pn=len(people)
// // Space: O(2^sn)
// func smallestSufficientTeam(reqSkills []string, people [][]string) []int {
//     indices := make(map[string]int)
//     for idx, skill := range reqSkills {
//         indices[skill] = idx
//     }
//     n := len(people)
//     skills := make([]int, n)
//     for person := range people {
//         for _, skill := range people[person] {
//             skills[person] |= 1 << indices[skill]
//         }
//     }

//     maximal := (1 << len(reqSkills)) - 1
//     memo := make([]*int, maximal+1)
//     memo[maximal] = new(int)

//     var dp func(smask int) int // smask: skills mask
//     dp = func(smask int) int {
//         if memo[smask] != nil {
//             return *memo[smask]
//         }
//         memo[smask] = new(int)
//         *memo[smask] = (1 << n) - 1
//         var pmask int
//         for person := range n {
//             if smask | skills[person] != smask {
//                 pmask = dp(smask | skills[person]) | (1 << person)
//                 if bits.OnesCount(uint(pmask)) < bits.OnesCount(uint(*memo[smask])) {
//                     *memo[smask] = pmask
//                 }
//             }
//         }
//         return *memo[smask]
//     }

//     mask := dp(0)
//     var team []int
//     for pid := range n {
//         if mask & (1 << pid) != 0 {
//             team = append(team, pid)
//         }
//     }
//     return team
// }
