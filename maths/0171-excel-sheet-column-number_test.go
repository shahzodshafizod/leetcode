package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./maths/ -run ^TestTitleToNumber$
func TestTitleToNumber(t *testing.T) {
	for _, tc := range []struct {
		columnTitle  string
		columnNumber int
	}{
		{columnTitle: "A", columnNumber: 1},
		{columnTitle: "AB", columnNumber: 28},
		{columnTitle: "ZY", columnNumber: 701},
	} {
		columnNumber := titleToNumber(tc.columnTitle)
		assert.Equal(t, tc.columnNumber, columnNumber)
	}
}
