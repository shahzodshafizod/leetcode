package bits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./bits/ -run ^TestHasAlternatingBits$
func TestHasAlternatingBits(t *testing.T) {
	for _, tc := range []struct {
		n      int
		result bool
	}{
		{n: 5, result: true},
		{n: 7, result: false},
		{n: 11, result: false},
		{n: 10, result: true},
		{n: 1, result: true},
		{n: 2, result: true},
		{n: 3, result: false},
		{n: 4, result: false},
		{n: 21, result: true},
		{n: 42, result: true},
		{n: 85, result: true},
	} {
		result := hasAlternatingBits(tc.n)
		assert.Equal(t, tc.result, result)
	}
}
