from bisect import bisect_left, bisect_right
import unittest
from typing import List

# https://leetcode.com/problems/shortest-matching-substring/

# python3 -m unittest strings/3455-shortest-matching-substring.py


class Solution(unittest.TestCase):
    # # Approach #1: Pattern matching without binary search (TLE)
    # # Time: O(n^2) where n = len(s)
    # # Space: O(1)
    # def minValidSubstring(self, s: str, p: str) -> int:
    #     parts = p.split("*")
    #     if len(parts) != 3:
    #         return -1

    #     prefix, middle, suffix = parts

    #     if not prefix and not middle and not suffix:
    #         return 0

    #     n = len(s)
    #     min_len = float("inf")

    #     for start in range(n):
    #         if prefix:
    #             if start + len(prefix) > n:
    #                 continue
    #             if s[start : start + len(prefix)] != prefix:
    #                 continue
    #             pos = start + len(prefix)
    #         else:
    #             pos = start

    #         if middle:
    #             mid_pos = s.find(middle, pos)
    #             if mid_pos == -1:
    #                 continue
    #             pos = mid_pos + len(middle)

    #         if suffix:
    #             suf_pos = s.find(suffix, pos)
    #             if suf_pos == -1:
    #                 continue
    #             end = suf_pos + len(suffix)
    #         else:
    #             end = pos

    #         min_len = min(min_len, end - start)

    #     return int(min_len) if min_len != float("inf") else -1

    # Approach #2: Binary Search with Pattern Matching
    # Time: O(n log n) where n = len(s)
    # Space: O(n) for storing occurrences
    def minValidSubstring(self, s: str, p: str) -> int:
        # Split pattern into A*B*C
        parts = p.split('*')
        A, B, C = parts[0], parts[1], parts[2]
        a, b, c = len(A), len(B), len(C)

        # KMP to find all occurrences of a pattern in text
        def kmp_occurrences(text: str, pat: str) -> List[int]:
            m = len(pat)
            if m == 0:
                return []  # we won't call this for empty patterns
            # Build LPS
            lps = [0] * m
            j = 0
            for i in range(1, m):
                while j > 0 and pat[i] != pat[j]:
                    j = lps[j - 1]
                if pat[i] == pat[j]:
                    j += 1
                    lps[i] = j
            # Search
            res: List[int] = []
            j = 0
            for i in range(len(text)):
                while j > 0 and text[i] != pat[j]:
                    j = lps[j - 1]
                if text[i] == pat[j]:
                    j += 1
                    if j == m:
                        res.append(i - m + 1)
                        j = lps[j - 1]
            return res

        # Occurrences lists for non-empty tokens
        posA = kmp_occurrences(s, A) if a > 0 else []
        posB = kmp_occurrences(s, B) if b > 0 else []
        posC = kmp_occurrences(s, C) if c > 0 else []

        # Early impossibility checks
        if a > 0 and not posA:
            return -1
        if b > 0 and not posB:
            return -1
        if c > 0 and not posC:
            return -1

        INF = 10**18
        ans = INF

        # Case breakdown

        # Case A empty, C empty
        if a == 0 and c == 0:
            if b == 0:
                return 0  # "**"
            else:
                return b if posB else -1  # "*B*"

        # Case A non-empty, C empty
        if a > 0 and c == 0:
            if b == 0:
                # Pattern A* -> shortest is just A if it exists
                return a
            else:
                # For each A start, take earliest B at/after i+a
                for i in posA:
                    idxB = bisect_left(posB, i + a)
                    if idxB == len(posB):
                        continue
                    bStart = posB[idxB]
                    cand = bStart + b - i
                    if cand < ans:
                        ans = cand
                return ans if ans < INF else -1

        # Case A empty, C non-empty
        if a == 0 and c > 0:
            if b == 0:
                # Pattern *C -> shortest is just C if it exists
                return c
            else:
                # For each C start, pick latest B start with bStart + b <= cStart
                for cStart in posC:
                    xMax = cStart - b
                    if xMax < 0:
                        continue
                    idxB = bisect_right(posB, xMax) - 1
                    if idxB < 0:
                        continue
                    bStart = posB[idxB]
                    cand = cStart + c - bStart
                    if cand < ans:
                        ans = cand
                return ans if ans < INF else -1

        # Case A non-empty, C non-empty
        # General form: A * B * C
        if b == 0:
            # Reduce to A * C: for each A start, take earliest C start >= i + a
            for i in posA:
                idxC = bisect_left(posC, i + a)
                if idxC == len(posC):
                    continue
                cStart = posC[idxC]
                cand = cStart + c - i
                if cand < ans:
                    ans = cand
            return ans if ans < INF else -1
        else:
            # For each A start, take earliest B at/after i+a, then earliest C at/after bStart+b
            for i in posA:
                idxB = bisect_left(posB, i + a)
                if idxB == len(posB):
                    continue
                bStart = posB[idxB]
                y = bStart + b
                idxC = bisect_left(posC, y)
                if idxC == len(posC):
                    continue
                cStart = posC[idxC]
                cand = cStart + c - i
                if cand < ans:
                    ans = cand
            return ans if ans < INF else -1

    def test(self):
        for s, pattern, expected in [
            # Valid test cases (pattern has exactly two '*')
            ("abaacbaecebce", "ba*c*ce", 8),
            ("baccbaadbc", "cc*baa*adb", -1),
            ("a", "**", 0),
            ("madlogic", "*adlogi*", 6),
            ("hy", "h**", 1),
        ]:
            output = self.minValidSubstring(s, pattern)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
