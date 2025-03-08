import unittest

# https://leetcode.com/problems/minimum-recolors-to-get-k-consecutive-black-blocks/

# python3 -m unittest slidingwindows/2379-minimum-recolors-to-get-k-consecutive-black-blocks.py

class Solution(unittest.TestCase):
    def minimumRecolors(self, blocks: str, k: int) -> int:
        res = float("inf")
        ops, k = 0, k-1
        for idx in range(len(blocks)):
            if blocks[idx] == 'W':
                ops += 1
            if idx-k >= 0:
                res = min(res, ops)
                if blocks[idx-k] == 'W':
                    ops -= 1
        return res

    def test(self):
        for blocks, k, expected in [
            ("WBBWWBBWBW", 7, 3),
		    ("WBWBBBW", 2, 0),
            ("BWWWBB", 6, 3),
        ]:
            output = self.minimumRecolors(blocks, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
