package learning

import (
	"strings"
)

/*
Top-Down approach follows the memorization technique.
It consists of two distinct events: recursion and caching.

The easiest way to remember them is that
Bottom-Up is iterative and Top-Down is recursive.
*/

// Source: Dynamic Programming - Learn to Solve Algorithmic Problems & Coding Challenges
// https://www.youtube.com/watch?v=oBt53YbR9Kk

type DP interface {
	/*
		Write a function `fib(n)` that takes in a number as an argument.
		The function should return the n-th number of the Fibonacci sequence.

		The 0th number of the sequence is 0. The 1st number of the sequence is 1.
		To generate the next number of the sequence, we sum the previous two.
	*/
	Fib(n int) int

	/*
		Say that you are a traveler on a 2D grid. You begin in the
		top-left corner and your goal is to travel to the bottom-right
		corner. You may only move down or right.

		In how many ways can you travel to the goal on a grid with
		dimensions m * n?

		Write a function `gridTraveler(m, n)` that calculates this.
	*/
	GridTravaler(m int, n int) int

	/*
		Write a function `canSum(targetSum, numbers)` that takes in a
		targetSum and an array of numbers as arguments.

		The function should return a boolean indicating whether or not it
		is possible to generate the targetSum using numbers from the array.

		You may use an element of the array as many times as needed.

		You may assume that all input numbers are non-negative.
	*/
	CanSum(targetSum int, numbers []int) bool

	/*
		Write a function `howSum(targetSum, numbers)` that takes in a
		targetSum and an array of numbers as arguments.

		The function should return an array containing any combination of
		elements that add up to exactly the targetSum. If there is no
		combination that adds up to the targetSum, then return null.

		If there are multiple combinations possible, you may return any
		single one.
	*/
	HowSum(targetSum int, numbers []int) []int

	/*
		Write a function `bestSum(targetSum, numbers)` that takes in a
		targetSum and an array of numbers as arguments.

		The function should return an array containing the shortest
		combination of numbers that add up to exactly the targetSum.

		If there is a tie for the shortest combination, you may return any
		one of the shortest.
	*/
	BestSum(targetSum int, numbers []int) []int

	/*
		Write a function `canConstruct(target, wordBank)` that accepts a
		target string and an array of strings.

		The function should return a boolean indicating whether or not the
		`target` can be constructed by concatenating elements of the
		`wordBank` array.

		You may reuse elements of `wordBank` as many times as needed.
	*/
	CanConstruct(target string, wordBank []string) bool

	/*
		Write a function `countConstruct(target, wordBank)` that accepts a
		target string and an array of strings.

		The function should return the number of ways that the `target` can
		be constructed by concatenating elements of the `wordBank` array.

		You may reuse elements of `wordBank` as many times as needed.
	*/
	CountConstruct(target string, wordBank []string) int

	/*
		Write a function `allConstruct(target, wordBank)` that accepts a
		target string and an array of strings.

		The function should return a 2D array containing `all of the ways`
		that the `target` can be constructed by concatenating elements of
		the `wordBank` array. Each element of the 2D array should represent
		one combination that construct the `target`.

		You may reuse of `wordBank` as many times as needed.
	*/
	AllConstruct(target string, wordBank []string) [][]string
}

/*
	Memoization Recipe from Alvin Zablan

1. Make it work.
  - visualize the problem as a tree
  - implement the tree using recursion
  - test it

2. Make it efficient
  - add a memo object
  - add a base case to return memo values
  - store return values into the memo
*/
type memoization struct{}

func NewMemoization() DP { return &memoization{} }

// time: O(n)
// space: O(n)
func (m *memoization) Fib(n int) int {
	var dp func(n int, memo []int) int

	dp = func(n int, memo []int) int {
		if n <= 2 {
			return 1
		}

		if memo[n-1] != 0 {
			return memo[n-1]
		}

		memo[n-1] = dp(n-1, memo) + dp(n-2, memo)

		return memo[n-1]
	}

	return dp(n, make([]int, n))
}

