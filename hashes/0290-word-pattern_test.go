package hashes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./hashes/ -run ^TestWordPattern$
func TestWordPattern(t *testing.T) {
	for _, tc := range []struct {
		pattern string
		s       string
		same    bool
	}{
		{pattern: "abba", s: "dog cat cat dog", same: true},
		{pattern: "abba", s: "dog cat cat fish", same: false},
		{pattern: "aaaa", s: "dog cat cat dog", same: false},
	} {
		same := wordPattern(tc.pattern, tc.s)
		assert.Equal(t, tc.same, same)
	}
}
