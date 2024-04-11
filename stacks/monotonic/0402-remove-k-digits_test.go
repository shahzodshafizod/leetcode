package monotonic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./stacks/monotonic/ -run ^TestRemoveKdigits$
func TestRemoveKdigits(t *testing.T) {
	for _, tc := range []struct {
		num   string
		k     int
		small string
	}{
		{num: "1432219", k: 3, small: "1219"},
		{num: "1143229", k: 3, small: "1122"},
		{num: "10200", k: 1, small: "200"},
		{num: "10", k: 2, small: "0"},
		{num: "1234567890", k: 9, small: "0"},
		{num: "123456789", k: 9, small: "0"},
		{num: "100001", k: 5, small: "0"},
	} {
		small := removeKdigits(tc.num, tc.k)
		assert.Equal(t, tc.small, small)
	}
}
