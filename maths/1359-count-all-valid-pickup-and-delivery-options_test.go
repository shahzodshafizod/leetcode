package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./maths/ -run ^TestCountOrders$
func TestCountOrders(t *testing.T) {
	for _, tc := range []struct {
		n       int
		options int
	}{
		{n: 1, options: 1},
		{n: 2, options: 6},
		{n: 3, options: 90},
		// {n: 500, options: 764678010},
	} {
		options := countOrders(tc.n)
		assert.Equal(t, tc.options, options)
	}
}
