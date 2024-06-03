package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestIsSubsequence$
func TestIsSubsequence(t *testing.T) {
	for _, tc := range []struct {
		s  string
		t  string
		is bool
	}{
		{s: "abc", t: "ahbgdc", is: true},
		{s: "axc", t: "ahbgdc", is: false},
	} {
		is := isSubsequence(tc.s, tc.t)
		assert.Equal(t, tc.is, is)
	}
}
