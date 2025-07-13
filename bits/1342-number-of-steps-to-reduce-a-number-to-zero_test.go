package bits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./bits/ -run ^TestNumberOfSteps$
func TestNumberOfSteps(t *testing.T) {
	for _, tc := range []struct {
		num int
		steps int
	}{
		{num: 14, steps: 6},
		{num: 8, steps: 4},
		{num: 123, steps: 12},
	} {
		steps := numberOfSteps(tc.num)
		assert.Equal(t, tc.steps, steps)
	}
}
