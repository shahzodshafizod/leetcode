package stacks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./stacks/ -run ^TestMakeGood$
func TestMakeGood(t *testing.T) {
	for _, tc := range []struct {
		s    string
		good string
	}{
		{s: "leEeetcode", good: "leetcode"},
		{s: "abBAcC", good: ""},
		{s: "ABbaCc", good: ""},
		{s: "s", good: "s"},
	} {
		good := makeGood(tc.s)
		assert.Equal(t, tc.good, good)
	}
}
