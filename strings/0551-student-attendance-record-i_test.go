package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./strings/ -run ^TestCheckRecord$
func TestCheckRecord(t *testing.T) {
	for _, tc := range []struct {
		s        string
		eligible bool
	}{
		{s: "PPALLP", eligible: true},
		{s: "PPALLL", eligible: false},
	} {
		eligible := checkRecord(tc.s)
		assert.Equal(t, tc.eligible, eligible)
	}
}
