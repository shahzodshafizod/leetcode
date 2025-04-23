import math
import unittest

# https://leetcode.com/problems/count-the-number-of-ideal-arrays/

# python3 -m unittest dp/2338-count-the-number-of-ideal-arrays.py

class Solution(unittest.TestCase):
    def idealArrays(self, n: int, maxValue: int) -> int:
        combs = [0] * 15
        for k in range(1, 15):
            combs[k] = math.comb(n-1,k-1)
        memo = [[None]*15 for _ in range(maxValue+1)]
        def dp(num: int, k: int) -> int:
            if memo[num][k] != None:
                return memo[num][k]
            # num is in the last pos, combs[k] is the 
            # combination of n-1 previous positions
            # k is the number of items to be chosen
            res = combs[k]
            nxt = num * 2
            while nxt <= maxValue:
                res += dp(nxt, k+1)
                nxt += num
            memo[num][k] = res
            return res
        return sum(dp(i,1) for i in range(1, maxValue+1)) % (10**9 + 7)

    def test(self):
        for n, maxValue, expected in [
            (2, 5, 10),
            (5, 3, 11),
        ]:
            output = self.idealArrays(n, maxValue)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
