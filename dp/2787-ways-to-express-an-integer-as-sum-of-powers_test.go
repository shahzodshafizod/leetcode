package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestNumberOfWays$
func TestNumberOfWays(t *testing.T) {
	for _, tc := range []struct {
		n    int
		x    int
		ways int
	}{
		{n: 10, x: 2, ways: 1},
		{n: 4, x: 1, ways: 2},
	} {
		ways := numberOfWays(tc.n, tc.x)
		assert.Equal(t, tc.ways, ways)
	}
}
