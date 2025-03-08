from math import sqrt
from typing import List
import unittest

# https://leetcode.com/problems/closest-prime-numbers-in-range/

# python3 -m unittest maths/2523-closest-prime-numbers-in-range.py

class Solution(unittest.TestCase):
    # Approach: Sieve of Eratosthenes
    # Time: O(N∗Log(Log(N)))
    # Space: O(N)
    def closestPrimes(self, left: int, right: int) -> List[int]:
        prime = [True] * (right+1)
        prime[0] = prime[1] = False
        for num in range(2, int(sqrt(right))+1):
            if prime[num]:
                for factor in range(num+num, right+1, num):
                    prime[factor] = False
            num += 1
        ans = [-1, -1]
        diff = right-left+1
        prev = -1
        for curr in range(left, right+1):
            if not prime[curr]:
                continue
            if prev != -1 and curr - prev < diff:
                diff = curr - prev
                ans = [prev, curr]
            prev = curr
        return ans

    def test(self):
        for left, right, expected in [
            (10, 19, [11,13]),
            (4, 6, [-1,-1]),
            (1, 2, [-1,-1]),
        ]:
            output = self.closestPrimes(left, right)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
