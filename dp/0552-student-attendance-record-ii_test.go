package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestCheckRecord$
func TestCheckRecord(t *testing.T) {
	for _, tc := range []struct {
		n       int
		records int
	}{
		{n: 2, records: 8},
		{n: 1, records: 3},
		// {n: 10101, records: 183236316},
		// {n: 59, records: 667741397},
		{n: 3, records: 19},
	} {
		records := checkRecord(tc.n)
		assert.Equal(t, tc.records, records)
	}
}
