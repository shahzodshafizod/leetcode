package design

import (
	"container/heap"

	"github.com/shahzodshafizod/leetcode/pkg"
)

// https://leetcode.com/problems/design-a-food-rating-system/

type foodItem struct {
	rating  int
	food    string
	cuisine string
}

type FoodRatings struct {
	cuisines map[string]*pkg.Heap[*foodItem]
	foods    map[string]*foodItem
}

func NewFoodRatings(foods []string, cuisines []string, ratings []int) FoodRatings {
	f := FoodRatings{
		cuisines: make(map[string]*pkg.Heap[*foodItem]),
		foods:    make(map[string]*foodItem),
	}
	for idx := range foods {
		if f.cuisines[cuisines[idx]] == nil {
			f.cuisines[cuisines[idx]] = pkg.NewHeap(
				make([]*foodItem, 0),
				func(x, y *foodItem) bool {
					return x.rating > y.rating || x.rating == y.rating && x.food < y.food
				},
			)
		}

		item := &foodItem{rating: ratings[idx], food: foods[idx], cuisine: cuisines[idx]}
		heap.Push(f.cuisines[cuisines[idx]], item)
		f.foods[foods[idx]] = item
	}

	return f
}

func (f *FoodRatings) ChangeRating(food string, newRating int) {
	cuisine := f.foods[food].cuisine
	item := &foodItem{
		rating:  newRating,
		food:    food,
		cuisine: cuisine,
	}
	f.foods[food] = item
	heap.Push(f.cuisines[cuisine], item)
}

func (f *FoodRatings) HighestRated(cuisine string) string {
	for f.cuisines[cuisine].Peak().rating != f.foods[f.cuisines[cuisine].Peak().food].rating {
		heap.Pop(f.cuisines[cuisine])
	}

	return f.cuisines[cuisine].Peak().food
}

// /**
//  * Your FoodRatings object will be instantiated and called as such:
//  * obj := Constructor(foods, cuisines, ratings);
//  * obj.ChangeRating(food,newRating);
//  * param_2 := obj.HighestRated(cuisine);
//  */
