import unittest

# https://leetcode.com/problems/count-good-numbers/

# python3 -m unittest maths/1922-count-good-numbers.py

class Solution(unittest.TestCase):
    # Approach: Binary Exponentiation
    # Time: O(logn)
    # Space: O(logn)
    def countGoodNumbers(self, n: int) -> int:
        MOD = 10**9 + 7
        def pow(x: int, n: int) -> int:
            if n == 0: return 1
            half = pow(x, n>>1)
            if n&1:
                return half * half * x % MOD
            return half * half % MOD
        evens = pow(5, n-n//2)
        primes = pow(4, n>>1)
        return (evens * primes) % MOD

    def test(self):
        for n, expected in [
            (1, 5),
            (4, 400),
            (50, 564908303),
        ]:
            output = self.countGoodNumbers(n)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
