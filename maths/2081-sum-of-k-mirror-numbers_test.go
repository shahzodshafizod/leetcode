package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./maths/ -run ^TestKMirror$
func TestKMirror(t *testing.T) {
	for _, tc := range []struct {
		k   int
		n   int
		sum int64
	}{
		{k: 2, n: 5, sum: 25},
		{k: 3, n: 7, sum: 499},
		{k: 7, n: 17, sum: 20379000},
		{k: 4, n: 5, sum: 66},
	} {
		sum := kMirror(tc.k, tc.n)
		assert.Equal(t, tc.sum, sum)
	}
}
