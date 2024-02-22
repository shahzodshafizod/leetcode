package bits

// https://leetcode.com/problems/reverse-bits/

func reverseBits(num uint32) uint32 {
	var reversed uint32 = 0
	var power uint32 = 31
	for num > 0 {
		reversed += (num & 1) << power
		num >>= 1
		power--
	}
	return reversed
}

// func reverseBits(num uint32) uint32 {
//     var reversed uint32 = 0
// 	for i := 0; i < 32; i++ {
// 		reversed = (reversed << 1) + num&1
// 		num >>= 1
// 	}
// 	return reversed
// }
