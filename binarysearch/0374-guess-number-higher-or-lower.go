package binarysearch

// https://leetcode.com/problems/guess-number-higher-or-lower/

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

var GUESSED__NUMBER__ int

func guess(num int) int {
	if num > GUESSED__NUMBER__ {
		return -1
	}

	if num < GUESSED__NUMBER__ {
		return 1
	}

	return 0
}

func guessNumber(n int) int {
	left, right := 1, n

	var num, guessed int

	for left <= right {
		num = left + (right-left)/2
		guessed = guess(num)

		if guessed == -1 {
			right = num - 1
		} else if guessed == 1 {
			left = num + 1
		} else {
			break
		}
	}

	return num
}
