package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./strings/ -run ^TestCapitalizeTitle$
func TestCapitalizeTitle(t *testing.T) {
	for _, tc := range []struct {
		title       string
		capitalized string
	}{
		{title: "capiTalIze tHe titLe", capitalized: "Capitalize The Title"},
		{title: "First leTTeR of EACH Word", capitalized: "First Letter of Each Word"},
		{title: "i lOve leetcode", capitalized: "i Love Leetcode"},
	} {
		capitalized := capitalizeTitle(tc.title)
		assert.Equal(t, tc.capitalized, capitalized)
	}
}
