package intervals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./intervals/ -run ^TestCheckValidCuts$
func TestCheckValidCuts(t *testing.T) {
	for _, tc := range []struct {
		n          int
		rectangles [][]int
		valid      bool
	}{
		{n: 5, rectangles: [][]int{{1, 0, 5, 2}, {0, 2, 2, 4}, {3, 2, 5, 3}, {0, 4, 4, 5}}, valid: true},
		{n: 4, rectangles: [][]int{{0, 0, 1, 1}, {2, 0, 3, 4}, {0, 2, 2, 3}, {3, 0, 4, 3}}, valid: true},
		{n: 4, rectangles: [][]int{{0, 2, 2, 4}, {1, 0, 3, 2}, {2, 2, 3, 4}, {3, 0, 4, 2}, {3, 2, 4, 4}}, valid: false},
	} {
		valid := checkValidCuts(tc.n, tc.rectangles)
		assert.Equal(t, tc.valid, valid)
	}
}
