package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestMaxEnvelopes$
func TestMaxEnvelopes(t *testing.T) {
	for _, tc := range []struct {
		envelopes [][]int
		count     int
	}{
		{envelopes: [][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}, count: 3},
		{envelopes: [][]int{{1, 1}, {1, 1}, {1, 1}}, count: 1},
	} {
		count := maxEnvelopes(tc.envelopes)
		assert.Equal(t, tc.count, count)
	}
}
