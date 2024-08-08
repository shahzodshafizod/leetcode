package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./strings/ -run ^TestCountSeniors$
func TestCountSeniors(t *testing.T) {
	for _, tc := range []struct {
		details []string
		count   int
	}{
		{details: []string{"7868190130M7522", "5303914400F9211", "9273338290F4010"}, count: 2},
		{details: []string{"1313579440F2036", "2921522980M5644"}, count: 0},
		{details: []string{"5612624052M0130", "5378802576M6424", "5447619845F0171", "2941701174O9078"}, count: 2},
	} {
		count := countSeniors(tc.details)
		assert.Equal(t, tc.count, count)
	}
}
