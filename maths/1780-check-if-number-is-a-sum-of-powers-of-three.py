import unittest

# https://leetcode.com/problems/check-if-number-is-a-sum-of-powers-of-three/

# python3 -m unittest maths/1780-check-if-number-is-a-sum-of-powers-of-three.py

class Solution(unittest.TestCase):
    # # Approach #1: Backtracking
    # # Time: O(2 ^ log3(n))
    # # Space: O( log3(n) )
    # def checkPowersOfThree(self, n: int) -> bool:
    #     def backtrack(x: int, cur_sum: int) -> bool:
    #         if cur_sum == n:
    #             return True
    #         if cur_sum > n or 3**x > n:
    #             return False
    #         # include
    #         if backtrack(x+1, cur_sum + 3**x):
    #             return True
    #         # skip
    #         return backtrack(x+1, cur_sum)
    #     return backtrack(0, 0)
    
    # # Approach #2: Iterative (Optimized)
    # # Time: O( log3(n) )
    # # Space: O(1)
    # def checkPowersOfThree(self, n: int) -> bool:
    #     power = 1
    #     while power*3 <= n:
    #         power *= 3
    #     while n and power > 0:
    #         if n >= power:
    #             n -= power
    #         power //= 3
    #     return n == 0

    # Approach #3: Ternary Representation
    # Time: O( log3(n) )
    # Space: O(1)
    def checkPowersOfThree(self, n: int) -> bool:
        while n:
            if n % 3 == 2:
                return False
            n //= 3
        return True

    def test(self):
        for n, expected in [
            (12, True),
            (91, True),
            (21, False),
        ]:
            output = self.checkPowersOfThree(n)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
