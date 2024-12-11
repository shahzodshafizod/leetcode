package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./binarysearch/ -run ^TestFirstBadVersion$
func TestFirstBadVersion(t *testing.T) {
	for _, tc := range []struct {
		n       int
		setting int
		first   int
	}{
		{n: 5, setting: 4, first: 4},
		{n: 1, setting: 1, first: 1},
		{n: 3, setting: 2, first: 2},
	} {
		FIRST__BAD__VERSION__ = tc.setting
		first := firstBadVersion(tc.n)
		assert.Equal(t, tc.first, first)
	}
}
