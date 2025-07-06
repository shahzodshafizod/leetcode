from typing import List
from collections import Counter, deque
import unittest

# https://leetcode.com/problems/longest-subsequence-repeated-k-times/

# python3 -m unittest strings/2014-longest-subsequence-repeated-k-times.py


class Solution(unittest.TestCase):
    # Approach: Backtracking + Breadth-First Search
    # Time: O(7! * n)
    # Space: O(7!)
    def longestSubsequenceRepeatedK(self, s: str, k: int) -> str:
        def calc_count(subseq: str) -> int:
            i, n = 0, len(subseq)
            count = 0
            for c in s:
                if c == subseq[i]:
                    i += 1
                    if i == n:
                        count += 1
                        i = 0
            return count

        candidates = sorted([c for c, cnt in Counter(s).items() if cnt >= k])
        subseq = ""
        queue = deque([subseq])
        while queue:
            subseq = queue.popleft()
            for c in candidates:
                if calc_count(subseq + c) >= k:
                    queue.append(subseq + c)
        return subseq

    def test(self):
        for s, k, expected in [
            ("bb", 2, "b"),
            ("ab", 2, ""),
            ("gbjbjigjbji", 2, "gjbji"),
            ("letsleetcode", 2, "let"),
            ("nhsbbeonhsbbeonhsbbeo", 3, "nhsbbeo"),
        ]:
            output = self.longestSubsequenceRepeatedK(s, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
