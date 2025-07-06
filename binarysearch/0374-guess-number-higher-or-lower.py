import unittest

# https://leetcode.com/problems/guess-number-higher-or-lower/

# python3 -m unittest binarysearch/0374-guess-number-higher-or-lower.py

# The guess API is already defined for you.
# @param num, your guess
# @return -1 if num is higher than the picked number
#          1 if num is lower than the picked number
#          otherwise return 0
# def guess(num: int) -> int:


class Solution(unittest.TestCase):
    guessed = 0

    def guess(self, num: int) -> int:
        if num > self.guessed:
            return -1
        if num < self.guessed:
            return 1
        return 0

    def guessNumber(self, n: int) -> int:
        left, right = 1, n
        while left <= right:
            num = (left + right) // 2
            guessed = self.guess(num)
            if guessed == -1:
                right = num - 1
            elif guessed == 1:
                left = num + 1
            else:
                break
        return num

    def test(self):
        for n, picked, expected in [
            (10, 6, 6),
            (1, 1, 1),
            (2, 1, 1),
        ]:
            self.guessed = picked
            output = self.guessNumber(n)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
