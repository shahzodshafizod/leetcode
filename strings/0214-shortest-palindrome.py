import unittest

# https://leetcode.com/problems/shortest-palindrome/

# python3 -m unittest strings/0214-shortest-palindrome.py

class Solution(unittest.TestCase):
    # # Approach #1: Brute-Force
    # # Time: O(n^2)
    # # Space: O(1)
    # def shortestPalindrome(self, s: str) -> str:
    #     n = len(s)
    #     def is_palindrome(left: int, right: int) -> tuple[bool, int]:
    #         while left >= 0 and right < n and s[left] == s[right]:
    #             left -= 1
    #             right += 1
    #         return left < 0, right
    #     # 1. find the longest palindrome starting from 0
    #     left = (n-1)//2
    #     length = 0
    #     while left >= 0 and length == 0:
    #         for right in [left+1, left]:
    #             ok, l = is_palindrome(left, right)
    #             if ok:
    #                 length = l
    #                 break
    #         left -= 1
    #     # 2. reverse the remaining and add in the beginning
    #     return s[length:][::-1] + s

    # # Approach #2: KMP (Knuth-Morris-Pratt) Algorithm
    # # Time: O(n)
    # # Space: O(n)
    # def shortestPalindrome(self, s: str) -> str:
    #     rev = s[::-1]
    #     combined = s + '#' + rev
    #     kmp = [0] * len(combined)
    #     length = 0
    #     for idx in range(1, len(combined)):
    #         length = kmp[idx-1]
    #         while length > 0 and combined[idx] != combined[length]:
    #             length = kmp[length-1]
    #         if combined[idx] == combined[length]:
    #             length += 1
    #         kmp[idx] = length
    #     return rev[:len(s)-length] + s

    # Approach #3: Rabin Karp Algorithm
    # Time: O(n)
    # Space: O(1)
    def shortestPalindrome(self, s: str) -> str:
        forward, backward = 0, 0
        base, mod = 29, int(1e9+7)
        power = 1
        length = 0
        for idx, c in enumerate(s):
            digit = ord(c) - ord('a') + 1
            forward = (forward*base + digit) % mod
            backward = (digit*power + backward) % mod
            if forward == backward:
                length = idx
            power = (power * base) % mod
        return s[length+1:][::-1] + s

    def test(self):
        for s, expected in [
            ("aacecaaa", "aaacecaaa"),
            ("abcd", "dcbabcd"),
            ("aabba", "abbaabba"),
            ("abb", "bbabb"),
            ("", ""),
            ("a", "a"),
		    ("zz", "zz"),
            # ("abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz", "zyxwvutsrqponmlkjihgfedcbazyxwvutsrqponmlkjihgfedcbazyxwvutsrqponmlkjihgfedcbazyxwvutsrqponmlkjihgfedcbazyxwvutsrqponmlkjihgfedcbazyxwvutsrqponmlkjihgfedcbazyxwvutsrqponmlkjihgfedcbazyxwvutsrqponmlkjihgfedcbazyxwvutsrqponmlkjihgfedcbazyxwvutsrqponmlkjihgfedcbazyxwvutsrqponmlkjihgfedcbazyxwvutsrqponmlkjihgfedcbazyxwvutsrqponmlkjihgfedcbazyxwvutsrqponmlkjihgfedcbazyxwvutsrqponmlkjihgfedcbazyxwvutsrqponmlkjihgfedcbazyxwvutsrqponmlkjihgfedcbazyxwvutsrqponmlkjihgfedcbazyxwvutsrqponmlkjihgfedcbazyxwvutsrqponmlkjihgfedcbazyxwvutsrqponmlkjihgfedcbazyxwvutsrqponmlkjihgfedcbazyxwvutsrqponmlkjihgfedcbazyxwvutsrqponmlkjihgfedcbazyxwvutsrqponmlkjihgfedcbazyxwvutsrqponmlkjihgfedcbazyxwvutsrqponmlkjihgfedcbazyxwvutsrqponmlkjihgfedcbazyxwvutsrqponmlkjihgfedcbazyxwvutsrqponmlkjihgfedcbazyxwvutsrqponmlkjihgfedcbazyxwvutsrqponmlkjihgfedcbazyxwvutsrqponmlkjihgfedcbazyxwvutsrqponmlkjihgfedcbabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz"),
        ]:
            output = self.shortestPalindrome(s)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
