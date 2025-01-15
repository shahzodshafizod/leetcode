package bits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./bits/ -run ^TestMinimizeXor$
func TestMinimizeXor(t *testing.T) {
	for _, tc := range []struct {
		num1 int
		num2 int
		x    int
	}{
		{num1: 3, num2: 5, x: 3},
		{num1: 1, num2: 12, x: 3},
	} {
		x := minimizeXor(tc.num1, tc.num2)
		assert.Equal(t, tc.x, x)
	}
}
