package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestGetRow$
func TestGetRow(t *testing.T) {
	for _, tc := range []struct {
		rowIndex int
		row      []int
	}{
		{rowIndex: 3, row: []int{1, 3, 3, 1}},
		{rowIndex: 0, row: []int{1}},
		{rowIndex: 1, row: []int{1, 1}},
	} {
		row := getRow(tc.rowIndex)
		assert.Equal(t, tc.row, row)
	}
}
