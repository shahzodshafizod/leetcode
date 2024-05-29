package bits

// https://leetcode.com/problems/number-of-steps-to-reduce-a-number-in-binary-representation-to-one/

// time: O(n)
// space: O(1)
func numSteps(s string) int {
	var operations, carry = 0, 0
	for idx := len(s) - 1; idx > 0; idx-- {
		operations++
		if (int(s[idx]-'0')+carry)&1 == 1 {
			operations++
			carry = 1
		}
	}
	return operations + carry
}

// func numSteps(s string) int {
// 	var operations = 0
// 	var carry = 0
// 	for idx := len(s) - 1; idx > 0; idx-- {
// 		switch (int(s[idx]-'0') + carry) & 1 {
// 		case 1:
// 			operations += 2
// 			carry = 1
// 		case 0:
// 			operations++
// 		}
// 	}
// 	return operations + carry
// }
