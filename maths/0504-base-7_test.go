package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./maths/ -run ^TestConvertToBase7$
func TestConvertToBase7(t *testing.T) {
	for _, tc := range []struct {
		num   int
		base7 string
	}{
		{num: 100, base7: "202"},
		{num: -7, base7: "-10"},
		{num: 0, base7: "0"},
		{num: 7, base7: "10"},
		{num: -8, base7: "-11"},
		{num: 49, base7: "100"},
	} {
		base7 := convertToBase7(tc.num)
		assert.Equal(t, tc.base7, base7)
	}
}