// time: O(m*n)
// space: O(m*n)
func (m *memoization) GridTravaler(mm int, n int) int {
	var dp func(mm int, n int, memo [][]int) int

	dp = func(mm int, n int, memo [][]int) int {
		if mm == 0 || n == 0 {
			return 0
		}

		if mm == 1 && n == 1 {
			return 1
		}

		if memo[mm-1][n-1] != 0 {
			return memo[mm-1][n-1]
		}

		memo[mm-1][n-1] = dp(mm-1, n, memo) + dp(mm, n-1, memo)

		return memo[mm-1][n-1]
	}

	memo := make([][]int, mm)
	for i := 0; i < mm; i++ {
		memo[i] = make([]int, n)
	}

	return dp(mm, n, memo)
}

// m = targetSum
// n = len(numbers)
// time: O(n * m)
// space: O(m)
func (m *memoization) CanSum(targetSum int, numbers []int) bool {
	var dp func(targetSum int, idx int, memo []*bool) bool

	dp = func(targetSum int, idx int, memo []*bool) bool {
		if idx < 0 || targetSum <= 0 {
			return targetSum == 0
		}

		if memo[targetSum] != nil {
			return *memo[targetSum]
		}
		// decision to include
		targetSum -= numbers[idx]

		can := dp(targetSum, idx, memo)
		if can {
			memo[targetSum] = &can

			return true
		}
		// decision NOT to include
		targetSum += numbers[idx]
		can = dp(targetSum, idx-1, memo)
		memo[targetSum] = &can

		return can
	}

	return dp(targetSum, len(numbers)-1, make([]*bool, targetSum+1))
}

// m = targetSum
// n = len(numbers)
// time: O(n * m^2) - for every unit of targetSum (m) iterate every numbers (n) and copy slice of size m
// space: O(m^2)
func (m *memoization) HowSum(targetSum int, numbers []int) []int {
	var dp func(targetSum int, idx int, memo []*[]int) []int

	dp = func(targetSum int, idx int, memo []*[]int) []int {
		if idx < 0 || targetSum < 0 {
			return nil
		}

		if targetSum == 0 {
			return []int{}
		}

		if memo[targetSum] != nil {
			return *memo[targetSum]
		}
		// decision to include
		targetSum -= numbers[idx]

		combination := dp(targetSum, idx, memo)
		if combination != nil {
			combination = append(combination, numbers[idx])
			memo[targetSum] = &combination

			return combination
		}
		// decision NOT to include
		targetSum += numbers[idx]
		combination = dp(targetSum, idx-1, memo)
		memo[targetSum] = &combination

		return combination
	}

	return dp(targetSum, len(numbers)-1, make([]*[]int, targetSum+1))
}

// m = targetSum
// n = len(numbers)
// time: O(n * m^2)
// space: O(m^2)
func (m *memoization) BestSum(targetSum int, numbers []int) []int {
	var dp func(targetSum int, memo []*[]int) []int

	dp = func(targetSum int, memo []*[]int) []int {
		if targetSum < 0 {
			return nil
		}

		if targetSum == 0 {
			return []int{}
		}

		if memo[targetSum] != nil {
			return *memo[targetSum]
		}

		var shortest []int

		for _, number := range numbers {
			remainder := targetSum - number

			combination := dp(remainder, memo)
			if combination != nil {
				combination = append(combination, number)
				if shortest == nil || len(combination) < len(shortest) {
					shortest = combination
				}
			}
		}

		memo[targetSum] = &shortest

		return shortest
	}

	return dp(targetSum, make([]*[]int, targetSum+1))
	// // BACKTRACKING
	// var shortest []int = nil
	// var dp func(targetSum int, idx int, combination []int)
	// var count = 0
	// dp = func(targetSum int, idx int, combination []int) {
	// 	count++
	// 	if targetSum < 0 || idx < 0 {
	// 		return
	// 	}
	// 	if targetSum == 0 {
	// 		if shortest == nil || len(combination) < len(shortest) {
	// 			shortest = append([]int{}, combination...)
	// 		}
	// 		return
	// 	}
	// 	// decision to include
	// 	targetSum -= numbers[idx]
	// 	dp(targetSum, idx, append(combination, numbers[idx]))
	// 	// decision NOT to include
	// 	targetSum += numbers[idx]
	// 	dp(targetSum, idx-1, combination)
	// }
	// dp(targetSum, len(numbers)-1, []int{})
	// return shortest
}

