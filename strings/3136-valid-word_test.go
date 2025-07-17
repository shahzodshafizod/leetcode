package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./strings/ -run ^TestIsValid$
func TestIsValid(t *testing.T) {
	for _, tc := range []struct {
		word string
		valid bool
	}{
		{word: "234Adas", valid: true},
		{word: "b3", valid: false},
		{word: "a3$e", valid: false},
	} {
		valid := isValid(tc.word)
		assert.Equal(t, tc.valid, valid)
	}
}
