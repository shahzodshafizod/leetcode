package strings

// https://leetcode.com/problems/find-beautiful-indices-in-the-given-array-ii/

// Approach #2: KMP Pattern Matching + Binary Search
// Time: O(N+M1+M2)
// Space: O(M1+M2)
func beautifulIndices(s string, a string, b string, k int) []int {
	kmp := func(substr string) []int {
		n := len(substr)
		lps := make([]int, n)
		preflen, idx := 0, 1

		for idx < n {
			if substr[idx] == substr[preflen] {
				preflen++
				lps[idx] = preflen
				idx++
			} else if preflen > 0 {
				preflen = lps[preflen-1]
			} else {
				lps[idx] = 0
				idx++
			}
		}

		return lps
	}

	na, nb := len(a), len(b)
	v1 := kmp(a + "#" + s)
	v2 := kmp(b + "#" + s)

	var aindices []int

	for i, v := range v1 {
		if v >= na {
			aindices = append(aindices, i-na-na)
		}
	}

	var bindices []int

	for j, v := range v2 {
		if v >= nb {
			bindices = append(bindices, j-nb-nb)
		}
	}

	var res []int

	j, nj := 0, len(bindices)
	for _, i := range aindices {
		for j < nj && bindices[j] < i-k {
			j++
		}

		if j < nj && bindices[j] <= i+k {
			res = append(res, i)
		}
	}

	return res
}

// // Approach #1: Brute-force
// // Time: O(n*m + n*p + |a_indices| * |b_indices|), where n=len(s), m=len(a), p=len(b)
// // Space: O(|a_indices| + |b_indices|)
// func beautifulIndices(s string, a string, b string, k int) []int {
// 	abs := func(x int) int {
// 		if x < 0 {
// 			return -x
// 		}

// 		return x
// 	}

// 	n := len(s)

// 	findOccurrences := func(substr string) []int {
// 		// Find all occurrences of 'substr' in 's'
// 		var indices []int

// 		subn := len(substr)
// 		for i := range n - subn + 1 {
// 			if s[i:i+subn] == substr {
// 				indices = append(indices, i)
// 			}
// 		}

// 		return indices
// 	}

// 	aIndices := findOccurrences(a)
// 	bIndices := findOccurrences(b)

// 	result := []int{}

// 	for _, i := range aIndices {
// 		for _, j := range bIndices {
// 			if abs(j-i) <= k {
// 				result = append(result, i)

// 				break
// 			}
// 		}
// 	}

// 	return result
// }
