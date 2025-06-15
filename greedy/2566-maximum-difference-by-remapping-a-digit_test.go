package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./greedy/ -run ^TestMinMaxDifference$
func TestMinMaxDifference(t *testing.T) {
	for _, tc := range []struct {
		num  int
		diff int
	}{
		{num: 11891, diff: 99009},
		{num: 90, diff: 99},
	} {
		diff := minMaxDifference(tc.num)
		assert.Equal(t, tc.diff, diff)
	}
}
