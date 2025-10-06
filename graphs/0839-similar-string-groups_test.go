package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./graphs/ -run ^TestNumSimilarGroups$
func TestNumSimilarGroups(t *testing.T) {
	for _, tc := range []struct {
		strs   []string
		groups int
	}{
		{strs: []string{"tars", "rats", "arts", "star"}, groups: 2},
		{strs: []string{"omv", "ovm"}, groups: 1},
		{strs: []string{"kccomwcgcs", "socgcmcwkc", "sgckwcmcoc", "coswcmcgkc", "cowkccmsgc", "cosgmccwkc", "sgmkwcccoc", "coswmccgkc", "kowcccmsgc", "kgcomwcccs"}, groups: 5},
	} {
		groups := numSimilarGroups(tc.strs)
		assert.Equal(t, tc.groups, groups)
	}
}
