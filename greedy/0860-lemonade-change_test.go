package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./greedy/ -run ^TestLemonadeChange$
func TestLemonadeChange(t *testing.T) {
	for _, tc := range []struct {
		bills     []int
		canChange bool
	}{
		{bills: []int{5, 5, 5, 10, 20}, canChange: true},
		{bills: []int{5, 5, 10, 10, 20}, canChange: false},
		{bills: []int{5, 5, 5, 5, 10, 20, 10}, canChange: true},
		{bills: []int{5, 5, 5, 20}, canChange: true},
		{bills: []int{5, 5, 10, 20, 5, 5, 5, 5, 5, 5, 5, 5, 5, 10, 5, 5, 20, 5, 20, 5}, canChange: true},
		{bills: []int{5, 5, 5, 5, 20, 20, 5, 5, 20, 5}, canChange: false},
		{bills: []int{5, 5, 10, 10, 5, 20, 5, 10, 5, 5}, canChange: true},
	} {
		canChange := lemonadeChange(tc.bills)
		assert.Equal(t, tc.canChange, canChange)
	}
}
