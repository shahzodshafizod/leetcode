import unittest
import math

# https://leetcode.com/problems/count-square-sum-triples/

# python3 -m unittest maths/1925-count-square-sum-triples.py


class Solution(unittest.TestCase):
    # Approach 1: Brute Force - Try All Combinations
    # Try all (a, b, c) and check if a² + b² = c²
    # Time Complexity: O(n³)
    # Space Complexity: O(1)

    # # Approach 2: Optimized - Two Loops
    # # For each pair (a, b), check if a² + b² is perfect square <= n
    # # Time Complexity: O(n²)
    # # Space Complexity: O(1)
    # def countTriples(self, n: int) -> int:
    #     count = 0

    #     # Try all pairs (a, b)
    #     for a in range(1, n + 1):
    #         for b in range(1, n + 1):
    #             sum_squares = a * a + b * b

    #             # Check if sum is a perfect square
    #             c = int(math.sqrt(sum_squares))
    #             if c * c == sum_squares and c <= n:
    #                 count += 1

    #     return count

    # Alternative: For each c, find valid (a, b) pairs
    def countTriples(self, n: int) -> int:
        count = 0

        for c in range(1, n + 1):
            c_squared = c * c

            for a in range(1, c):
                a_squared = a * a
                b_squared = c_squared - a_squared

                # Check if b_squared is perfect square
                b = int(math.sqrt(b_squared))
                if b * b == b_squared and b <= n and b > 0:
                    count += 1

        return count

    def test(self):
        for n, expected in [
            (5, 2),  # (3,4,5), (4,3,5)
            (10, 4),  # (3,4,5), (4,3,5), (6,8,10), (8,6,10)
            (1, 0),
            (2, 0),
            (15, 8),  # 8 triples
        ]:
            output = self.countTriples(n)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
