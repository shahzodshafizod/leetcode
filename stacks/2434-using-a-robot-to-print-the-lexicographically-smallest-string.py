from collections import defaultdict
import unittest

# https://leetcode.com/problems/using-a-robot-to-print-the-lexicographically-smallest-string/

# python3 -m unittest stacks/2434-using-a-robot-to-print-the-lexicographically-smallest-string.py


class Solution(unittest.TestCase):
    # Approach: Greedy + Stack
    # Time: O(n)
    # Space: O(n)
    def robotWithString(self, s: str) -> str:
        cnt = defaultdict(int)
        for c in s:
            cnt[c] += 1
        t, paper = [], []
        min_char = "a"
        for c in s:
            t.append(c)
            cnt[c] -= 1
            while min_char != "z" and cnt[min_char] == 0:
                min_char = chr(ord(min_char) + 1)
            while t and t[-1] <= min_char:
                paper.append(t.pop())
        return "".join(paper)

    def test(self):
        for s, expected in [
            ("zza", "azz"),
            ("bac", "abc"),
            ("bdda", "addb"),
        ]:
            output = self.robotWithString(s)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
