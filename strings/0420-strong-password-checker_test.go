package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./strings/ -run ^TestStrongPasswordChecker$
func TestStrongPasswordChecker(t *testing.T) {
	for _, tc := range []struct {
		password string
		output   int
	}{
		{"a", 5},        // Too short, missing types
		{"aA1", 3},      // Too short
		{"1337C0d3", 0}, // Already strong
	} {
		output := strongPasswordChecker(tc.password)
		assert.Equal(t, tc.output, output)
	}
}