// m = len(target)
// n = len(wordBank)
// time: O(n * m^2)
// space: O(m^2) - saving every combination
func (m *memoization) CanConstruct(target string, wordBank []string) bool {
	// backtracking is not suitable, bc it never uses first word with the last word combined
	var dp func(target string, memo map[string]bool) bool

	dp = func(target string, memo map[string]bool) bool {
		if target == "" {
			return true
		}

		if can, exists := memo[target]; exists {
			return can
		}

		for _, word := range wordBank {
			if strings.Index(target, word) == 0 && dp(target[len(word):], memo) {
				return true
			}
		}

		memo[target] = false

		return false
	}

	return dp(target, make(map[string]bool))
}

// m = len(target)
// n = len(numbers)
// time: O(n * m^2)
// space: O(m^2) - saving every combination
func (m *memoization) CountConstruct(target string, wordBank []string) int {
	var dp func(target string, memo map[string]int) int

	dp = func(target string, memo map[string]int) int {
		if target == "" {
			return 1
		}

		if count, exists := memo[target]; exists {
			return count
		}

		count := 0

		for _, word := range wordBank {
			if strings.Index(target, word) == 0 {
				restOfTarget := target[len(word):]
				ways := dp(restOfTarget, memo)
				memo[restOfTarget] = ways
				count += ways
			}
		}

		memo[target] = count

		return count
	}

	return dp(target, make(map[string]int))
}

// m = len(target)
// n = len(numbers)
// time: O(n^m)
// space: O(m) - ?????
func (m *memoization) AllConstruct(target string, wordBank []string) [][]string {
	var dp func(target string, memo map[string][][]string) [][]string

	dp = func(target string, memo map[string][][]string) [][]string {
		if target == "" {
			return [][]string{{}}
		}

		if ways, exists := memo[target]; exists {
			return ways
		}

		allWays := make([][]string, 0)

		for _, word := range wordBank {
			if strings.Index(target, word) == 0 {
				restOfTarget := target[len(word):]

				ways := dp(restOfTarget, memo)
				for i := range ways {
					ways[i] = append([]string{word}, ways[i]...)
				}

				allWays = append(allWays, ways...)
			}
		}

		memo[target] = append([][]string{}, allWays...)

		return allWays
	}

	return dp(target, make(map[string][][]string))
}

/*
	Tabulation Recipe

- visualize the problem as a table
- size the table based on the inputs
- initialize the table with default values
- seed the trivial answer into the table
- iterate through the table
- fill further positions based on the current position
*/
type tabulation struct{}

func NewTabulation() DP { return &tabulation{} }

// time: O(n)
// space: O(1)
func (t *tabulation) Fib(n int) int {
	// // space: O(n)
	// var dp = make([]int, n+1)
	// dp[0] = 0
	// dp[1] = 1
	// for i := 2; i <= n; i++ {
	// 	dp[i] = dp[i-1] + dp[i-2]
	// }
	// return dp[n]
	prev := 0
	curr := 1

	for i := 2; i <= n; i++ {
		prev, curr = curr, prev+curr
	}

	return curr
}

// time: O(m*n)
// space: O(m)
func (t *tabulation) GridTravaler(m int, n int) int {
	// example: m:3, n:3
	// [1][1][1]
	// [1][2][3]
	// [1][3][6]
	array := make([]int, m)
	array[0] = 1

	for col := 0; col < n; col++ {
		for row := 1; row < m; row++ {
			array[row] += array[row-1]
		}
	}

	return array[m-1]
	// // space: O(m*n)
	// var grid = make([][]int, m)
	// for i := 0; i < m; i++ {
	// 	grid[i] = make([]int, n)
	// 	grid[i][0] = 1
	// }
	// for col := 0; col < n; col++ {
	// 	grid[0][col] = 1
	// }
	// for row := 1; row < m; row++ {
	// 	for col := 1; col < n; col++ {
	// 		grid[row][col] = grid[row-1][col] + grid[row][col-1]
	// 	}
	// }
	// return grid[m-1][n-1]
}

