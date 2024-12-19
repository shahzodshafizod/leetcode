package template

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./template/ -run ^Test$
func Test(t *testing.T) {
	for _, tc := range []struct {
		output any
	}{
		// {},
	} {
		output := a()
		assert.Equal(t, tc.output, output)
	}
}
