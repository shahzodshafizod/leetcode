import unittest

# https://leetcode.com/problems/circular-sentence/

# python3 -m unittest strings/2490-circular-sentence.py

class Solution(unittest.TestCase):
    # # Time: O(n)
    # # Space: O(1)
    # def isCircularSentence(self, sentence: str) -> bool:
    #     for i in range(len(sentence)):
    #         if sentence[i] == ' ' and sentence[i-1] != sentence[i+1]:
    #             return False
    #     return sentence[0] == sentence[-1]

    # Faster
    # Time: O(n)
    # Space: O(n)
    def isCircularSentence(self, sentence: str) -> bool:
        words = sentence.split()
        for idx in range(len(words)):
            # Indices in Python are circular
            if words[idx][0] != words[idx-1][-1]:
                return False
        return True
    
    def test(self):
        for sentence, expected in [
            ("L", True),
            ("sus", True),
            ("eetcode", True),
            ("leetcode", False),
            ("a b c d e a", False),
            ("Leetcode is cool", False),
            ("MuFoevIXCZzrpXeRmTssj lYSW U jM", False),
            ("leetcode Exercises sound delightfu", False),
            ("leetcode exercises sound delightful", True),
            ("Leetcode Exercises sound delightful", False),
            ("IuTiUtGGsNydmacGduehPPGksKQyT TmOraUbCcQdnZUCpGCYtGp p pG GCcRvZDRawqGKOiBSLwjIDOjdhnHiisfddYoeHqxOqkUvOEyI", True),
        ]:
            output = self.isCircularSentence(sentence)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
