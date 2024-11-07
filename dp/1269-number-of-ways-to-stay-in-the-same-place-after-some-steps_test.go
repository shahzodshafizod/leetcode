package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestNumWays$
func TestNumWays(t *testing.T) {
	for _, tc := range []struct {
		steps  int
		arrLen int
		ways   int
	}{
		{steps: 3, arrLen: 2, ways: 4},
		{steps: 2, arrLen: 4, ways: 2},
		{steps: 4, arrLen: 2, ways: 8},
		// {steps: 500, arrLen: 1000000, ways: 374847123},
	} {
		ways := numWays(tc.steps, tc.arrLen)
		assert.Equal(t, tc.ways, ways)
	}
}
