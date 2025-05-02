package twopointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./twopointers/ -run ^TestPushDominoes$
func TestPushDominoes(t *testing.T) {
	for _, tc := range []struct {
		dominoes string
		final    string
	}{
		{dominoes: "RR.L", final: "RR.L"},
		{dominoes: ".L.R...LR..L..", final: "LL.RR.LLRRLL.."},
	} {
		final := pushDominoes(tc.dominoes)
		assert.Equal(t, tc.final, final)
	}
}
