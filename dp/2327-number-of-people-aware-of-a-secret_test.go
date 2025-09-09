package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestPeopleAwareOfSecret$
func TestPeopleAwareOfSecret(t *testing.T) {
	for _, tc := range []struct {
		n      int
		delay  int
		forget int
		people int
	}{
		{n: 6, delay: 2, forget: 4, people: 5},
		{n: 4, delay: 1, forget: 3, people: 6},
	} {
		people := peopleAwareOfSecret(tc.n, tc.delay, tc.forget)
		assert.Equal(t, tc.people, people)
	}
}
