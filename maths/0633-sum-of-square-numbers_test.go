package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./maths/ -run ^TestJudgeSquareSum$
func TestJudgeSquareSum(t *testing.T) {
	for _, tc := range []struct {
		c     int
		exist bool
	}{
		{c: 5, exist: true},
		{c: 3, exist: false},
		{c: 2, exist: true},
	} {
		exist := judgeSquareSum(tc.c)
		assert.Equal(t, tc.exist, exist)
	}
}
