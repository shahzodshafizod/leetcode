package bits

// https://leetcode.com/problems/number-of-steps-to-reduce-a-number-in-binary-representation-to-one/

// time: O(n)
// space: O(1)
func numSteps(s string) int {
	operations, carry := 0, 0

	var bit int

	for idx := len(s) - 1; idx > 0; idx-- {
		bit = int(s[idx]-'0') ^ carry
		carry |= bit
		operations += 1 + bit
	}

	return operations + carry
}

// // time: O(n)
// // space: O(1)
// func numSteps(s string) int {
//     var operations, carry = 0, 0
// 	for idx := len(s) - 1; idx > 0; idx-- {
// 		operations++
// 		if (int(s[idx]-'0')+carry)&1 == 1 {
// 			operations++
// 			carry = 1
// 		}
// 	}
// 	return operations + carry
// }

// // time: O(n x n)
// // space: O(n)
// func numSteps(s string) int {
// 	var n = len(s)
// 	var bits = make([]bool, n)
// 	for idx := range bits {
// 		bits[idx] = s[idx] == '1'
// 	}
// 	var operations = 0
// 	var one int
// 	for idx := n - 1; idx > 0; idx-- {
// 		operations++
// 		if bits[idx] {
// 			operations++
// 			for one = idx; one >= 0 && bits[one]; one-- {
// 				bits[one] = false
// 			}
// 			if one >= 0 {
// 				bits[one] = true
// 			}
// 		}
// 	}
// 	if !bits[0] {
// 		operations++
// 	}
// 	return operations
// }
