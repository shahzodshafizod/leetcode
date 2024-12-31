package bits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./bits/ -run ^TestAddBinary$
func TestAddBinary(t *testing.T) {
	for _, tc := range []struct {
		a     string
		b     string
		bisum string
	}{
		{a: "11", b: "1", bisum: "100"},
		{a: "1010", b: "1011", bisum: "10101"},
	} {
		bisum := addBinary(tc.a, tc.b)
		assert.Equal(t, tc.bisum, bisum)
	}
}
