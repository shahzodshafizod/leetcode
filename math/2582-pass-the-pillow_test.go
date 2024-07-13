package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./math/ -run ^TestPassThePillow$
func TestPassThePillow(t *testing.T) {
	for _, tc := range []struct {
		n      int
		time   int
		person int
	}{
		{n: 4, time: 5, person: 2},
		{n: 3, time: 2, person: 3},
	} {
		person := passThePillow(tc.n, tc.time)
		assert.Equal(t, tc.person, person)
	}
}
