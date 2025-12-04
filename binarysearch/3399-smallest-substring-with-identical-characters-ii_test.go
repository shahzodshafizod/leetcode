package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./binarysearch/ -run ^TestMinLength$
func TestMinLength(t *testing.T) {
	for _, tc := range []struct {
		s      string
		numOps int
		length int
	}{
		{s: "000001", numOps: 1, length: 2},
		{s: "0000", numOps: 2, length: 1},
		{s: "0101", numOps: 0, length: 1},
		{s: "000000", numOps: 2, length: 2},
		{s: "111111", numOps: 1, length: 3},
		{s: "00", numOps: 0, length: 2},
		{s: "0", numOps: 0, length: 1},
		{s: "1111100000", numOps: 2, length: 2},
		{s: "0110", numOps: 1, length: 2},
	} {
		length := minLength(tc.s, tc.numOps)
		assert.Equal(t, tc.length, length)
	}
}
