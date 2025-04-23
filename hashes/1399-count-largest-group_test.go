package hashes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./hashes/ -run ^TestCountLargestGroup$
func TestCountLargestGroup(t *testing.T) {
	for _, tc := range []struct {
		n     int
		count int
	}{
		{n: 13, count: 4},
		{n: 2, count: 2},
	} {
		count := countLargestGroup(tc.n)
		assert.Equal(t, tc.count, count)
	}
}
