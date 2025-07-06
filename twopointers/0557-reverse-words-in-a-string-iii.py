import unittest

# https://leetcode.com/problems/reverse-words-in-a-string-iii/

# python3 -m unittest twopointers/0557-reverse-words-in-a-string-iii.py


class Solution(unittest.TestCase):
    # # Approach: Two Pointers
    # # Time: O(n)
    # # Space: O(n)
    # def reverserWords(self, s: str) -> str:
    #     s = list(s) # [c for c in s]
    #     start = 0
    #     n = len(s)
    #     for end in range(n):
    #         if s[end] == ' ' or end == n-1:
    #             temp = end
    #             end -= int(s[end] == ' ')
    #             while start < end:
    #                 s[start], s[end] = s[end], s[start]
    #                 start += 1
    #                 end -= 1
    #             start = temp+1
    #     return "".join(s)

    # Approach: One-Liner
    # Time: O(n)
    # Space: O(n)
    def reverserWords(self, s: str) -> str:
        return " ".join([w[::-1] for w in s.split(" ")])

    def testReverserWords(self) -> None:
        for s, expected in [
            ("Let's take LeetCode contest", "s'teL ekat edoCteeL tsetnoc"),
            ("Mr Ding", "rM gniD"),
        ]:
            output = self.reverserWords(s)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
