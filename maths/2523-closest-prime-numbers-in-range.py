from typing import List
import unittest

# https://leetcode.com/problems/closest-prime-numbers-in-range/

# python3 -m unittest maths/2523-closest-prime-numbers-in-range.py

class Solution(unittest.TestCase):
    # Approach: Sieve of Eratosthenes
    # Time: O(N∗Log(Log(N)))
    # Space: O(N)
    def closestPrimes(self, left: int, right: int) -> List[int]:
        sieve = [True] * (right+1)
        sieve[0] = sieve[1] = False
        primes = []
        for num in range(2, right+1):
            if sieve[num]:
                for factor in range(num*num, right+1, num):
                    sieve[factor] = False
                if num >= left:
                    primes.append(num)
        ans = [-1, -1]
        diff = right-left+1
        prev = -1
        for curr in primes:
            if prev != -1 and curr - prev < diff:
                diff = curr - prev
                ans = [prev, curr]
            prev = curr
        return ans

    def test(self):
        for left, right, expected in [
            (10, 19, [11,13]),
            (4, 6, [-1,-1]),
        ]:
            output = self.closestPrimes(left, right)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
