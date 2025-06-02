package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./greedy/ -run ^TestCandy$
func TestCandy(t *testing.T) {
	for _, tc := range []struct {
		ratings []int
		candies int
	}{
		{ratings: []int{1, 0, 2}, candies: 5},
		{ratings: []int{1, 2, 2}, candies: 4},
		{ratings: []int{60, 80, 100, 100, 100, 100, 100}, candies: 10},
		{ratings: []int{100, 80, 70, 60, 70, 80, 90, 100, 90, 80, 70, 60, 60}, candies: 35},
		{ratings: []int{6, 7, 6, 5, 4, 3, 2, 1, 0, 0, 0, 1, 0}, candies: 42},
		// {ratings: []int{20000, 20000, 16001, 16001, 16002, 16002, 16003, 16003, 16004, 16004, 16005, 16005, 16006, 16006, 16007, 16007, 16008, 16008, 16009, 16009, 16010, 16010, 16011, 16011, 16012, 16012, 16013, 16013, 16014, 16014, 16015, 16015, 16016, 16016, 16017, 16017, 16018, 16018, 16019, 16019, 16020, 16020, 16021, 16021, 16022, 16022, 16023, 16023, 16024, 16024}, candies: 74},
	} {
		candies := candy(tc.ratings)
		assert.Equal(t, tc.candies, candies)
	}
}
