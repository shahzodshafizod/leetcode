package bits

// https://leetcode.com/problems/number-of-steps-to-reduce-a-number-to-zero/

func numberOfSteps(num int) int {
	steps := 0
	for num > 0 {
		if num&1 == 1 {
			num ^= 1
		} else {
			num >>= 1
		}
		steps++
	}
	return steps
	// if num == 0 {
	// 	return 0
	// }
	// var steps = 0
	// for num > 0 {
	// 	steps += 1 + num&1
	// 	num >>= 1
	// }
	// return steps-1
}
