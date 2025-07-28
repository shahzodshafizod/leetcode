package backtracking

// https://leetcode.com/problems/find-the-count-of-good-integers/

func countGoodIntegers(n int, k int) int64 {
	// 1. collecting all k-palindromic integer containing n digits
	palindromes := make(map[int]struct{})

	var backtrack func(number int, digit int, left int, right int)

	backtrack = func(number int, digit int, left int, right int) {
		if left < right {
			if number%k == 0 {
				// collecting sorted numbers by digits
				var freq [10]int
				for number > 0 {
					freq[number%10]++
					number /= 10
				}

				for digit, cnt := range freq {
					for ; cnt > 0; cnt-- {
						number = number*10 + digit
					}
				}

				palindromes[number] = struct{}{}
			}

			return
		}

		if digit > 9 {
			return
		}

		if left == right {
			number += digit * left
			backtrack(number, 0, left/10, right*10)
			number -= digit * left
			backtrack(number, digit+1, left, right)
		} else {
			number += digit*left + digit*right
			backtrack(number, 0, left/10, right*10)
			number -= digit*left + digit*right
			backtrack(number, digit+1, left, right)
		}
	}

	left := 1
	for idx := n - 1; idx > 0; idx-- {
		left *= 10
	}

	backtrack(0, 1, left, 1)
	// 2. precalculation of factorials
	var facts [11]int64

	facts[0] = 1
	for num := 1; num <= 10; num++ {
		facts[num] = facts[num-1] * int64(num)
	}
	// 3. calculating valid (non 0-leading) permutations
	var total int64 = 0

	var count int64

	for number := range palindromes {
		var freq [10]int
		for idx := 0; idx < n; idx++ {
			freq[number%10]++
			number /= 10
		}

		count = int64(n-freq[0]) * facts[n-1]
		for _, f := range freq {
			count /= facts[f]
		}

		total += count
	}

	return total
}
