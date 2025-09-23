package design

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./design/ -run ^TestFoodRatings$
func TestFoodRatings(t *testing.T) {
	for _, tc := range []struct {
		commands []string
		values   []any
		output   []any
	}{
		{
			[]string{"FoodRatings", "highestRated", "highestRated", "changeRating", "highestRated", "changeRating", "highestRated"},
			[]any{
				[][]any{{"kimchi", "miso", "sushi", "moussaka", "ramen", "bulgogi"}, {"korean", "japanese", "japanese", "greek", "japanese", "korean"}, {9, 12, 8, 15, 14, 7}},
				[]any{"korean"},
				[]any{"japanese"},
				[]any{"sushi", 16},
				[]any{"japanese"},
				[]any{"ramen", 16},
				[]any{"japanese"},
			},
			[]any{nil, "kimchi", "ramen", nil, "sushi", nil, "ramen"},
		},
	} {
		var foodRatings FoodRatings

		var output any
		for idx, command := range tc.commands {
			output = nil

			switch command {
			case "FoodRatings":
				values, ok := tc.values[idx].([][]any)
				_ = ok

				foods := make([]string, 0, len(values[0]))
				for _, food := range values[0] {
					foods = append(foods, food.(string))
				}

				cuisines := make([]string, 0, len(values[1]))
				for _, cuisine := range values[1] {
					cuisines = append(cuisines, cuisine.(string))
				}

				ratings := make([]int, 0, len(values[2]))
				for _, rating := range values[2] {
					ratings = append(ratings, rating.(int))
				}

				foodRatings = NewFoodRatings(foods, cuisines, ratings)
			case "changeRating":
				values, ok := tc.values[idx].([]any)
				_ = ok
				food, ok := values[0].(string)
				_ = ok
				rating, ok := values[1].(int)
				_ = ok

				foodRatings.ChangeRating(food, rating)
			case "highestRated":
				values, ok := tc.values[idx].([]any)
				_ = ok
				cuisine, ok := values[0].(string)
				_ = ok
				output = foodRatings.HighestRated(cuisine)
			default:
			}

			assert.Equal(t, tc.output[idx], output)
		}
	}
}
