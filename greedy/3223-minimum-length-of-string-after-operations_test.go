package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./greedy/ -run ^TestMinimumLength$
func TestMinimumLength(t *testing.T) {
	for _, tc := range []struct {
		s      string
		length int
	}{
		{s: "abaacbcbb", length: 5},
		{s: "aa", length: 2},
	} {
		length := minimumLength(tc.s)
		assert.Equal(t, tc.length, length)
	}
}

