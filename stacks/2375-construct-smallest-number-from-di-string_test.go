package stacks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./stacks/ -run ^TestSmallestNumber$
func TestSmallestNumber(t *testing.T) {
	for _, tc := range []struct {
		pattern string
		num     string
	}{
		{pattern: "IIIDIDDD", num: "123549876"},
		{pattern: "DDD", num: "4321"},
	} {
		num := smallestNumber(tc.pattern)
		assert.Equal(t, tc.num, num)
	}
}
