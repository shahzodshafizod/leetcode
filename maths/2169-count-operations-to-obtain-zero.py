import unittest

# https://leetcode.com/problems/count-operations-to-obtain-zero/

# python3 -m unittest maths/2169-count-operations-to-obtain-zero.py


class Solution(unittest.TestCase):
    # # Approach #1: Simulation
    # # Time: O(max(num1, num2))
    # # Space: O(1)
    # def countOperations(self, num1: int, num2: int) -> int:
    #     ops = 0
    #     while num1 and num2:
    #         if num1 >= num2:
    #             num1 -= num2
    #         else:
    #             num2 -= num1
    #         ops += 1
    #     return ops

    # Approach #2: Euclidean Algorithm
    # Time: O(log(max(num1, num2)))
    # Space: O(1)
    def countOperations(self, num1: int, num2: int) -> int:
        ops = 0
        while num1 and num2:
            ops += num1 // num2
            num1 %= num2
            num1, num2 = num2, num1
        return ops

    def test(self):
        for num1, num2, expected in [
            (2, 3, 3),
            (10, 10, 1),
        ]:
            output = self.countOperations(num1, num2)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
