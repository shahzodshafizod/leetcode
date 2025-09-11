package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./strings/ -run ^TestSortVowels$
func TestSortVowels(t *testing.T) {
	for _, tc := range []struct {
		s      string
		sorted string
	}{
		{s: "lEetcOde", sorted: "lEOtcede"},
		{s: "lYmpH", sorted: "lYmpH"},
	} {
		sorted := sortVowels(tc.s)
		assert.Equal(t, tc.sorted, sorted)
	}
}
