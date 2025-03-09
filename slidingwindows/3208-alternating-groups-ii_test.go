package slidingwindows

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./slidingwindows/ -run ^TestNumberOfAlternatingGroups$
func TestNumberOfAlternatingGroups(t *testing.T) {
	for _, tc := range []struct {
		colors []int
		k      int
		groups int
	}{
		{colors: []int{0, 1, 0, 1, 0}, k: 3, groups: 3},
		{colors: []int{0, 1, 0, 0, 1, 0, 1}, k: 6, groups: 2},
		{colors: []int{1, 1, 0, 1}, k: 4, groups: 0},
	} {
		groups := numberOfAlternatingGroups(tc.colors, tc.k)
		assert.Equal(t, tc.groups, groups)
	}
}
