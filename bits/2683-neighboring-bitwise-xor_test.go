package bits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./bits/ -run ^TestDoesValidArrayExist$
func TestDoesValidArrayExist(t *testing.T) {
	for _, tc := range []struct {
		derived []int
		exists  bool
	}{
		{derived: []int{1, 1, 0}, exists: true},
		{derived: []int{1, 1, 0}, exists: true},
		{derived: []int{1, 0}, exists: false},
	} {
		exists := doesValidArrayExist(tc.derived)
		assert.Equal(t, tc.exists, exists)
	}
}
