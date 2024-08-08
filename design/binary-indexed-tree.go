package design

/*
Move to the parent node in BIT
Flip the last set bit: idx -= idx & -idx
ex.: idx=7 and n=35
	00000111(7)
	00000111(7) - 00000111(7) & 11111001(-7) = 00000111(7) - 00000001(1) = 00000110(6)
	00000110(6) - 00000110(6) & 11111011(-6) = 00000110(6) - 00000010(2) = 00000100(4)
	00000100(4) - 00000100(4) & 11111111(-4) = 00000100(4) - 00000100(4) = 00000000(0)
	[7],[6],[4]

Move to the next relevant index in BIT
Add the last set bit: idx += idx & -idx
ex.: idx=7 and n=35
	00000111(7)
	00000111(7) + 00000111(7) & 11111001(-7) = 00000111(7) + 00000001(1) = 00001000(8)
	00001000(8) + 00001000(8) & 11111111(-8) = 00001000(8) + 00001000(8) = 00010000(16)
	00010000(16) + 00010000(16) & 00010000(-16) = 00010000(16) + 00010000(16) = 00100000(32)
	00100000(32) + 00100000(32) & 00100000(-32) = 00100000(32) + 00100000(32) = 00100000(64) -> 64 > 35
	[7],[8],[16],[32]
*/

type BIT interface {
	Update(index int, delta int)
	Query(index int) int
}

type bit []int

func NewBIT(n int) BIT {
	var b = bit(make([]int, n+1))
	return &b
}

func (b *bit) Update(index int, delta int) {
	for idx, n := index, len(*b); idx < n; idx += idx & -idx { // add the last set bit
		(*b)[idx] += delta
	}
}

func (b *bit) Query(index int) int {
	sum := 0
	for idx := index; idx > 0; idx -= idx & -idx { // flip the last set bit
		sum += (*b)[idx]
	}
	return sum
}
