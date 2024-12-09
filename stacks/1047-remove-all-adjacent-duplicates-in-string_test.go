package stacks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./stacks/ -run ^TestRemoveDuplicates$
func TestRemoveDuplicates(t *testing.T) {
	for _, tc := range []struct {
		s        string
		modified string
	}{
		{s: "abbaca", modified: "ca"},
		{s: "azxxzy", modified: "ay"},
		{s: "aaaaaa", modified: ""},
		{s: "aaaaaaaaa", modified: "a"},
	} {
		modified := removeDuplicates(tc.s)
		assert.Equal(t, tc.modified, modified)
	}
}
