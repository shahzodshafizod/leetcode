package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestNumDistinct$
func TestNumDistinct(t *testing.T) {
	for _, tc := range []struct {
		s     string
		t     string
		count int
	}{
		{s: "rabbbit", t: "rabbit", count: 3},
		{s: "babgbag", t: "bag", count: 5},
		{s: "a", t: "", count: 1},
		{s: "", t: "a", count: 0},
	} {
		count := numDistinct(tc.s, tc.t)
		assert.Equal(t, tc.count, count)
	}
}
