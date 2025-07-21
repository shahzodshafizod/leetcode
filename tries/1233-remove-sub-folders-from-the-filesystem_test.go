package tries

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./tries/ -run ^TestRemoveSubfolders$
func TestRemoveSubfolders(t *testing.T) {
	for _, tc := range []struct {
		folder []string
		roots  []string
	}{
		{folder: []string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"}, roots: []string{"/a", "/c/d", "/c/f"}},
		{folder: []string{"/a", "/a/b/c", "/a/b/d"}, roots: []string{"/a"}},
		{folder: []string{"/a/b/c", "/a/b/ca", "/a/b/d"}, roots: []string{"/a/b/c", "/a/b/ca", "/a/b/d"}},
	} {
		roots := removeSubfolders(tc.folder)
		sort.Strings(tc.roots)
		sort.Strings(roots)
		assert.Equal(t, tc.roots, roots)
	}
}
