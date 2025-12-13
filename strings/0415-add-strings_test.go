package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./strings/ -run ^TestAddStrings$
func TestAddStrings(t *testing.T) {
	for _, tc := range []struct {
		num1   string
		num2   string
		output string
	}{
		{num1: "11", num2: "123", output: "134"},
		{num1: "456", num2: "77", output: "533"},
		{num1: "0", num2: "0", output: "0"},
		{num1: "1", num2: "9", output: "10"},
		{num1: "999", num2: "1", output: "1000"},
		{num1: "123456789", num2: "987654321", output: "1111111110"},
	} {
		output := addStrings(tc.num1, tc.num2)
		assert.Equal(t, tc.output, output)
	}
}
