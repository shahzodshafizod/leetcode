package bits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./bits/ -run ^TestToHex$
func TestToHex(t *testing.T) {
	for _, tc := range []struct {
		num int
		hex string
	}{
		{num: 26, hex: "1a"},
		{num: -1, hex: "ffffffff"},
		{num: 0, hex: "0"},
		{num: 1, hex: "1"},
		{num: 15, hex: "f"},
		{num: 16, hex: "10"},
		{num: 255, hex: "ff"},
	} {
		hex := toHex(tc.num)
		assert.Equal(t, tc.hex, hex)
	}
}
