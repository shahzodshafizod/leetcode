import unittest

# https://leetcode.com/problems/count-and-say/

# python3 -m unittest strings/0038-count-and-say.py


class Solution(unittest.TestCase):
    # # Approach: Recursive
    # # Time: O(nxm), m=max(len(RLE))
    # # Space: O(n)
    # def countAndSay(self, n: int) -> str:
    #     if n == 1: return "1"
    #     s = self.countAndSay(n-1)
    #     rle, count = "", 1
    #     for idx in range(1, len(s)):
    #         if s[idx] != s[idx-1]:
    #             rle = f"{rle}{count}{s[idx-1]}"
    #             count = 0
    #         count += 1
    #     return f"{rle}{count}{s[-1]}"

    # Approach: Iterative
    # Time: O(n)
    # Space: O(n)
    def countAndSay(self, n: int) -> str:
        rle = "1"
        for _ in range(n - 1):
            new_rle, count = "", 1
            for idx in range(1, len(rle)):
                if rle[idx] != rle[idx - 1]:
                    new_rle = f"{new_rle}{count}{rle[idx-1]}"
                    count = 0
                count += 1
            rle = f"{new_rle}{count}{rle[-1]}"
        return rle

    def test(self):
        for n, expected in [
            (4, "1211"),
            (1, "1"),
        ]:
            output = self.countAndSay(n)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
