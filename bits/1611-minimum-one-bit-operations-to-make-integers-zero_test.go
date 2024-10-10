package bits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./bits/ -run ^TestMinimumOneBitOperations$
func TestMinimumOneBitOperations(t *testing.T) {
	for _, tc := range []struct {
		n          int
		operations int
	}{
		{n: 3, operations: 2},
		{n: 6, operations: 4},
		{n: 333, operations: 393},
	} {
		operations := minimumOneBitOperations(tc.n)
		assert.Equal(t, tc.operations, operations)
	}
}
