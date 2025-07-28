package design

// https://leetcode.com/problems/my-calendar-iii/

// Approach: Line Sweep
type MyCalendarThree struct {
	count  map[int]int
	points []int
}

func NewMyCalendarThree() MyCalendarThree {
	return MyCalendarThree{count: make(map[int]int)}
}

// Time: O(3N*N) = O(N^2)
// Space: O(2N) = O(N)
func (m *MyCalendarThree) Book(startTime int, endTime int) int {
	m.addKey(startTime) // O(N)
	m.addKey(endTime)   // O(N)

	m.count[startTime]++
	m.count[endTime]--

	overlap, maxOverlap := 0, 0
	for _, point := range m.points { // O(N)
		overlap += m.count[point]
		maxOverlap = max(maxOverlap, overlap)
	}

	return maxOverlap
}

func (m *MyCalendarThree) addKey(key int) {
	if _, ok := m.count[key]; !ok {
		m.points = append(m.points, key)
		for idx := len(m.points) - 1; idx > 0; idx-- {
			if m.points[idx-1] <= m.points[idx] {
				break
			}

			m.points[idx-1], m.points[idx] = m.points[idx], m.points[idx-1]
		}
	}
}

/**
 * Your MyCalendarThree object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(startTime,endTime);
 */
