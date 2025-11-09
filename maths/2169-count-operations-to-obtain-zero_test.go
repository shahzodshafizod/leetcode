package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./maths/ -run ^TestCountOperations$
func TestCountOperations(t *testing.T) {
	for _, tc := range []struct {
		num1 int
		num2 int
		ops  int
	}{
		{num1: 2, num2: 3, ops: 3},
		{num1: 10, num2: 10, ops: 1},
	} {
		ops := countOperations(tc.num1, tc.num2)
		assert.Equal(t, tc.ops, ops)
	}
}
