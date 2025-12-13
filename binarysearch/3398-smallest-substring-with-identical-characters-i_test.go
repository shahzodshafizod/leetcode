package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./binarysearch/ -run ^TestMinLengthI$
func TestMinLengthI(t *testing.T) {
	for _, tc := range []struct {
		s      string
		numOps int
		output int
	}{
		{s: "000001", numOps: 1, output: 2},
		{s: "0000", numOps: 2, output: 1},
		{s: "0101", numOps: 0, output: 1},
		{s: "1111", numOps: 1, output: 2},
		{s: "00", numOps: 0, output: 2},
		{s: "01", numOps: 0, output: 1},
		{s: "000", numOps: 1, output: 1},
	} {
		output := minLengthI(tc.s, tc.numOps)
		assert.Equal(t, tc.output, output)
	}
}
