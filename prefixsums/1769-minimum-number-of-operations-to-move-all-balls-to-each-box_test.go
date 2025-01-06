package prefixsums

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./prefixsums/ -run ^TestMinOperations$
func TestMinOperations(t *testing.T) {
	for _, tc := range []struct {
		boxes  string
		answer []int
	}{
		{boxes: "110", answer: []int{1, 1, 3}},
		{boxes: "001011", answer: []int{11, 8, 5, 4, 3, 4}},
	} {
		answer := minOperations(tc.boxes)
		assert.Equal(t, tc.answer, answer)
	}
}
