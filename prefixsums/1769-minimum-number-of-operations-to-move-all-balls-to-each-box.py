from typing import List
import unittest

# https://leetcode.com/problems/minimum-number-of-operations-to-move-all-balls-to-each-box/

# python3 -m unittest prefixsums/1769-minimum-number-of-operations-to-move-all-balls-to-each-box.py


class Solution(unittest.TestCase):
    def minOperations(self, s: str) -> List[int]:
        n = len(s)
        right, rboxes = 0, 0
        for idx in range(n - 1, -1, -1):
            if s[idx] == '1':
                rboxes += 1
            right += rboxes
        answer = [0] * n
        left, lboxes = 0, 0
        for idx in range(n):
            left += lboxes
            right -= rboxes
            answer[idx] = left + right
            if s[idx] == '1':
                lboxes += 1
                rboxes -= 1
        return answer

    def test(self):
        for s, expected in [
            ("110", [1, 1, 3]),
            ("001011", [11, 8, 5, 4, 3, 4]),
        ]:
            output = self.minOperations(s)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
