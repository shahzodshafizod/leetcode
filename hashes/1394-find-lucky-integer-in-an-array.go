package hashes

// https://leetcode.com/problems/find-lucky-integer-in-an-array/

func findLucky(arr []int) int {
	counter := make(map[int]int, len(arr))
	for _, num := range arr {
		counter[num]++
	}
	lucky := -1
	for num, cnt := range counter {
		if num == cnt && num > lucky {
			lucky = num
		}
	}
	return lucky
}
