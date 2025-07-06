import unittest

# https://leetcode.com/problems/lexicographically-smallest-equivalent-string/

# python3 -m unittest unionfinds/1061-lexicographically-smallest-equivalent-string.py


class Solution(unittest.TestCase):
    # Time: O(n+m)
    # Space: O(1)
    def smallestEquivalentString(self, s1: str, s2: str, baseStr: str) -> str:
        root = list(range(26))

        def find(x: int) -> int:
            if root[x] != x:
                root[x] = find(root[x])
            return root[x]

        orda = ord('a')
        for c1, c2 in zip(s1, s2):
            x = find(ord(c1) - orda)
            y = find(ord(c2) - orda)
            if x != y:
                root[max(x, y)] = min(x, y)
        return "".join(chr(orda + find(ord(c) - orda)) for c in baseStr)

    def test(self):
        for s1, s2, baseStr, expected in [
            ("parker", "morris", "parser", "makkek"),
            ("hello", "world", "hold", "hdld"),
            ("leetcode", "programs", "sourcecode", "aauaaaaada"),
        ]:
            output = self.smallestEquivalentString(s1, s2, baseStr)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
