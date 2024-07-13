package stacks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./stacks/ -run ^TestMaximumGain$
func TestMaximumGain(t *testing.T) {
	for _, tc := range []struct {
		s      string
		x      int
		y      int
		points int
	}{
		{s: "cdbcbbaaabab", x: 4, y: 5, points: 19},
		{s: "aabbaaxybbaabb", x: 5, y: 4, points: 20},
		{s: "abbaab", x: 1, y: 100, points: 201},
	} {
		points := maximumGain(tc.s, tc.x, tc.y)
		assert.Equal(t, tc.points, points)
	}
}
