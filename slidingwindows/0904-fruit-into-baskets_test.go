package slidingwindows

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./slidingwindows/ -run ^TestTotalFruit$
func TestTotalFruit(t *testing.T) {
	for _, tc := range []struct {
		fruits []int
		total  int
	}{
		{fruits: []int{1, 2, 1}, total: 3},
		{fruits: []int{0, 1, 2, 2}, total: 3},
		{fruits: []int{1, 2, 3, 2, 2}, total: 4},
	} {
		total := totalFruit(tc.fruits)
		assert.Equal(t, tc.total, total)
	}
}
