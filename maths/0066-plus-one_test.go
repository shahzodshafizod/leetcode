package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./maths/ -run ^TestPlusOne$
func TestPlusOne(t *testing.T) {
	for _, tc := range []struct {
		digits   []int
		modified []int
	}{
		{digits: []int{1, 2, 3}, modified: []int{1, 2, 4}},
		{digits: []int{4, 3, 2, 1}, modified: []int{4, 3, 2, 2}},
		{digits: []int{9}, modified: []int{1, 0}},
		{digits: []int{9, 9, 9, 9}, modified: []int{1, 0, 0, 0, 0}},
		{digits: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, modified: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 1}},
		{digits: []int{6, 1, 4, 5, 3, 9, 0, 1, 9, 5, 1, 8, 6, 7, 0, 5, 5, 4, 3}, modified: []int{6, 1, 4, 5, 3, 9, 0, 1, 9, 5, 1, 8, 6, 7, 0, 5, 5, 4, 4}},
		{digits: []int{9, 9}, modified: []int{1, 0, 0}},
	} {
		modified := plusOne(tc.digits)
		assert.Equal(t, tc.modified, modified)
	}
}
