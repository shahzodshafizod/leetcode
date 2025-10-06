package graphs

// https://leetcode.com/problems/bus-routes/

// Approach: Breadth-First Search
// Time: O(n*m), n=len(routes), m=len(routes[i])
// Space: O(n*m)
func numBusesToDestination(routes [][]int, source int, target int) int {
	graph := make(map[int][]int) // stop -> [bus]

	for bus := range routes {
		for _, stop := range routes[bus] {
			graph[stop] = append(graph[stop], bus)
		}
	}

	seenStops := make(map[int]bool)
	seenBuses := make([]bool, len(routes))
	queue := []int{source}

	var currStop, bus, nextStop int

	cnt := 0

	for n := len(queue); n > 0; n = len(queue) {
		for qid := range n {
			currStop = queue[qid]

			if currStop == target {
				return cnt
			}

			for _, bus = range graph[currStop] {
				if seenBuses[bus] {
					continue
				}

				for _, nextStop = range routes[bus] {
					if !seenStops[nextStop] {
						queue = append(queue, nextStop)
						seenStops[nextStop] = true
					}
				}

				seenBuses[bus] = true
			}
		}

		queue = queue[n:]

		cnt++
	}

	return -1
}
