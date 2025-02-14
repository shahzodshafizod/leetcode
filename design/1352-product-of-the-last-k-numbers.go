package design

// https://leetcode.com/problems/product-of-the-last-k-numbers/

type ProductOfNumbers struct {
	presum []int
	size   int
}

func NewProductOfNumbers() ProductOfNumbers {
	return ProductOfNumbers{
		presum: []int{1},
		size:   0,
	}
}

func (p *ProductOfNumbers) Add(num int) {
	if num == 0 {
		p.presum = []int{1}
		p.size = 0
		return
	}
	p.presum = append(p.presum, p.presum[p.size]*num)
	p.size++
}

func (p *ProductOfNumbers) GetProduct(k int) int {
	if k > p.size {
		return 0
	}
	return p.presum[p.size] / p.presum[p.size-k]
}

/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */
