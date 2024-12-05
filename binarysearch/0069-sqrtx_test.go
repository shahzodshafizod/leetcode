package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./binarysearch/ -run ^TestMySqrt$
func TestMySqrt(t *testing.T) {
	for _, tc := range []struct {
		x    int
		root int
	}{
		{x: 0, root: 0},
		{x: 1, root: 1},
		{x: 4, root: 2},
		{x: 8, root: 2},
		{x: 2147395599, root: 46339},
		{x: 2147395600, root: 46340},
	} {
		root := mySqrt(tc.x)
		assert.Equal(t, tc.root, root)
	}
}
