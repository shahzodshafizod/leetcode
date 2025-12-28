package stacks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./stacks/ -run ^TestCalPoints$
func TestCalPoints(t *testing.T) {
	for _, tc := range []struct {
		operations []string
		result     int
	}{
		{
			operations: []string{"5", "2", "C", "D", "+"},
			result:     30,
		},
		{
			operations: []string{"5", "-2", "4", "C", "D", "9", "+", "+"},
			result:     27,
		},
		{
			operations: []string{"1", "C"},
			result:     0,
		},
		{
			operations: []string{"1"},
			result:     1,
		},
		{
			operations: []string{"1", "2", "+", "D", "C"},
			result:     6,
		},
		{
			operations: []string{"-60", "D", "-36", "30", "13", "C", "C", "-33", "53", "79"},
			result:     -117,
		},
	} {
		result := calPoints(tc.operations)
		assert.Equal(t, tc.result, result)
	}
}
