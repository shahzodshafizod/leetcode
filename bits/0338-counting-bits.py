import unittest
from typing import List

# https://leetcode.com/problems/counting-bits/

# python3 -m unittest bits/0338-counting-bits.py


class Solution(unittest.TestCase):
    def countBits(self, n: int) -> List[int]:
        # mask = [0] * 32
        # ans = [0] * (n + 1)
        # count = 0
        # for i in range(1, n+1):
        #     idx = 0
        #     while idx < 32 and mask[idx] == 1:
        #         mask[idx] = 0
        #         idx += 1
        #         count -= 1
        #     mask[idx] = 1
        #     count += 1
        #     ans[i] = count
        # return ans
        ans = [0] * (n + 1)
        for i in range(1, n + 1):
            ans[i] = ans[i >> 1] + (i & 1)
        return ans

    def testCountBits(self) -> None:
        for n, expected in [
            (2, [0, 1, 1]),
            (5, [0, 1, 1, 2, 1, 2]),
            (8, [0, 1, 1, 2, 1, 2, 2, 3, 1]),
        ]:
            output = self.countBits(n)
            self.assertListEqual(expected, output, f"expected: {expected}, output: {output}")
