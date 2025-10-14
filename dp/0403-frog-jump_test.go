package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestCanCross$
func TestCanCross(t *testing.T) {
	for _, tc := range []struct {
		stones []int
		can    bool
	}{
		{stones: []int{0, 1, 3, 5, 6, 8, 12, 17}, can: true},
		{stones: []int{0, 1, 2, 3, 4, 8, 9, 11}, can: false},
	} {
		can := canCross(tc.stones)
		assert.Equal(t, tc.can, can)
	}
}
