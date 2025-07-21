package tries

import (
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./tries/ -run ^TestDeleteDuplicateFolder$
func TestDeleteDuplicateFolder(t *testing.T) {
	for _, tc := range []struct {
		paths   [][]string
		cleaned [][]string
	}{
		{paths: [][]string{{"a"}, {"c"}, {"d"}, {"a", "b"}, {"c", "b"}, {"d", "a"}}, cleaned: [][]string{{"d"}, {"d", "a"}}},
		{paths: [][]string{{"a"}, {"c"}, {"a", "b"}, {"c", "b"}, {"a", "b", "x"}, {"a", "b", "x", "y"}, {"w"}, {"w", "y"}}, cleaned: [][]string{{"c"}, {"c", "b"}, {"a"}, {"a", "b"}}},
		{paths: [][]string{{"a", "b"}, {"c", "d"}, {"c"}, {"a"}}, cleaned: [][]string{{"c"}, {"c", "d"}, {"a"}, {"a", "b"}}},
	} {
		cleaned := deleteDuplicateFolder(tc.paths)
		sort.Slice(tc.cleaned, func(i, j int) bool {
			sort.Strings(tc.cleaned[i])
			sort.Strings(tc.cleaned[j])
			return strings.Join(tc.cleaned[i], "") < strings.Join(tc.cleaned[j], "")
		})
		sort.Slice(cleaned, func(i, j int) bool {
			sort.Strings(cleaned[i])
			sort.Strings(cleaned[j])
			return strings.Join(cleaned[i], "") < strings.Join(cleaned[j], "")
		})
		assert.Equal(t, tc.cleaned, cleaned)
	}
}
