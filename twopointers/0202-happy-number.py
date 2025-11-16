import unittest

# https://leetcode.com/problems/happy-number/

# python3 -m unittest twopointers/0202-happy-number.py


class Solution(unittest.TestCase):
    # Approach: Floyd's Tortoise and Hare (Sangpusht va Xargushi Floyd) for cycle detection
    # Space: O(1)
    def isHappy(self, n: int) -> bool:
        def get_sum(n: int) -> int:
            total = 0
            while n > 0:
                digit = n % 10
                total += digit * digit
                n //= 10
            return total

        slow, fast = n, n
        while True:
            slow = get_sum(slow)
            fast = get_sum(fast)
            fast = get_sum(fast)
            if fast == 1:
                return True
            if fast == slow:
                break
        return False

    def test(self):
        for n, expected in [
            (19, True),
            (2, False),
        ]:
            output = self.isHappy(n)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
