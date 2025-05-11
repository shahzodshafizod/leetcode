import unittest

# https://leetcode.com/problems/license-key-formatting/

# python3 -m unittest strings/0482-license-key-formatting.py

class Solution(unittest.TestCase):
    # Approach: Left to Right Traversal
    # Time: O(n)
    # Space: O(1)
    def licenseKeyFormatting(self, s: str, k: int) -> str:
        count = (len(s) - s.count('-')) % k
        count = k if count == 0 else count
        res = []
        for c in s:
            if c != '-':
                if count == 0:
                    res.append('-')
                    count = k
                res.append(c.upper())
                count -= 1
        return "".join(res)

    def test(self):
        for s, k, expected in [
            ("5F3Z-2e-9-w", 4, "5F3Z-2E9W"),
            ("2-5g-3-J", 2, "2-5G-3J"),
            ("---", 3, ""),
            ("2-4A0r7-4k", 4, "24A0-R74K"),
            ("5F3Z-2e-9-w", 44, "5F3Z2E9W"),
            ("--a-a-a-a--", 2, "AA-AA"),
        ]:
            output = self.licenseKeyFormatting(s, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
