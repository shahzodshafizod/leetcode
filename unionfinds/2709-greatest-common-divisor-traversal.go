package unionfinds

import "github.com/shahzodshafizod/leetcode/pkg"

// https://leetcode.com/problems/greatest-common-divisor-traversal/

func canTraverseAllPairs(nums []int) bool {
	n := len(nums)
	uf := pkg.NewUnionFindRanked(n)
	factors := make(map[int]int)
	var factor int
	for idx, num := range nums {
		factor = 2
		for factor*factor <= num {
			if num%factor == 0 {
				if index, exists := factors[factor]; exists {
					if uf.Union(idx, index) {
						n--
					}
				} else {
					factors[factor] = idx
				}
				for num%factor == 0 {
					num /= factor
				}
			}
			factor++
		}
		if num > 1 {
			if index, exists := factors[num]; exists {
				if uf.Union(idx, index) {
					n--
				}
			} else {
				factors[num] = idx
			}
		}
	}
	return n == 1
}

// // Ghalberi Eratosfen
// func canTraverseAllPairs(nums []int) bool {
// 	// edge case
// 	if len(nums) == 1 {
// 		return true
// 	}
// 	var max = 100000
// 	var sieve = make([]int, max+1)
// 	for i := 2; i < max; i++ {
// 		if sieve[i] == 0 {
// 			for j := i; j < max; j += i {
// 				sieve[j] = i
// 			}
// 		}
// 	}
// 	var uf = pkg.NewUnionFindRanked(max + 1)
// 	for _, num := range nums {
// 		// edge case
// 		if num == 1 {
// 			return false
// 		}
// 		var temp = num
// 		var factor = sieve[temp]
// 		for temp > 1 && factor != 0 {
// 			uf.Union(num, factor)
// 			temp /= factor
// 			factor = sieve[temp]
// 		}
// 	}
// 	var root = uf.Find(nums[0])
// 	for _, num := range nums {
// 		if uf.Find(num) != root {
// 			return false
// 		}
// 	}
// 	return true
// }
