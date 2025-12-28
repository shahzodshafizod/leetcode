package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./strings/ -run ^TestMinDeletionSize$
func TestMinDeletionSize(t *testing.T) {
	for _, tc := range []struct {
		strs   []string
		result int
	}{
		{
			strs:   []string{"cba", "daf", "ghi"},
			result: 1,
		},
		{
			strs:   []string{"a", "b"},
			result: 0,
		},
		{
			strs:   []string{"zyx", "wvu", "tsr"},
			result: 3,
		},
		{
			strs:   []string{"abc", "bce", "cae"},
			result: 1,
		},
		{
			strs:   []string{"a"},
			result: 0,
		},
		{
			strs:   []string{"abc"},
			result: 0,
		},
		{
			strs:   []string{"rrjk", "furt", "guzm"},
			result: 2,
		},
	} {
		result := minDeletionSize(tc.strs)
		assert.Equal(t, tc.result, result)
	}
}
