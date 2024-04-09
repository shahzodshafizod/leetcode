package queues

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./queues/ -run ^TestTimeRequiredToBuy$
func TestTimeRequiredToBuy(t *testing.T) {
	for _, tc := range []struct {
		tickets []int
		k       int
		time    int
	}{
		{tickets: []int{2, 3, 2}, k: 2, time: 6},
		{tickets: []int{5, 1, 1, 1}, k: 0, time: 8},
		{tickets: []int{84, 49, 5, 24, 70, 77, 87, 8}, k: 3, time: 154},
	} {
		time := timeRequiredToBuy(tc.tickets, tc.k)
		assert.Equal(t, tc.time, time)
	}
}
