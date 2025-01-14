package bits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./bits/ -run ^TestTheFindThePrefixCommonArray$
func TestTheFindThePrefixCommonArray(t *testing.T) {
	for _, tc := range []struct {
		A []int
		B []int
		C []int
	}{
		{A: []int{1, 3, 2, 4}, B: []int{3, 1, 2, 4}, C: []int{0, 2, 3, 4}},
		{A: []int{2, 3, 1}, B: []int{3, 1, 2}, C: []int{0, 1, 3}},
	} {
		C := findThePrefixCommonArray(tc.A, tc.B)
		assert.Equal(t, tc.C, C)
	}
}
