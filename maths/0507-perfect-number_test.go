package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./maths/ -run ^TestCheckPerfectNumber$
func TestCheckPerfectNumber(t *testing.T) {
	for _, tc := range []struct {
		num     int
		perfect bool
	}{
		{num: 28, perfect: true}, // 1 + 2 + 4 + 7 + 14 = 28
		{num: 7, perfect: false},
		{num: 6, perfect: true},    // 1 + 2 + 3 = 6
		{num: 496, perfect: true},  // Perfect number
		{num: 8128, perfect: true}, // Perfect number
		{num: 1, perfect: false},
		{num: 2, perfect: false},
	} {
		perfect := checkPerfectNumber(tc.num)
		assert.Equal(t, tc.perfect, perfect)
	}
}
