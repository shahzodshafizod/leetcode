import unittest

# https://leetcode.com/problems/reverse-string-ii/

# python3 -m unittest twopointers/0541-reverse-string-ii.py


class Solution(unittest.TestCase):
    # Approach: Two Pointers
    # Time: O(n)
    # Space: O(n), for b = list(s)
    def reverseStr(self, s: str, k: int) -> str:
        b, n = list(s), len(s)
        for left in range(0, n, 2 * k):
            right = min(n - 1, left + k - 1)
            while left < right:
                b[left], b[right] = b[right], b[left]
                left += 1
                right -= 1
        return "".join(b)

    def test(self):
        for s, k, expected in [
            ("abcd", 2, "bacd"),
            ("abcdefg", 2, "bacdfeg"),
            ("abcdefg", 8, "gfedcba"),
        ]:
            output = self.reverseStr(s, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
