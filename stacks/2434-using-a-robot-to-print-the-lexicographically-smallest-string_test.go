package stacks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./stacks/ -run ^TestRobotWithString$
func TestRobotWithString(t *testing.T) {
	for _, tc := range []struct {
		s string
		p string
	}{
		{s: "zza", p: "azz"},
		{s: "bac", p: "abc"},
		{s: "bdda", p: "addb"},
	} {
		p := robotWithString(tc.s)
		assert.Equal(t, tc.p, p)
	}
}
