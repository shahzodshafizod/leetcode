package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestRearrangeSticks$
func TestRearrangeSticks(t *testing.T) {
	for _, tc := range []struct {
		n    int
		k    int
		ways int
	}{
		{n: 3, k: 2, ways: 3},
		{n: 5, k: 5, ways: 1},
		{n: 20, k: 11, ways: 647427950},
		{n: 978, k: 575, ways: 103376920},
	} {
		ways := rearrangeSticks(tc.n, tc.k)
		assert.Equal(t, tc.ways, ways)
	}
}
