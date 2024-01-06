package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dsa/arrays/ -run ^TestTrap$
func TestTrap(t *testing.T) {
	for _, data := range []struct {
		height []int
		total  int
	}{
		{height: []int{0, 1, 0, 2, 1, 0, 3, 1, 0, 1, 2}, total: 8},
		{height: []int{}, total: 0},
		{height: []int{3}, total: 0},
		{height: []int{3, 4, 3}, total: 0},
		{height: []int{5, 0, 3, 0, 0, 0, 2, 3, 4, 2}, total: 20},
		{height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, total: 6},
		{height: []int{4, 2, 0, 3, 2, 5}, total: 9},
	} {
		total := trap(data.height)
		assert.Equal(t, data.total, total)
	}
}
