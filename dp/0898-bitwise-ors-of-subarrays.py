from typing import List, Set
import unittest

# https://leetcode.com/problems/bitwise-ors-of-subarrays/

# python3 -m unittest dp/0898-bitwise-ors-of-subarrays.py


class Solution(unittest.TestCase):
    def subarrayBitwiseORs(self, arr: List[int]) -> int:
        res: Set[int] = set()
        cur: Set[int] = set()
        for num in arr:
            cur = {num | x for x in cur} | {num}
            res |= cur
        return len(res)

    def test(self):
        for arr, expected in [
            ([0], 1),
            ([1, 1, 2], 3),
            ([1, 2, 4], 6),
        ]:
            output = self.subarrayBitwiseORs(arr)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
