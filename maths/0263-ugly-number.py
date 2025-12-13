import unittest

# https://leetcode.com/problems/ugly-number/

# python3 -m unittest maths/0263-ugly-number.py


class Solution(unittest.TestCase):
    # Approach: Divide by 2, 3, 5 repeatedly
    # Keep dividing by 2, 3, 5 while divisible.
    # If result is 1, it's an ugly number.
    # Time: O(log n) - divide operations
    # Space: O(1) - constant space
    def isUgly(self, n: int) -> bool:
        # if n <= 0: return False
        # for num in range(7, n+1):
        #     is_prime = True
        #     for i in range(2, ceil(sqrt(num))+1):
        #         if num % i == 0: is_prime = False
        #     if is_prime and n % num == 0:
        #         return False
        # return True

        # if n <= 0: return False
        # while n%2 == 0: n //= 2
        # while n%3 == 0: n //= 3
        # while n%5 == 0: n //= 5
        # return n == 1

        if n <= 0:
            return False

        for factor in [2, 3, 5]:
            while n % factor == 0:
                n //= factor

        return n == 1

    def test(self):
        for n, expected in [
            (6, True),  # 6 = 2 * 3
            (1, True),  # 1 is ugly
            (14, False),  # 14 = 2 * 7
            (0, False),  # non-positive
            (-6, False),  # negative
        ]:
            output = self.isUgly(n)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
