import unittest

# https://leetcode.com/problems/repeated-substring-pattern/

# python3 -m unittest strings/0459-repeated-substring-pattern.py


class Solution(unittest.TestCase):
    # Approach: KMP (Knuth-Morris-Pratt) Algorithm
    # Time: O(n)
    # Space: O(n)
    def repeatedSubstringPattern(self, s: str) -> bool:
        # return s in (s+s)[1:2*len(s)-1]
        n = len(s)
        lps = [0] * n
        pref_len, idx = 0, 1
        while idx < n:
            if s[idx] == s[pref_len]:
                pref_len += 1
                lps[idx] = pref_len
            elif pref_len > 0:
                pref_len = lps[pref_len - 1]
                idx -= 1
            idx += 1
        return lps[-1] != 0 and lps[-1] % (n - lps[-1]) == 0

    def test(self):
        for s, expected in [
            ("a", False),
            ("aa", True),
            ("aba", False),
            ("abab", True),
            ("abac", False),
            ("abaababaab", True),
            ("abcabcabcabc", True),
            ("babbabbabbabbab", True),
            ("ababababababaababababababaababababababa", True),
        ]:
            output = self.repeatedSubstringPattern(s)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
