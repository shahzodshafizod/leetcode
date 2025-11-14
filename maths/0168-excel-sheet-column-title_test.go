package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./maths/ -run ^TestConvertToTitle$
func TestConvertToTitle(t *testing.T) {
	for _, tc := range []struct {
		columnNumber int
		columnTitle  string
	}{
		{columnNumber: 1, columnTitle: "A"},
		{columnNumber: 28, columnTitle: "AB"},
		{columnNumber: 701, columnTitle: "ZY"},
	} {
		columnTitle := convertToTitle(tc.columnNumber)
		assert.Equal(t, tc.columnTitle, columnTitle)
	}
}
