package backtracking

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./backtracking/ -run ^TestAddOperators$
func TestAddOperators(t *testing.T) {
	for _, tc := range []struct {
		num      string
		target   int
		expected []string
	}{
		{"123", 6, []string{"1*2*3", "1+2+3"}},
		{"232", 8, []string{"2*3+2", "2+3*2"}},
		{"3456237490", 9191, []string{}},
		{"105", 5, []string{"1*0+5", "10-5"}},
		{"00", 0, []string{"0*0", "0+0", "0-0"}},
	} {
		output := addOperators(tc.num, tc.target)
		sort.Strings(tc.expected)
		sort.Strings(output)
		assert.Equal(t, tc.expected, output)
	}
}
