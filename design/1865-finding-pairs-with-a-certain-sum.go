package design

// https://leetcode.com/problems/finding-pairs-with-a-certain-sum/

type FindSumPairs struct {
	nums1 []int
	nums2 []int
	freq  map[int]int
}

// Time: O(n2)
func NewFindSumPairs(nums1 []int, nums2 []int) FindSumPairs {
	freq := make(map[int]int, len(nums2))
	for _, num := range nums2 {
		freq[num]++
	}

	return FindSumPairs{
		nums1: nums1,
		nums2: nums2,
		freq:  freq,
	}
}

// Time: O(1)
func (f *FindSumPairs) Add(index int, val int) {
	f.freq[f.nums2[index]]--
	f.nums2[index] += val
	f.freq[f.nums2[index]]++
}

// Time: O(n1)
func (f *FindSumPairs) Count(tot int) int {
	count := 0
	for _, num := range f.nums1 {
		count += f.freq[tot-num]
	}

	return count
}

/**
 * Your FindSumPairs object will be instantiated and called as such:
 * obj := Constructor(nums1, nums2);
 * obj.Add(index,val);
 * param_2 := obj.Count(tot);
 */
