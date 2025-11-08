package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./maths/ -run ^TestTotalMoney$
func TestTotalMoney(t *testing.T) {
	for _, tc := range []struct {
		n     int
		total int
	}{
		{n: 4, total: 10},
		{n: 10, total: 37},
		{n: 20, total: 96},
	} {
		total := totalMoney(tc.n)
		assert.Equal(t, tc.total, total)
	}
}
