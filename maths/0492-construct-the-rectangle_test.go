package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./maths/ -run ^TestConstructRectangle$
func TestConstructRectangle(t *testing.T) {
	for _, tc := range []struct {
		area   int
		output []int
	}{
		{area: 4, output: []int{2, 2}},
		{area: 37, output: []int{37, 1}},
		{area: 122122, output: []int{427, 286}},
		{area: 1, output: []int{1, 1}},
		{area: 10, output: []int{5, 2}},
	} {
		output := constructRectangle(tc.area)
		assert.Equal(t, tc.output, output)
	}
}
