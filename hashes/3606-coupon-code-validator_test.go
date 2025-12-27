package hashes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./hashes/ -run ^TestValidateCoupons$
func TestValidateCoupons(t *testing.T) {
	for _, tc := range []struct {
		code         []string
		businessLine []string
		isActive     []bool
		output       []string
	}{
		{
			code:         []string{"SAVE20", "", "PHARMA5", "SAVE@20"},
			businessLine: []string{"restaurant", "grocery", "pharmacy", "restaurant"},
			isActive:     []bool{true, true, true, true},
			output:       []string{"PHARMA5", "SAVE20"},
		},
		{
			code:         []string{"GROCERY15", "ELECTRONICS_50", "DISCOUNT10"},
			businessLine: []string{"grocery", "electronics", "invalid"},
			isActive:     []bool{false, true, true},
			output:       []string{"ELECTRONICS_50"},
		},
		{
			code:         []string{"E1", "E2", "G1", "P1", "R1"},
			businessLine: []string{"electronics", "electronics", "grocery", "pharmacy", "restaurant"},
			isActive:     []bool{true, true, true, true, true},
			output:       []string{"E1", "E2", "G1", "P1", "R1"},
		},
	} {
		output := validateCoupons(tc.code, tc.businessLine, tc.isActive)
		assert.Equal(t, tc.output, output)
	}
}
