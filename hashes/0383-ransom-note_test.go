package hashes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./hashes/ -run ^TestCanConstruct0383$
func TestCanConstruct0383(t *testing.T) {
	for _, tc := range []struct {
		ransomNote string
		magazine   string
		can        bool
	}{
		{ransomNote: "a", magazine: "b", can: false},
		{ransomNote: "aa", magazine: "ab", can: false},
		{ransomNote: "aa", magazine: "aab", can: true},
	} {
		can := canConstruct0383(tc.ransomNote, tc.magazine)
		assert.Equal(t, tc.can, can)
	}
}
