package stacks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./stacks/ -run ^TestCountOfAtoms$
func TestCountOfAtoms(t *testing.T) {
	for _, tc := range []struct {
		formula   string
		atomcount string
	}{
		{formula: "H2O", atomcount: "H2O"},
		{formula: "Mg(OH)2", atomcount: "H2MgO2"},
		{formula: "K4(ON(SO3)2)2", atomcount: "K4N2O14S4"},
	} {
		atomcount := countOfAtoms(tc.formula)
		assert.Equal(t, tc.atomcount, atomcount)
	}
}
