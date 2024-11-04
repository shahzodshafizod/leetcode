package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./strings/ -run ^TestRotateString$
func TestRotateString(t *testing.T) {
	for _, tc := range []struct {
		s    string
		goal string
		can  bool
	}{
		{s: "a", goal: "a", can: true},
		{s: "bcad", goal: "cdba", can: false},
		{s: "bcad", goal: "abcd", can: false},
		{s: "abcd", goal: "cdeab", can: false},
		{s: "abcde", goal: "abce", can: false},
		{s: "abcde", goal: "cdeab", can: true},
		{s: "abcde", goal: "abced", can: false},
		{s: "aaaaaaaaaa", goal: "aaaaaaaaaaaa", can: false},
		{s: "bbbacddceeb", goal: "ceebbbbacdd", can: true},
		{s: "bbbacddceeb", goal: "ceebbbbacdd", can: true},
		{s: "abcdebcdefin", goal: "abcdebcdefin", can: true},
		{s: "aaaaaaaaaaaaaaabc", goal: "caaaaaaaaaaaaaaab", can: true},
		{s: "dvadererrerdvaddf", goal: "rerdvaddfdvaderer", can: true},
	} {
		can := rotateString(tc.s, tc.goal)
		assert.Equal(t, tc.can, can)
	}
}
