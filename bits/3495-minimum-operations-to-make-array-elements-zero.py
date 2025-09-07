from typing import List
import unittest

# https://leetcode.com/problems/minimum-operations-to-make-array-elements-zero/

# python3 -m unittest bits/3495-minimum-operations-to-make-array-elements-zero.py


class Solution(unittest.TestCase):
    # # Approach #1
    # # Time: O(32*n)
    # # Space: O(1)
    # def minOperations(self, queries: List[List[int]]) -> int:
    #     def get(num: int) -> int:  # O(32)
    #         cnt: int = 0
    #         base, bits = 1, 1
    #         while base <= num:
    #             cnt += (min(base * 2 - 1, num) - base + 1) * ((bits + 1) // 2)
    #             bits += 1
    #             base <<= 1
    #         return cnt

    #     return sum((get(r) - get(l - 1) + 1) // 2 for l, r in queries)

    # # Approach #2
    # # Time: O(16*n)
    # # Space: O(1)
    # def minOperations(self, queries: List[List[int]]) -> int:
    #     def get(num: int) -> int:
    #         cnt: int = 0
    #         base, quats = 1, 1
    #         while base <= num:
    #             cnt += (min(base * 4 - 1, num) - base + 1) * quats
    #             quats += 1
    #             base <<= 2
    #         return cnt

    #     return sum((get(r) - get(l - 1) + 1) // 2 for l, r in queries)

    # Approach #3
    # Time: O(16*n)
    # Space: O(1)
    def minOperations(self, queries: List[List[int]]) -> int:
        cnt: int = 0
        for left, right in queries:
            ops = 0
            for d in range(1, 17):
                start = 1 << (2 * (d - 1))
                end = (start << 2) - 1
                start = max(start, left)
                end = min(end, right)
                if start <= end:
                    ops += (end - start + 1) * d
            cnt += (ops + 1) // 2
        return cnt

    def test(self):
        for queries, expected in [
            ([[1, 2], [2, 4]], 3),
            ([[2, 6]], 4),
            ([[1, 8]], 7),
        ]:
            output = self.minOperations(queries)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
