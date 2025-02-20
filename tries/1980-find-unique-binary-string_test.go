package tries

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./tries/ -run ^TestFindDifferentBinaryString$
func TestFindDifferentBinaryString(t *testing.T) {
	for _, tc := range []struct {
		nums []string
		num  string
	}{
		{nums: []string{"01", "10"}, num: "11"},
		{nums: []string{"00", "01"}, num: "10"},
		{nums: []string{"111", "011", "001"}, num: "000"},
	} {
		num := findDifferentBinaryString(tc.nums)
		assert.Equal(t, tc.num, num)
	}
}
