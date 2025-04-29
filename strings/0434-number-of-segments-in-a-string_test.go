package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./strings/ -run ^TestCountSegments$
func TestCountSegments(t *testing.T) {
	for _, tc := range []struct {
		s     string
		count int
	}{
		{s: "Hello, my name is John", count: 5},
		{s: "Hello", count: 1},
		{s: ", , , , a, eaefa", count: 6},
		{s: " ", count: 0},
		{s: ", , , , a, e a e f a", count: 10},
		{s: "", count: 0},
	} {
		count := countSegments(tc.s)
		assert.Equal(t, tc.count, count)
	}
}
