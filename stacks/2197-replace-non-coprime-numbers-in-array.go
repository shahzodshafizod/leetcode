package stacks

// https://leetcode.com/problems/replace-non-coprime-numbers-in-array/

func replaceNonCoprimes(nums []int) []int {
	gcd := func(a int, b int) int {
		for b > 0 {
			a, b = b, a%b
		}

		return a
	}
	stack := make([]int, len(nums))

	var size, g, lcm int
	for idx := range nums {
		lcm = nums[idx]
		for size > 0 {
			g = gcd(stack[size-1], lcm)
			if g == 1 {
				break
			}

			lcm = stack[size-1] * lcm / g
			size--
		}

		stack[size] = lcm
		size++
	}

	return stack[:size]
}
