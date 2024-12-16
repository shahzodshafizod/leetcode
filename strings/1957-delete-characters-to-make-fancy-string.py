import unittest

# https://leetcode.com/problems/delete-characters-to-make-fancy-string/

# python3 -m unittest strings/1957-delete-characters-to-make-fancy-string.py

class Solution(unittest.TestCase):
    # # Approach: Two Pointers
    # # Time: O(n)
    # # Space: O(n)
    # def makeFancyString(self, s: str) -> str:
    #     fancy = list(s)
    #     slow = 2
    #     for fast in range(2, len(s)):
    #         if s[fast] != fancy[slow-1] or s[fast] != fancy[slow-2]:
    #             fancy[slow] = s[fast]
    #             slow += 1
    #     return "".join(fancy[:min(len(s), slow)])

    # Approach: Stack
    # Time: O(n)
    # Space: O(n)
    def makeFancyString(self, s: str) -> str:
        stack = []
        for c in s:
            if len(stack) < 2 or stack[-1] != c or stack[-2] != c:
                stack.append(c)
        return "".join(stack)

    def test(self):
        for s, expected in [
            ("a", "a"),
            ("aab", "aab"),
            ("aaabaaaa", "aabaa"),
            ("leeetcode", "leetcode"),
            ("leeetcodeeeeeeeeeeeeeeeeeeeeeedddddddddddddddddddddd", "leetcodeedd"),
        ]:
            output = self.makeFancyString(s)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
