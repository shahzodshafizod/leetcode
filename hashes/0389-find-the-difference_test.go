package hashes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./hashes/ -run ^TestFindTheDifference$
func TestFindTheDifference(t *testing.T) {
	for _, tc := range []struct {
		s     string
		t     string
		extra byte
	}{
		{s: "abcd", t: "abcde", extra: 'e'},
		{s: "", t: "y", extra: 'y'},
		{s: "a", t: "aa", extra: 'a'},
	} {
		extra := findTheDifference(tc.s, tc.t)
		assert.Equal(t, tc.extra, extra)
	}
}
