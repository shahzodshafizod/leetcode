package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./maths/ -run ^TestFlowerGame$
func TestFlowerGame(t *testing.T) {
	for _, tc := range []struct {
		n     int
		m     int
		count int64
	}{
		{n: 3, m: 2, count: 3},
		{n: 1, m: 1, count: 0},
	} {
		count := flowerGame(tc.n, tc.m)
		assert.Equal(t, tc.count, count)
	}
}
