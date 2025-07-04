from typing import List
import unittest

# https://leetcode.com/problems/find-the-k-th-character-in-string-game-ii/

# python3 -m unittest bits/3307-find-the-k-th-character-in-string-game-ii.py


class Solution(unittest.TestCase):
    # Approach: Mathematics
    # Time: O(logk)
    # Space: O(1)
    def kthCharacter(self, k: int, operations: List[int]) -> str:
        ans = 0
        for bits in range(k.bit_length(), -1, -1):
            if k > 1 << bits:
                k -= 1 << bits
                ans += operations[bits]
        return chr(ord('a') + ans % 26)

    def test(self):
        for k, operations, expected in [
            (5, [0, 0, 0], "a"),
            (10, [0, 1, 0, 1], "b"),
            (33031255, [0, 1, 0, 0, 1, 1, 1, 0, 1, 0, 1, 1, 1, 0, 0, 1, 0, 0, 1, 1, 1, 1, 0, 1, 1], "j"),
            (14493383, [0, 0, 1, 1, 0, 1, 0, 1, 1, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 0, 1, 0, 1, 1, 1], "i"),
            (32961520, [1, 1, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 0, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1], "k"),
        ]:
            output = self.kthCharacter(k, operations)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
