import unittest

# https://leetcode.com/problems/lexicographically-minimum-string-after-removing-stars/

# python3 -m unittest greedy/3170-lexicographically-minimum-string-after-removing-stars.py


class Solution(unittest.TestCase):
    # Approach: Greedy with Stack
    # Time: O(n)
    # Space: O(n)
    def clearStars(self, s: str) -> str:
        stack = [[] for _ in range(26)]
        cleared = list(s)
        for idx, c in enumerate(s):
            if c != '*':
                stack[ord(c) - ord('a')].append(idx)
            else:
                cleared[idx] = ''
                for letter in range(26):
                    if stack[letter]:
                        cleared[stack[letter].pop()] = ''
                        break
        return "".join(cleared)

    def test(self):
        for s, expected in [
            ("aaba*", "aab"),
            ("abc", "abc"),
            ("abc*de*fgh*", "defgh"),
            ("a*b*c*d*", ""),
            ("abcde*f*", "cdef"),
            ("abc*", "bc"),
            ("aaaa***", "a"),
            ("de*", "e"),
            ("d*yed", "yed"),
        ]:
            output = self.clearStars(s)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
