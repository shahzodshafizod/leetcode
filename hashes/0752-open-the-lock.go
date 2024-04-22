package hashes

// https://leetcode.com/problems/open-the-lock/

func openLock(deadends []string, target string) int {
	var visited = make(map[string]bool)
	for _, deadend := range deadends {
		visited[deadend] = true
	}
	if visited["0000"] || visited[target] {
		return -1
	}
	if target == "0000" {
		return 0
	}
	var queue = []string{"0000"}
	visited["0000"] = true
	var level = 1
	for length := len(queue); length > 0; length = len(queue) {
		for idx := 0; idx < length; idx++ {
			for digit := 0; digit < 4; digit++ {
				for _, op := range []int{-1, 1} {
					next := []byte(queue[idx])
					next[digit] = '0' + uint8((int(next[digit]-'0')+op+10)%10)
					if !visited[string(next)] {
						if string(next) == target {
							return level
						}
						queue = append(queue, string(next))
						visited[string(next)] = true
					}
				}
			}
		}
		queue = queue[length:]
		level++
	}
	return -1
}
