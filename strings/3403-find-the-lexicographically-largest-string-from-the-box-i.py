import unittest

# https://leetcode.com/problems/find-the-lexicographically-largest-string-from-the-box-i/

# python3 -m unittest strings/3403-find-the-lexicographically-largest-string-from-the-box-i.py

class Solution(unittest.TestCase):
    # Approach 1: Enumeration
    # Time: O(nk), n=len(word), k=n-numFriends+1
    # Space: O(1)
    def answerString(self, word: str, numFriends: int) -> str:
        if numFriends == 1: return word
        n = len(word)
        k = n - numFriends + 1
        return max( word[i:i+k] for i in range(n) )

    def test(self):
        for word, numFriends, expected in [
            ("dbca", 2, "dbc"),
            ("gggg", 4, "g"),
            ("aann", 2, "nn"),
            ("gh", 1, "gh"),
            ("yxz", 3, "z"),
            ("xzy", 2, "zy"),
            ("xxyy", 2, "yy"),
            ("xy", 1, "xy"),
            ("xzyz", 3, "zy"),
            ("adif", 1, "adif"),
            ("bagj", 3, "j"),
            ("gm", 1, "gm"),
            ("gm", 2, "m"),
            ("gmn", 3, "n"),
            ("gmhm", 3, "mh"),
            ("gmng", 2, "ng"),
            ("gmnaangggg", 4, "ngggg"),
        ]:
            output = self.answerString(word, numFriends)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
