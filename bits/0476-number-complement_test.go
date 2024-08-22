package bits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./bits/ -run ^TestFindComplement$
func TestFindComplement(t *testing.T) {
	for _, tc := range []struct {
		num     int
		flipped int
	}{
		{num: 5, flipped: 2},
		{num: 1, flipped: 0},
	} {
		flipped := findComplement(tc.num)
		assert.Equal(t, tc.flipped, flipped)
	}
}
