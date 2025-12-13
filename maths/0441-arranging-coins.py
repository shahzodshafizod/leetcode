import unittest
import math

# https://leetcode.com/problems/arranging-coins/

# python3 -m unittest maths/0441-arranging-coins.py


class Solution(unittest.TestCase):
    # # Approach #1: Brute Force - Simulation
    # # Time: O(sqrt(n)) - approximately sqrt(n) rows
    # # Space: O(1) - constant space
    # def arrangeCoins(self, n: int) -> int:
    #     row = 0
    #     while row + 1 <= n:
    #         row += 1
    #         n -= row
    #     return row

    # # Approach #2: Binary Search
    # # Time: O(log n) - binary search
    # # Space: O(1) - constant space
    # def arrangeCoins(self, n: int) -> int:
    #     # For each mid, check if mid * (mid + 1) / 2 <= n
    #     left, right = 0, n
    #     result = 0

    #     while left <= right:
    #         mid = (left + right) // 2
    #         total = mid * (mid + 1) // 2

    #         if total <= n:
    #             result = mid
    #             left = mid + 1
    #         else:
    #             right = mid - 1

    #     return result

    # Approach #3: Mathematical Formula (Quadratic Equation)
    # Time: O(1) - constant time calculation
    # Space: O(1) - constant space
    def arrangeCoins(self, n: int) -> int:
        # We need k such that 1 + 2 + ... + k <= n
        # Using formula: k * (k + 1) / 2 <= n
        # Solving quadratic: k^2 + k - 2n <= 0
        # Using quadratic formula: k = (-1 + sqrt(1 + 8n)) / 2
        return int((-1 + math.sqrt(1 + 8 * n)) / 2)

    def test(self):
        for n, expected in [
            (5, 2),  # rows: [1] [2,3] (4,5 incomplete)
            (8, 3),  # rows: [1] [2,3] [4,5,6] (7,8 incomplete)
            (1, 1),  # rows: [1]
            (3, 2),  # rows: [1] [2,3]
            (0, 0),  # no rows
            (10, 4),  # rows: [1] [2,3] [4,5,6] [7,8,9,10]
            (2147483647, 65535),  # large n
        ]:
            output = self.arrangeCoins(n)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
