from collections import Counter
import unittest

# https://leetcode.com/problems/ransom-note/

# python3 -m unittest hashes/0383-ransom-note.py

class Solution(unittest.TestCase):
    # Approach: Counting via Hash Table
    # Time: O(n)
    # Space: O(26) = O(1)
    def canConstruct(self, ransomNote: str, magazine: str) -> bool:
        count = Counter(magazine)
        for c in ransomNote:
            count[c] -= 1
            if count[c] < 0:
                return False
        return True

    def test(self):
        for ransomNote, magazine, expected in [
            ("a", "b", False),
            ("aa", "ab", False),
            ("aa", "aab", True),
        ]:
            output = self.canConstruct(ransomNote, magazine)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
