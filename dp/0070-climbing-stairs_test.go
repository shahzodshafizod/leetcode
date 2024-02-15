package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestClimbStairs$
func TestClimbStairs(t *testing.T) {
	for _, tc := range []struct {
		n            int
		combinations int
	}{
		{n: 2, combinations: 2},
		{n: 3, combinations: 3},
		{n: 10, combinations: 89},
	} {
		ways := climbStairs(tc.n)
		assert.Equal(t, tc.combinations, ways)
	}
}
