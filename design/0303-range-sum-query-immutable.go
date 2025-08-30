package design

// https://leetcode.com/problems/range-sum-query-immutable/

type NumArray struct {
	prefix []int
}

func NewNumArray(nums []int) NumArray {
	n := len(nums)
	prefix := make([]int, n+1)

	for idx := range n {
		prefix[idx+1] = prefix[idx] + nums[idx]
	}

	return NumArray{prefix: prefix}
}

func (n *NumArray) SumRange(left int, right int) int {
	return n.prefix[right+1] - n.prefix[left]
}

/*
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
