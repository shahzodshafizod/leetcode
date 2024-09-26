package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestMinStickers$
func TestMinStickers(t *testing.T) {
	for _, tc := range []struct {
		stickers []string
		target   string
		count    int
	}{
		{stickers: []string{"with", "example", "science"}, target: "thehat", count: 3},
		{stickers: []string{"notice", "possible"}, target: "basicbasic", count: -1},
		{stickers: []string{"charge", "mind", "bottom"}, target: "centorder", count: 4},
		{stickers: []string{"fly", "me", "charge", "mind", "bottom"}, target: "centorder", count: 4},
		{stickers: []string{"these", "guess", "about", "garden", "him"}, target: "atomher", count: 3},
	} {
		count := minStickers(tc.stickers, tc.target)
		assert.Equal(t, tc.count, count)
	}
}
