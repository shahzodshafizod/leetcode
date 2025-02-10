package stacks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./stacks/ -run ^TestClearDigits$
func TestClearDigits(t *testing.T) {
	for _, tc := range []struct {
		s       string
		cleared string
	}{
		{s: "abc", cleared: "abc"},
		{s: "cb34", cleared: ""},
		{s: "a", cleared: "a"},
		{s: "xzuzr2ghilydk", cleared: "xzuzghilydk"},
		{s: "qm93xjkpaaovhqckjhg5j1o4rndtg3bobgeke", cleared: "xjkpaaovhqckjhrndtbobgeke"},
	} {
		cleared := clearDigits(tc.s)
		assert.Equal(t, tc.cleared, cleared)
	}
}
