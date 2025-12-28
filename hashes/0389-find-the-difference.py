import unittest

# https://leetcode.com/problems/find-the-difference/

# python3 -m unittest hashes/0389-find-the-difference.py


class Solution(unittest.TestCase):
    # # Approach: Counter via Hash Set
    # # Time: O(n)
    # # Space: O(26) = O(1)
    # def findTheDifference(self, s: str, t: str) -> str:
    #     return list((Counter(t)-Counter(s)).keys())[0]

    # Approach: Bit Manipulation
    # Time: O(n)
    # Space: O(1)
    def findTheDifference(self, s: str, t: str) -> str:
        extra = 0
        for sc, tc in zip(s, t):
            extra ^= ord(sc)
            extra ^= ord(tc)
        extra ^= ord(t[len(t) - 1])
        return chr(extra)

    def testFindTheDifference(self) -> None:
        for s, t, expected in [("abcd", "abcde", "e"), ("", "y", "y"), ("a", "aa", "a")]:
            output = self.findTheDifference(s, t)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
