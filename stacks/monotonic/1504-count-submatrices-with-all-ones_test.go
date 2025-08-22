package monotonic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./stacks/monotonic/ -run ^TestNumSubmat$
func TestNumSubmat(t *testing.T) {
	for _, tc := range []struct {
		mat [][]int
		cnt int
	}{
		{mat: [][]int{{1, 0, 1}, {1, 1, 0}, {1, 1, 0}}, cnt: 13},
		{mat: [][]int{{0, 1, 1, 0}, {0, 1, 1, 1}, {1, 1, 1, 0}}, cnt: 24},
	} {
		cnt := numSubmat(tc.mat)
		assert.Equal(t, tc.cnt, cnt)
	}
}
