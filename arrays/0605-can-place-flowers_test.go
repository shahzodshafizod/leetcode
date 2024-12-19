package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./arrays/ -run ^TestCanPlaceFlowers$
func TestCanPlaceFlowers(t *testing.T) {
	for _, tc := range []struct {
		flowerbed []int
		n         int
		canPlant  bool
	}{
		{flowerbed: []int{1, 0, 0, 0, 1}, n: 1, canPlant: true},
		{flowerbed: []int{1, 0, 0, 0, 1}, n: 2, canPlant: false},
		{flowerbed: []int{0, 0, 1, 0, 0}, n: 1, canPlant: true},
		{flowerbed: []int{0, 0, 1, 0, 0}, n: 2, canPlant: true},
		{flowerbed: []int{0, 0, 1, 0, 0}, n: 3, canPlant: false},
		{flowerbed: []int{1, 0, 0, 0, 1}, n: 1, canPlant: true},
		{flowerbed: []int{1, 0, 0, 0, 1}, n: 2, canPlant: false},
		{flowerbed: []int{1, 0, 0, 0, 1}, n: 1, canPlant: true},
		{flowerbed: []int{0, 0, 1, 0, 0}, n: 1, canPlant: true},
		{flowerbed: []int{0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0}, n: 4, canPlant: true},
		{flowerbed: []int{1, 0, 0, 0, 0, 0, 1}, n: 2, canPlant: true},
		{flowerbed: []int{0}, n: 1, canPlant: true},
		{flowerbed: []int{1, 0, 0, 0, 0, 1}, n: 2, canPlant: false},
		{flowerbed: []int{1, 0, 0, 1, 0, 0, 0, 1}, n: 1, canPlant: true},
	} {
		canPlant := canPlaceFlowers(tc.flowerbed, tc.n)
		assert.Equal(t, tc.canPlant, canPlant)
	}
}
