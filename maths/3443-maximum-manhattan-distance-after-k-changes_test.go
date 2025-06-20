package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./maths/ -run ^TestMaxDistance$
func TestMaxDistance(t *testing.T) {
	for _, tc := range []struct {
		s        string
		k        int
		distance int
	}{
		{s: "NWSE", k: 1, distance: 3},
		{s: "NSWWEW", k: 3, distance: 6},
		{s: "EWWE", k: 0, distance: 1},
		{s: "NSES", k: 1, distance: 4},
		{s: "SN", k: 0, distance: 1},
	} {
		distance := maxDistance(tc.s, tc.k)
		assert.Equal(t, tc.distance, distance)
	}
}
