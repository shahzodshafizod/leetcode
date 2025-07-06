import unittest

# https://leetcode.com/problems/find-the-count-of-good-integers/

# python3 -m unittest backtracking/3272-find-the-count-of-good-integers.py


class Solution(unittest.TestCase):
    def countGoodIntegers(self, n: int, k: int) -> int:
        # 1. collecting all k-palindromic integer containing n digits
        palindromes = set()

        def backtrack(number: int, digit: int, left: int, right: int):
            if left < right:
                if number % k == 0:
                    # collecting sorted numbers by digits
                    palindromes.add("".join(sorted(str(number))))
                return
            if digit > 9:
                return
            if left == right:
                number += digit * left
                backtrack(number, 0, left // 10, right * 10)
                number -= digit * left
                backtrack(number, digit + 1, left, right)
                return
            number += digit * left + digit * right
            backtrack(number, 0, left // 10, right * 10)
            number -= digit * left + digit * right
            backtrack(number, digit + 1, left, right)

        left = 1
        for _ in range(n - 1):
            left *= 10
        backtrack(0, 1, left, 1)
        # 2. precalculation of factorials
        fact = [0] * (n + 1)
        fact[0] = 1
        for num in range(1, n + 1):
            fact[num] = fact[num - 1] * num
        # 3. calculating valid (non 0-leading) permutations
        total = 0
        for snum in palindromes:
            freq = [0] * 10
            for c in snum:
                freq[int(c)] += 1
            count = (n - freq[0]) * fact[n - 1]
            for f in freq:
                count //= fact[f]
            total += count
        return total

    def test(self):
        for n, k, expected in [
            (3, 5, 27),
            (1, 4, 2),
            (5, 6, 2468),
        ]:
            output = self.countGoodIntegers(n, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
