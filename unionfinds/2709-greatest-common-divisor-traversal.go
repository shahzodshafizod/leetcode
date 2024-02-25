package unionfinds

// https://leetcode.com/problems/greatest-common-divisor-traversal/

type UnionFind struct {
	par   []int
	size  []int
	count int
}

func NewUnionFind(n int) *UnionFind {
	var par = make([]int, n)
	var size = make([]int, n)
	for i := 0; i < n; i++ {
		par[i] = i
		size[i] = 1
	}
	return &UnionFind{par: par, size: size, count: n}
}

func (u *UnionFind) Find(x int) int {
	if u.par[x] != x {
		u.par[x] = u.Find(u.par[x])
	}
	return u.par[x]
}

func (u *UnionFind) Union(x int, y int) {
	var px, py = u.Find(x), u.Find(y)
	if px == py {
		return
	}
	if u.size[px] < u.size[py] {
		u.par[px] = py
		u.size[py] += u.size[px]
	} else {
		u.par[py] = px
		u.size[px] += u.size[py]
	}
	u.count--
}

func canTraverseAllPairs(nums []int) bool {
	var uf = NewUnionFind(len(nums))
	var factors = make(map[int]int)
	var factor int
	for idx, num := range nums {
		factor = 2
		for factor*factor <= num {
			if num%factor == 0 {
				if index, exists := factors[factor]; exists {
					uf.Union(idx, index)
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
				uf.Union(idx, index)
			} else {
				factors[num] = idx
			}
		}
	}
	return uf.count == 1
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
// 	var uf = NewUnionFind(max + 1)
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
