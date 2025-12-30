package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./strings/ -run ^TestToLowerCase$
func TestToLowerCase(t *testing.T) {
	for _, tc := range []struct {
		s      string
		output string
	}{
		{"Hello", "hello"},
		{"here", "here"},
		{"LOVELY", "lovely"},
		{"", ""},
		{"a", "a"},
		{"A", "a"},
		{"123!@# ABCxyz", "123!@# abcxyz"},
	} {
		output := toLowerCase(tc.s)
		assert.Equal(t, tc.output, output)
	}
}
