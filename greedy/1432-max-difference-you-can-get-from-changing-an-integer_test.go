package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./greedy/ -run ^TestMaxDiff$
func TestMaxDiff(t *testing.T) {
	for _, tc := range []struct {
		num  int
		diff int
	}{
		{num: 555, diff: 888},
		{num: 9, diff: 8},
		{num: 1101057, diff: 8808050},
		{num: 123456, diff: 820000},
		{num: 123, diff: 820},
		{num: 111, diff: 888},
	} {
		diff := maxDiff(tc.num)
		assert.Equal(t, tc.diff, diff)
	}
}
