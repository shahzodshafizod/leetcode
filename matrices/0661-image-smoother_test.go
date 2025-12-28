package matrices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./matrices/ -run ^TestImageSmoother$
func TestImageSmoother(t *testing.T) {
	for _, tc := range []struct {
		img    [][]int
		result [][]int
	}{
		{
			img:    [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}},
			result: [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
		},
		{
			img:    [][]int{{100, 200, 100}, {200, 50, 200}, {100, 200, 100}},
			result: [][]int{{137, 141, 137}, {141, 138, 141}, {137, 141, 137}},
		},
		{
			img:    [][]int{{1}},
			result: [][]int{{1}},
		},
		{
			img:    [][]int{{1, 2}, {3, 4}},
			result: [][]int{{2, 2}, {2, 2}},
		},
	} {
		result := imageSmoother(tc.img)
		assert.Equal(t, tc.result, result)
	}
}
