package bits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./bits/ -run ^TestCountBits$
func TestCountBits(t *testing.T) {
	for _, tc := range []struct {
		n    int
		ones []int
	}{
		{n: 0, ones: []int{0}},
		{n: 2, ones: []int{0, 1, 1}},
		{n: 5, ones: []int{0, 1, 1, 2, 1, 2}},
		{n: 14, ones: []int{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3}},
	} {
		ones := countBits(tc.n)
		assert.Equal(t, tc.ones, ones)
	}
}
