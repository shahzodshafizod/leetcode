package design

// https://leetcode.com/problems/data-stream-as-disjoint-intervals/

/*
[1,2,3,4,6,8,9] -> [[1,4], [6,6], [8,9]]
[1,2,3,4,6,7,8,9] -> [[1,4],[6,9]]
[1,2,3,4,5,6,7,8,9] -> [[1,9]]
*/

// Approach #2: Binary Search
// Time: O(NLogN)
// Space: O(N)
type SummaryRanges struct {
	intervals [][]int
}

func NewSummaryRanges() SummaryRanges {
	return SummaryRanges{intervals: make([][]int, 0)}
}

func (s *SummaryRanges) AddNum(value int) {
	n := len(s.intervals)
	left, right := 0, n-1

	var mid, start, end int

	for left <= right {
		mid = left + (right-left)>>1
		start, end = s.intervals[mid][0], s.intervals[mid][1]
		// 1. check if in a valid range
		if start <= value && value <= end {
			return
		}

		if start > value {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	prev, next := right, right+1
	// 2. check if can merge
	merged := false

	if prev >= 0 && s.intervals[prev][1]+1 == value {
		s.intervals[prev][1]++
		merged = true
	}

	if next < n && s.intervals[next][0]-1 == value {
		s.intervals[next][0]--
		merged = true
	}

	if merged {
		if prev >= 0 && next < n && s.intervals[prev][1] == s.intervals[next][0] {
			s.intervals[prev][1] = s.intervals[next][1]
			copy(s.intervals[next:], s.intervals[next+1:])
			s.intervals = s.intervals[:n-1]
		}

		return
	}
	// 3. insert a new range
	s.intervals = append(s.intervals, nil)
	copy(s.intervals[next+1:], s.intervals[next:])
	s.intervals[next] = []int{value, value}
}

func (s *SummaryRanges) GetIntervals() [][]int {
	return s.intervals
}

// // Approach #1: Union-Find
// // Time: O(N^2)
// // Space: O(N)
// type SummaryRanges struct {
// 	MAX    int
// 	parent []*int
// }

// func NewSummaryRanges() SummaryRanges {
// 	const max int = 1e4
// 	return SummaryRanges{
// 		MAX:    max,
// 		parent: make([]*int, max+2),
// 	}
// }

// func (s *SummaryRanges) AddNum(value int) {
// 	if value == 0 || s.parent[value-1] == nil {
// 		s.parent[value] = new(int)
// 		*s.parent[value] = value
// 	} else {
// 		s.parent[value] = s.parent[value-1]
// 	}
// 	if s.parent[value+1] != nil {
// 		s.parent[value+1] = s.parent[value]
// 	}
// }

// func (s *SummaryRanges) find(x int) int {
// 	if *s.parent[x] != x {
// 		*s.parent[x] = s.find(*s.parent[x])
// 	}
// 	return *s.parent[x]
// }

// func (s *SummaryRanges) GetIntervals() [][]int {
// 	var intervals = make([][]int, 0)
// 	var end int
// 	for idx := s.MAX; idx >= 0; idx-- {
// 		if s.parent[idx] == nil {
// 			continue
// 		}
// 		end = idx
// 		idx = s.find(idx)
// 		intervals = append([][]int{{idx, end}}, intervals...)
// 	}
// 	return intervals
// }

/**
 * Your SummaryRanges object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(value);
 * param_2 := obj.GetIntervals();
 */
