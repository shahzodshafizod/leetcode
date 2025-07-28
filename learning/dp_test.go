package learning

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./dp/ -run ^TestDPFib$
func TestDPFib(t *testing.T) {
	dps := []DP{NewMemoization(), NewTabulation()}
	assert := assert.New(t)

	for _, tc := range []struct {
		n int
		f int
	}{
		{n: 6, f: 8},
		{n: 7, f: 13},
		{n: 8, f: 21},
		{n: 50, f: 12586269025},
	} {
		for _, dp := range dps {
			f := dp.Fib(tc.n)
			assert.Equal(tc.f, f)
		}
	}
}

// go test -v -count=1 ./dp/ -run ^TestDPGridTraveler$
func TestDPGridTraveler(t *testing.T) {
	dps := []DP{NewMemoization(), NewTabulation()}
	assert := assert.New(t)

	for _, tc := range []struct {
		m    int
		n    int
		ways int
	}{
		{m: 1, n: 1, ways: 1},
		{m: 2, n: 3, ways: 3},
		{m: 3, n: 2, ways: 3},
		{m: 3, n: 3, ways: 6},
		{m: 18, n: 18, ways: 2333606220},
	} {
		for _, dp := range dps {
			ways := dp.GridTravaler(tc.m, tc.n)
			assert.Equal(tc.ways, ways)
		}
	}
}

// go test -v -count=1 ./dp/ -run ^TestDPCanSum$
func TestDPCanSum(t *testing.T) {
	dps := []DP{NewMemoization(), NewTabulation()}
	assert := assert.New(t)

	for _, tc := range []struct {
		targetSum int
		numbers   []int
		can       bool
	}{
		{targetSum: 7, numbers: []int{2, 3}, can: true},
		{targetSum: 7, numbers: []int{5, 3, 4, 7}, can: true},
		{targetSum: 7, numbers: []int{2, 4}, can: false},
		{targetSum: 8, numbers: []int{2, 3, 5}, can: true},
		{targetSum: 300, numbers: []int{7, 14}, can: false},
		{targetSum: 7, numbers: []int{3, 6, 7, 2}, can: true},
		{targetSum: 1, numbers: []int{2}, can: false},
		{targetSum: 6, numbers: []int{2, 3}, can: true},
	} {
		for _, dp := range dps {
			can := dp.CanSum(tc.targetSum, tc.numbers)
			assert.Equal(tc.can, can)
		}
	}
}

// go test -v -count=1 ./dp/ -run ^TestDPHowSum$
func TestDPHowSum(t *testing.T) {
	dps := []DP{NewMemoization(), NewTabulation()}
	assert := assert.New(t)

	for _, tc := range []struct {
		targetSum   int
		numbers     []int
		combination []int
	}{
		{targetSum: 7, numbers: []int{2, 3}, combination: []int{2, 2, 3}},
		{targetSum: 7, numbers: []int{5, 3, 4, 7}, combination: []int{7}},
		{targetSum: 7, numbers: []int{2, 4}, combination: nil},
		{targetSum: 8, numbers: []int{2, 3, 5}, combination: []int{3, 5}},
		{targetSum: 300, numbers: []int{7, 14}, combination: nil},
		// {targetSum: 7, numbers: []int{3, 6, 7, 2}, combination: []int{3, 2, 2}}, // or {7}
		{targetSum: 1, numbers: []int{2}, combination: nil},
		{targetSum: 6, numbers: []int{2, 3}, combination: []int{3, 3}},
	} {
		for _, dp := range dps {
			combination := dp.HowSum(tc.targetSum, tc.numbers)
			assert.Equal(tc.combination, combination)
		}
	}
}

// go test -v -count=1 ./dp/ -run ^TestDPBestSum$
func TestDPBestSum(t *testing.T) {
	dps := []DP{NewMemoization(), NewTabulation()}
	assert := assert.New(t)

	for _, tc := range []struct {
		targetSum   int
		numbers     []int
		combination []int
	}{
		{targetSum: 7, numbers: []int{2, 3}, combination: []int{3, 2, 2}},
		{targetSum: 7, numbers: []int{5, 3, 4, 7}, combination: []int{7}},
		{targetSum: 7, numbers: []int{2, 4}, combination: nil},
		{targetSum: 8, numbers: []int{2, 3, 5}, combination: []int{5, 3}},
		{targetSum: 300, numbers: []int{7, 14}, combination: nil},
		{targetSum: 7, numbers: []int{3, 6, 7, 2}, combination: []int{7}},
		{targetSum: 1, numbers: []int{2}, combination: nil},
		{targetSum: 6, numbers: []int{2, 3}, combination: []int{3, 3}},
		{targetSum: 7, numbers: []int{5, 3, 4, 7}, combination: []int{7}},
		{targetSum: 8, numbers: []int{1, 4, 5}, combination: []int{4, 4}},
		{targetSum: 100, numbers: []int{1, 5, 25, 2}, combination: []int{25, 25, 25, 25}},
	} {
		for _, dp := range dps {
			combination := dp.BestSum(tc.targetSum, tc.numbers)
			assert.Equal(tc.combination, combination)
		}
	}
}

