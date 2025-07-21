package unionfinds

// https://leetcode.com/problems/check-if-the-rectangle-corner-is-reachable/

// If the top/left border to the bottom/right border of the rectangle is
// blocked by a group of connected circles, then the answer will be false.
// AND the chain should go through the rectangle, not outside, ex. of edge case:
// beginning and end of the chain touch borders, but the lies goes outside

// TODO: Solve this problem by yourself

// Approach: Union-Find
// Time: O(N)
// Space: O(N)
func canReachCorner(xCorner int, yCorner int, circles [][]int) bool {
	n := len(circles)
	parent := make([]int, n+2)
	for idx := range parent {
		parent[idx] = idx
	}
	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	var x, y, r int
	for idx := range circles {
		x, y, r = circles[idx][0], circles[idx][1], circles[idx][2]
		// if the circle overlaps right-top
		if (x-xCorner)*(x-xCorner)+(y-yCorner)*(y-yCorner) <= r*r {
			return false
		}
		// // if the circle overlaps left-bottom
		// if x*x+y*y <= r*r {
		// 	return false
		// }
		if x <= r && y <= yCorner || y+r >= yCorner && x <= xCorner {
			parent[find(n)] = find(idx)
		}
		if y <= r && x <= xCorner || x+r >= xCorner && y <= yCorner {
			parent[find(n+1)] = find(idx)
		}
		// if the circle touches both borders
		if find(n) == find(n+1) {
			return false
		}
	}
	// if either border is untouched
	if find(n) == n || find(n+1) == n+1 {
		return true
	}
	var x2, y2, r2 int
	for i := 0; i < n; i++ {
		x, y, r = circles[i][0], circles[i][1], circles[i][2]
		// if circle 1 is out of usable range
		if x-r >= xCorner || y-r >= yCorner || x >= xCorner && y >= yCorner {
			continue
		}
		for j := 0; j < i; j++ {
			// if already unioned
			if find(i) == find(j) {
				continue
			}
			x2, y2, r2 = circles[j][0], circles[j][1], circles[j][2]
			// if pair is out of usable range
			if x+x2 >= 2*xCorner && y+y2 >= 2*yCorner {
				continue
			}
			// if circles intersect
			if (x-x2)*(x-x2)+(y-y2)*(y-y2) <= (r+r2)*(r+r2) {
				parent[find(i)] = find(j)
				// if both borders are unioned
				if find(n) == find(n+1) {
					return false
				}
			}
		}
	}
	return true
}

// // Approach #2: Graph
// // Time: O(N^2)
// // Space: O(N)
// func canReachCorner(xCorner int, yCorner int, circles [][]int) bool {
// 	var n = len(circles)
// 	var edges = make([]int, n)
// 	var inside = make([]bool, n)
// 	var x, y, r int
// 	for idx := range circles {
// 		x, y, r = circles[idx][0], circles[idx][1], circles[idx][2]
// 		inside[idx] = x <= xCorner && y <= yCorner
// 		// if the circle intersects with start/end point
// 		if x*x+y*y <= r*r || (x-xCorner)*(x-xCorner)+(y-yCorner)*(y-yCorner) <= r*r {
// 			return false
// 		}
// 		// check horizontal edges
// 		if y-r <= 0 && y+r >= 0 && x <= xCorner ||
// 			x+r >= xCorner && x-r <= xCorner && y <= yCorner {
// 			edges[idx] = 1
// 		}
// 		// check vertical edges
// 		if x-r <= 0 && x+r >= 0 && y <= yCorner ||
// 			y+r >= yCorner && y-r <= yCorner && x <= xCorner {
// 			edges[idx] += 2
// 		}
// 		if edges[idx] == 3 {
// 			return false
// 		}
// 	}
// 	var seen = make([]bool, n)
// 	var dfs func(idx int) int
// 	dfs = func(idx int) int {
// 		seen[idx] = true
// 		var mask = edges[idx]
// 		var x, y, r = circles[idx][0], circles[idx][1], circles[idx][2]
// 		var x2, y2, r2 int
// 		for j := 0; j < n; j++ {
// 			if seen[j] {
// 				continue
// 			}
// 			x2, y2, r2 = circles[j][0], circles[j][1], circles[j][2]
// 			var dx = x - x2
// 			var dy = y - y2
// 			var dr = r + r2
// 			if dx*dx+dy*dy <= dr*dr &&
// 				(inside[idx] || inside[j] || x+x2 <= 2*xCorner && y+y2 <= 2*yCorner) {
// 				mask |= dfs(j)
// 			}
// 		}
// 		return mask
// 	}
// 	for idx := 0; idx < n; idx++ {
// 		if !seen[idx] && (dfs(idx) == 3) {
// 			return false
// 		}
// 	}
// 	return true
// }

// // Approach #1: Graph
// // Time: O(N^2)
// // Space: O(N)
// func canReachCorner(xCorner int, yCorner int, circles [][]int) bool {
// 	// inCircle returns whether Point (px, py) is located
// 	// inside Circle (cx, cy, r) (including the border)
// 	var inCircle = func(circle []int, px, py int) bool {
// 		cx, cy, radius := circle[0], circle[1], circle[2]
// 		return (cx-px)*(cx-px)+(cy-py)*(cy-py) <= radius*radius
// 	}
// 	var abs = func(x int) int {
// 		if x < 0 {
// 			return -x
// 		}
// 		return x
// 	}
// 	visited := make([]bool, len(circles))
// 	var dfs func(i int) bool
// 	dfs = func(i int) bool {
// 		x1, y1, r1 := circles[i][0], circles[i][1], circles[i][2]
// 		// Circle i intersects with the bottom or right border of the rectangle
// 		if y1 <= yCorner && abs(x1-xCorner) <= r1 ||
// 			x1 <= xCorner && y1 <= r1 ||
// 			x1 > xCorner && inCircle(circles[i], xCorner, 0) {
// 			return true
// 		}
// 		visited[i] = true
// 		for j, circle := range circles {
// 			if visited[j] {
// 				continue
// 			}
// 			x2, y2, r2 := circle[0], circle[1], circle[2]
// 			// Let Point A which |O1A| / |O1O2| = r1 / (r1+r2).
// 			// If two circles are connected to each other then A must be inside the intersection
// 			// And its coordinate is: (x1·r2+x2·r1)/(r1+r2), (y1·r2+y2·r1)/(r1+r2)
// 			var distance = (x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)
// 			var rdistance = (r1 + r2) * (r1 + r2)
// 			if distance <= rdistance &&
// 				x1*r2+x2*r1 < (r1+r2)*xCorner &&
// 				y1*r2+y2*r1 < (r1+r2)*yCorner &&
// 				dfs(j) {
// 				return true
// 			}
// 		}
// 		return false
// 	}
// 	for i, circle := range circles {
// 		x, y, r := circle[0], circle[1], circle[2]
// 		// DFS starts from circles which intersects the left/top border of the rectangle
// 		if inCircle(circle, 0, 0) {
// 			return false
// 		}
// 		if inCircle(circle, xCorner, yCorner) {
// 			return false
// 		}
// 		if visited[i] {
// 			continue
// 		}
// 		if (x <= xCorner && abs(y-yCorner) <= r ||
// 			y <= yCorner && x <= r ||
// 			y > yCorner && inCircle(circle, 0, yCorner)) &&
// 			dfs(i) {
// 			return false
// 		}
// 	}
// 	return true
// }
