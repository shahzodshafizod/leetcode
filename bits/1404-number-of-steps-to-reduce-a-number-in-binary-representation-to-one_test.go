package bits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./bits/ -run ^TestNumSteps$
func TestNumSteps(t *testing.T) {
	for _, tc := range []struct {
		s     string
		steps int
	}{
		{s: "1101", steps: 6},
		{s: "10", steps: 1},
		{s: "1", steps: 0},
	} {
		steps := numSteps(tc.s)
		assert.Equal(t, tc.steps, steps)
	}
}
