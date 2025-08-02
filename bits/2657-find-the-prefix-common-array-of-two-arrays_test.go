package bits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./bits/ -run ^TestTheFindThePrefixCommonArray$
func TestTheFindThePrefixCommonArray(t *testing.T) {
	for _, tc := range []struct {
		a []int
		b []int
		c []int
	}{
		{a: []int{1, 3, 2, 4}, b: []int{3, 1, 2, 4}, c: []int{0, 2, 3, 4}},
		{a: []int{2, 3, 1}, b: []int{3, 1, 2}, c: []int{0, 1, 3}},
	} {
		c := findThePrefixCommonArray(tc.a, tc.b)
		assert.Equal(t, tc.c, c)
	}
}