// go test -v -count=1 ./dp/ -run ^TestDPCanConstruct$
func TestDPCanConstruct(t *testing.T) {
	dps := []DP{NewMemoization(), NewTabulation()}
	assert := assert.New(t)

	for _, tc := range []struct {
		target   string
		wordBank []string
		can      bool
	}{
		{target: "abcdef", wordBank: []string{"ab", "abc", "cd", "def", "abcd"}, can: true},
		{target: "skateboard", wordBank: []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}, can: false},
		{target: "enterapotentpot", wordBank: []string{"a", "p", "ent", "enter", "ot", "o", "t"}, can: true},
		{
			target:   "eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef",
			wordBank: []string{"e", "ee", "eee", "eeee", "eeeee", "eeeeee"},
			can:      false,
		},
	} {
		for _, dp := range dps {
			can := dp.CanConstruct(tc.target, tc.wordBank)
			assert.Equal(tc.can, can)
		}
	}
}

// go test -v -count=1 ./dp/ -run ^TestDPCountConstruct$
func TestDPCountConstruct(t *testing.T) {
	dps := []DP{NewMemoization(), NewTabulation()}
	assert := assert.New(t)

	for _, tc := range []struct {
		target   string
		wordBank []string
		count    int
	}{
		{target: "purple", wordBank: []string{"purp", "p", "ur", "le", "purpl"}, count: 2},
		{target: "abcdef", wordBank: []string{"ab", "abc", "cd", "def", "abcd"}, count: 1},
		{target: "skateboard", wordBank: []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}, count: 0},
		{target: "enterapotentpot", wordBank: []string{"a", "p", "ent", "enter", "ot", "o", "t"}, count: 4},
		{
			target:   "eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef",
			wordBank: []string{"e", "ee", "eee", "eeee", "eeeee", "eeeeee"},
			count:    0,
		},
	} {
		for _, dp := range dps {
			count := dp.CountConstruct(tc.target, tc.wordBank)
			assert.Equal(tc.count, count)
		}
	}
}

// go test -v -count=1 ./dp/ -run ^TestDPAllConstruct$
func TestDPAllConstruct(t *testing.T) {
	dps := []DP{NewMemoization(), NewTabulation()}
	assert := assert.New(t)

	for _, tc := range []struct {
		target   string
		wordBank []string
		allWays  [][]string
	}{
		{
			target:   "purple",
			wordBank: []string{"purp", "p", "ur", "le", "purpl"},
			allWays:  [][]string{{"purp", "le"}, {"p", "ur", "p", "le"}},
		},
		// {
		// 	target:   "abcdef",
		// 	wordBank: []string{"ab", "abc", "cd", "def", "abcd", "ef", "c"},
		// 	allWays:  [][]string{{"ab", "cd", "ef"}, {"ab", "c", "def"}, {"abc", "def"}, {"abcd", "ef"}},
		// }, // or {{"abc", "def"}, {"ab", "c", "def"}, {"abcd", "ef"}, {"ab", "cd", "ef"}}
		{target: "skateboard", wordBank: []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}, allWays: [][]string{}},
		// {
		// 	target:   "enterapotentpot",
		// 	wordBank: []string{"a", "p", "ent", "enter", "ot", "o", "t"},
		// 	allWays: [][]string{
		// 		{"enter", "a", "p", "ot", "ent", "p", "ot"},
		// 		{"enter", "a", "p", "ot", "ent", "p", "o", "t"},
		// 		{"enter", "a", "p", "o", "t", "ent", "p", "ot"},
		// 		{"enter", "a", "p", "o", "t", "ent", "p", "o", "t"},
		// 	},
		// 	// or
		// 	// allWays: [][]string{
		// 	// 	{"enter", "a", "p", "ot", "ent", "p", "ot"},
		// 	// 	{"enter", "a", "p", "o", "t", "ent", "p", "ot"},
		// 	// 	{"enter", "a", "p", "ot", "ent", "p", "o", "t"},
		// 	// 	{"enter", "a", "p", "o", "t", "ent", "p", "o", "t"},
		// 	// },
		// },
		{
			target:   "eeeeeef",
			wordBank: []string{"e", "ee", "eee", "eeee", "eeeee", "eeeeee"},
			allWays:  [][]string{},
		},
	} {
		for _, dp := range dps {
			allWays := dp.AllConstruct(tc.target, tc.wordBank)
			assert.Equal(tc.allWays, allWays)
		}
	}
}
