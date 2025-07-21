package learning

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./learning/ -run ^TestTricks$
func TestTricks(t *testing.T) {
	assert := assert.New(t)

	// 1. Check if a bit is set
	assert.Equal((0b_0101_1000>>2)&1, 0) // starting from 0 from rightmost, the second bit isn't set
	assert.Equal((0b_0101_1000>>3)&1, 1) // starting from 0 from rightmost, the third bit is set

	// 2. check if a number is even
	assert.Equal(46&1 == 0, true)
	assert.Equal(47&1 == 0, false)

	// 3. Check if an integer is a power of 2
	// A power of two is a number that has only a single bit in it
	for num, is := range map[int]bool{32: true, 47: false, 0: false} {
		// in the original source: n && !(n & (n - 1))
		assert.Equal(num != 0 && num&(num-1) == 0, is)
	}
}

// go test -v -count=1 ./learning/ -run ^TestBits$
func TestBits(t *testing.T) {
	// bitwise operators: "&", "|", "^", "~", "<<", ">>"
	assert := assert.New(t)

	// AND
	// 0b01011000 &
	// 0b01010111 =
	// 0b01010000
	assert.Equal(0b01011000&0b01010111, 0b01010000)

	// OR
	// 0b01011000 |
	// 0b01010111 =
	// 0b01011111
	assert.Equal(0b01011000|0b01010111, 0b01011111)

	// XOR (exclusive OR)
	// 0b01011000 ^
	// 0b01010111 =
	// 0b00001111
	assert.Equal(0b01011000^0b01010111, 0b00001111)

	// NEGATION
	// ^0 ^ x is equal to ^x
	// ^0 doesn't change the value of ^x in ^ operation in a similar manner
	// that +0 doesn't change the value of +x in + operation.
	//            ! 0b01010101 =
	//         ^0 ^ 0b01010101 =
	// 0b11111111 ^ 0b01010101 =
	//              0b10101010
	// for _, tc := range []struct {
	// 	num     int
	// 	negated int
	// }{
	// 	{num: 0b01010101, negated: 0b10101010}, // 170
	// 	{num: 0b00011001, negated: 0b11100110}, // 230
	// } {
	// 	assert.Equal((tc.num^^tc.num)^tc.num, tc.negated)
	// 	assert.Equal((1<<32-1)^tc.num, tc.negated)
	// }

	// SHIFT
	assert.Equal(2<<1, 4)
	assert.Equal(5>>1, 2)
}
