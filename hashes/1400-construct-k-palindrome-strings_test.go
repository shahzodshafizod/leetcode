package hashes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./hashes/ -run ^TestCanConstruct$
func TestCanConstruct(t *testing.T) {
	for _, tc := range []struct {
		s   string
		k   int
		can bool
	}{
		{s: "annabelle", k: 2, can: true},
		{s: "leetcode", k: 3, can: false},
		{s: "true", k: 4, can: true},
	} {
		can := canConstruct(tc.s, tc.k)
		assert.Equal(t, tc.can, can)
	}
}
