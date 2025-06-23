import unittest

# https://leetcode.com/problems/sum-of-k-mirror-numbers/

# python3 -m unittest maths/2081-sum-of-k-mirror-numbers.py

class Solution(unittest.TestCase):
    # Time: O(sqrt(10))
    # Space: O(1)
    def kMirror(self, k: int, n: int) -> int:
        def base_k_mirror(num: int) -> bool:
            digits = []
            while num > 0:
                digits.append(num%k)
                num //= k
            return digits == digits[::-1]
        sum, left = 0, 1
        while n:
            right = left * 10
            for p in [1,0]: # first odd-length, then even-length
                for i in range(left, right):
                    if n == 0: break
                    copy = str(i)
                    if p == 0: copy = copy + copy[::-1]
                    else: copy = copy[:-1] + copy[::-1]
                    copy = int(copy)
                    if base_k_mirror(copy):
                        sum += copy
                        n -= 1
            left = right
        return sum

    def test(self):
        for k, n, expected in [
            (2, 5, 25),
            (3, 7, 499),
            (7, 17, 20379000),
            (4, 5, 66),
        ]:
            output = self.kMirror(k, n)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
