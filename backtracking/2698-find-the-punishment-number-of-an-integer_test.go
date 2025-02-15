package backtracking

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./backtracking/ -run ^TestPunishmentNumber$
func TestPunishmentNumber(t *testing.T) {
	for _, tc := range []struct {
		n          int
		panishment int
	}{
		{n: 10, panishment: 182},
		{n: 37, panishment: 1478},
	} {
		panishment := punishmentNumber(tc.n)
		assert.Equal(t, tc.panishment, panishment)
	}
}