// m = targetSum
// n = len(numbers)
// time: O(m * n)
// space: O(m)
func (t *tabulation) CanSum(targetSum int, numbers []int) bool {
	// example: targetSum:7, numbers:[5,3,4]
	// idx: 0 1 2 3 4 5 6 7 (size: targetSum+1)
	// can: [T F F F F F F F]
	// can: [T] F F T T T F F (0+5->T, 0+3->T, 0+4->T)
	// can: T [F] [F] T T T F F skip, skip
	// can: T F F [T] T T T T (3+5->[out], 3+3->T, 3+4->T) and so on.
	dp := make([]bool, targetSum+1)
	dp[0] = true
	exit := false

	for i := 0; i <= targetSum && !exit; i++ {
		if !dp[i] {
			continue
		}

		exit = true

		for _, number := range numbers {
			if number+i <= targetSum {
				exit = false
				dp[number+i] = true
			}
		}
	}

	return dp[targetSum]
}

// m = targetSum
// n = len(numbers)
// time: O(n * m^2)
// space: O(m^2)
func (t *tabulation) HowSum(targetSum int, numbers []int) []int {
	dp := make([][]int, targetSum+1)
	dp[0] = []int{}
	exit := false

	for i := 0; i <= targetSum && !exit; i++ {
		if dp[i] == nil {
			continue
		}

		exit = true

		for _, number := range numbers {
			if number+i <= targetSum && dp[number+i] == nil {
				exit = false

				dp[number+i] = append(dp[i], number)
			}
		}
	}

	return dp[targetSum]
}

// m = targetSum
// n = len(numbers)
// time: O(n * m^2)
// space: O(m^2)
func (t *tabulation) BestSum(targetSum int, numbers []int) []int {
	dp := make([][]int, targetSum+1)
	dp[0] = []int{}
	exit := false

	for i := 0; i <= targetSum && !exit; i++ {
		if dp[i] == nil {
			continue
		}

		exit = true

		for _, number := range numbers {
			if number+i <= targetSum {
				exit = false

				if dp[number+i] == nil || len(dp[i])+1 < len(dp[number+i]) {
					dp[number+i] = append([]int{number}, dp[i]...)
				}
			}
		}
	}

	return dp[targetSum]
}

// m = len(target)
// n = len(wordBank)
// time: O(n * m^2)
// space: O(m)
func (t *tabulation) CanConstruct(target string, wordBank []string) bool {
	n := len(target)
	dp := make([]bool, n+1)
	dp[0] = true
	exit := false

	var length int

	for i := 0; i < n && !exit; i++ {
		if !dp[i] {
			continue
		}

		exit = true
		currTarget := target[i:]

		for _, word := range wordBank {
			length = len(word)
			if i+length <= n {
				exit = false

				if strings.Index(currTarget, word) == 0 {
					dp[i+length] = true
				}
			}
		}
	}

	return dp[n]
}

// m = len(target)
// n = len(numbers)
// time: O(n * m^2)
// space: O(m) - ?????
func (t *tabulation) CountConstruct(target string, wordBank []string) int {
	n := len(target)
	dp := make([]int, n+1)
	dp[0] = 1

	var length int

	exit := false
	for i := 0; i < n && !exit; i++ {
		if dp[i] == 0 {
			continue
		}

		exit = true
		currTarget := target[i:]

		for _, word := range wordBank {
			length = len(word)
			if i+length <= n {
				exit = false

				if strings.Index(currTarget, word) == 0 {
					dp[i+length] += dp[i]
				}
			}
		}
	}

	return dp[n]
}

// m = len(target)
// n = len(wordBank)
// time: O(n^m)
// space: O(n^m)
// Exponential Complexity!
func (t *tabulation) AllConstruct(target string, wordBank []string) [][]string {
	n := len(target)
	dp := make([][][]string, n+1)

	for i := 1; i <= n; i++ {
		dp[i] = [][]string{}
	}

	dp[0] = [][]string{{}}
	exit := false

	var length int

	for i := 0; i < n && !exit; i++ {
		if len(dp[i]) == 0 {
			continue
		}

		exit = true
		currTarget := target[i:]

		for _, word := range wordBank {
			length = len(word)
			if i+length <= n {
				exit = false

				if strings.Index(currTarget, word) == 0 {
					var combination [][]string
					for j, n := 0, len(dp[i]); j < n; j++ {
						combination = append(
							combination,
							append(append([]string{}, dp[i][j]...), word),
						)
					}

					dp[i+length] = append(dp[i+length], combination...)
				}
			}
		}
	}

	return dp[n]
}
