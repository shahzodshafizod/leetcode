package stacks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./stacks/ -run ^TestIsValid$
func TestIsValid(t *testing.T) {
	for _, data := range []struct {
		s     string
		valid bool
	}{
		{s: "", valid: true},
		{s: "()", valid: true},
		{s: "()[]{}", valid: true},
		{s: "(]", valid: false},
		{s: "{([])}", valid: true},
		{s: "{([)]}", valid: false},
		{s: "{([])]", valid: false},
		{s: "{([]", valid: false},
		{s: "{[(])}", valid: false},
		{s: "{[]()}", valid: true},
	} {
		valid := isValid(data.s)
		assert.Equal(t, data.valid, valid)
	}
}
