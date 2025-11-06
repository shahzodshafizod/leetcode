package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./strings/ -run ^TestHasSameDigits$
func TestHasSameDigits(t *testing.T) {
	for _, tc := range []struct {
		s    string
		same bool
	}{
		{s: "3902", same: true},
		{s: "34789", same: false},
	} {
		same := hasSameDigits(tc.s)
		assert.Equal(t, tc.same, same)
	}
}
