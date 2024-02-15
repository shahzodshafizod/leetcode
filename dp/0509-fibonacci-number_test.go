package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestFib$
func TestFib(t *testing.T) {
	for _, tc := range []struct {
		n         int
		fibonacci int
	}{
		{n: 2, fibonacci: 1},
		{n: 3, fibonacci: 2},
		{n: 4, fibonacci: 3},
		{n: 10, fibonacci: 55},
	} {
		fibonacci := fib(tc.n)
		assert.Equal(t, tc.fibonacci, fibonacci)
	}
}
