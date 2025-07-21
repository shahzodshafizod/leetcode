package maths

// https://leetcode.com/problems/count-all-valid-pickup-and-delivery-options/

/*
Explanation #1
number of permutations: 2! (P1D1, D1,P1)
half of total permutations is invalid: 2**n (D1P1)

n=1: (P1D1)/2 = 2!/2 = 2/2 = 1
n=2: (P1D1)/2 * (P2D2)/2 = 4!/2^2 = 24/4 = 6
n=3: (P1D1)/2 * (P2D2)/2 * (P3D3)/2 = 6!/2^3 = 720/8 = 90
*/

/*
Explanation #2
When n=1, we have 2 spots to fill _ _ with P1 & D1, options is 1
When n=2, we have 4 spots: _ _ _ _.
    For the 1st spot, we have 2 choices: P1, P2.
    For the next 3 spots, if we choose P1 for the spot,
    then we will have 3 places to choose D1.
    And for the rest of the 2 spaces we will fill P2 and D2.
    P1 x 3
        P1,D1,P2,D2
        P1,P2,D1,D2
        P1,P2,D2,D1
    P2 x 3
        P2,D2,P1,D1
        P2,P1,D2,D1
        P2,P1,D1,D2
    options is 2x3 = 6
When n=3, we have 6 spots: _ _ _ _ _ _.
    For the 1st spot, we have 3 choices, P1, P2, P3.
    For the next 5 spots, if we choose P1 for the spot,
    then we will have 5 places to choose D1.
    And for the rest of 4 spots, we already have 6 options.
    options is 3x5x6 = 90
*/

// Approach: Iterative
// Time: O(n)
// Space: O(1)
func countOrders(n int) int {
	const MOD int = 1e9 + 7
	options := 1
	for spots := n << 1; spots > 2; spots -= 2 {
		options = (options * (spots >> 1) * (spots - 1)) % MOD
	}
	// for ; n > 1; n-- {
	// 	options = (options * n * (2*n - 1)) % MOD
	// }
	return options
}

// // Approach: Recursive
// // Time: O(n)
// // Space: O(n)
// func countOrders(n int) int {
// 	const MOD int = 1e9 + 7
// 	var calcOptions func(n int) int
// 	calcOptions = func(n int) int {
// 		if n == 1 {
// 			return 1
// 		}
// 		return n * (2*n - 1) * calcOptions(n-1) % MOD
// 	}
// 	return calcOptions(n)
// 	// const MOD int = 1e9 + 7
// 	// var calcOptions func(spots int) int
// 	// calcOptions = func(spots int) int {
// 	// 	if spots == 2 {
// 	// 		return 1
// 	// 	}
// 	// 	return spots / 2 * (spots - 1) * calcOptions(spots-2) % MOD
// 	// }
// 	// return calcOptions(n * 2)
// }
