package backtracking

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./backtracking/ -run ^TestConstructDistancedSequence$
func TestConstructDistancedSequence(t *testing.T) {
	for _, tc := range []struct {
		n   int
		seq []int
	}{
		{n: 3, seq: []int{3, 1, 2, 3, 2}},
		{n: 5, seq: []int{5, 3, 1, 4, 3, 5, 2, 4, 2}},
	} {
		seq := constructDistancedSequence(tc.n)
		assert.Equal(t, tc.seq, seq)
	}
}
