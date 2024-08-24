import unittest

# https://leetcode.com/problems/find-the-closest-palindrome/

# python3 -m unittest strings/0564-find-the-closest-palindrome.py

class Solution(unittest.TestCase):
    def nearestPalindromic(self, n: str) -> str:
        def makePalindrome(num: int) -> int:
            s = str(num)
            palindrome = list(s)
            n = len(s)
            left, right = (n-1)//2, n//2
            while left >= 0:
                palindrome[right] = palindrome[left]
                left -= 1
                right += 1
            return int("".join(palindrome))

        def prevPalindrome(num: int) -> int:
            left, right = 0, num-1
            prev = 0
            while left <= right:
                mid = (right - left) // 2 + left
                palindrome = makePalindrome(mid)
                if palindrome < num:
                    prev = palindrome
                    left = mid+1
                else:
                    right = mid-1
            return prev

        def nextPalindrome(num: int) -> int:
            left, right = num+1, int(10e17+1)
            next = 0
            while left <= right:
                mid = (right - left) // 2 + left
                palindrome = makePalindrome(mid)
                if palindrome > num:
                    next = palindrome
                    right = mid-1
                else:
                    left = mid+1
            return next

        num = int(n)
        prev = prevPalindrome(num)
        next = nextPalindrome(num)
        if (num - prev) <= (next - num):
            return str(prev)
        return str(next)

    def testNearestPalindromic(self) -> None:
        for n, expected in [
            ("123", "121"),
            ("1", "0"),
            ("139", "141"),
            ("1234", "1221"),
            ("999", "1001"),
            ("1000", "999"),
            ("12932", "12921"),
            ("99800", "99799"),
            ("12120", "12121"),
            ("99899", "99799"),
            ("101", "99"),
            ("11", "9"),
            ("12345", "12321"),
            ("10", "9"),
            ("88", "77"),
            ("1837722381", "1837667381"),
            ("1000000000000000", "999999999999999"),
            ("999999999999999999", "1000000000000000001"),
            ("807045053224792883", "807045053350540708"),
        ]:
            output = self.nearestPalindromic(n)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
