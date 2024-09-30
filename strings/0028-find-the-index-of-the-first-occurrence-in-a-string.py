import unittest

# https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/

# python3 -m unittest strings/0028-find-the-index-of-the-first-occurrence-in-a-string.py

class Solution(unittest.TestCase):
    # def strStr(self, haystack: str, needle: str) -> int:
    #     return haystack.find(needle)

    # # Approach: KMP (Knuth-Morris-Pratt) Algorithm
    # # Time: O(M+N)
    # # Space: O(N)
    # def strStr(self, haystack: str, needle: str) -> int:
    #     m = len(needle)
    #     lps = [0] * m
    #     lps[0] = 0
    #     prevLPS = 0
    #     idx = 1
    #     while idx < m:
    #         if needle[idx] == needle[prevLPS]:
    #             lps[idx] = prevLPS+1
    #             prevLPS += 1
    #             idx += 1
    #         elif prevLPS == 0:
    #             idx += 1
    #         else:
    #             prevLPS = lps[prevLPS-1]
    #     n = len(haystack)
    #     hi, ni = 0, 0
    #     while hi < n:
    #         if haystack[hi] == needle[ni]:
    #             hi += 1
    #             ni += 1
    #         elif ni == 0:
    #             hi += 1
    #         else:
    #             ni = lps[ni-1]
    #         if ni == m:
    #             return hi - m
    #     return -1

    # Approach: Rabin Karp Algorithm
    # Time: O(M+N)
    # Space: O(1)
    def strStr(self, haystack: str, needle: str) -> int:
        base, mod = 29, int(1e9+7)
        subs = 0
        for c in needle:
            digit = ord(c) - ord('a') + 1
            subs = (subs * base + digit) % mod
        m = len(needle)
        power = base ** (m-1)
        s = 0
        for idx, c in enumerate(haystack):
            digit = ord(c) - ord('a') + 1
            s = (s * base + digit) % mod
            if idx+1 < m:
                continue
            if s == subs:
                return idx+1 - m
            digit = ord(haystack[idx+1-m]) - ord('a') + 1
            s -= (digit * power) % mod
        return -1

    def test(self) -> None:
        for haystack, needle, expected in [
            ("sadbutsad", "sad", 0),
            ("leetcode", "leeto", -1),
            ("sydbutsad", "sad", 6),
            ("aabaaabaaac", "aabaaac", 4),
            ("ababcaababcaabc", "ababcaabc", 6),
            # ("fourscoreandsevenyearsagoourfathersbroughtforthuponthiscontinentanewnation", "ourfathersbroughtforthuponthiscontinentanewnation", 25),
        ]:
            output = self.strStr(haystack, needle)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
