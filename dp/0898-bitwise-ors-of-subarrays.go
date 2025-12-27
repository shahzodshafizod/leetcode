package dp

// https://leetcode.com/problems/bitwise-ors-of-subarrays/

func subarrayBitwiseORs(arr []int) int {
	combined, prev := make(map[int]struct{}), make(map[int]struct{})

	for _, num := range arr {
		curr := map[int]struct{}{num: {}}

		combined[num] = struct{}{}
		for val := range prev {
			curr[val|num] = struct{}{}
			combined[val|num] = struct{}{}
		}

		prev = curr
	}

	return len(combined)
}
