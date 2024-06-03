package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./greedy/ -run ^TestAppendCharacters$
func TestAppendCharacters(t *testing.T) {
	for _, tc := range []struct {
		s   string
		t   string
		min int
	}{
		{s: "coaching", t: "coding", min: 4},
		{s: "coacdhing", t: "coding", min: 0},
		{s: "abcde", t: "a", min: 0},
		{s: "z", t: "abcde", min: 5},
		{s: "abba", t: "abba", min: 0},
	} {
		min := appendCharacters(tc.s, tc.t)
		assert.Equal(t, tc.min, min)
	}
}
