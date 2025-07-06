import unittest

# https://leetcode.com/problems/max-difference-you-can-get-from-changing-an-integer/

# python3 -m unittest greedy/1432-max-difference-you-can-get-from-changing-an-integer.py


class Solution(unittest.TestCase):
    # Time: O(log10(num))
    # Space: O(1)
    def maxDiff(self, num: int) -> int:
        num = str(num)
        mn, mx = num, num
        for d in num:
            if d != '9':
                mx = num.replace(d, '9')
                break
        if num[0] != '1':
            mn = num.replace(num[0], '1')
        else:
            for d in num:
                if d != '0' and d != '1':
                    mn = num.replace(d, '0')
                    break
        return int(mx) - int(mn)

    def test(self):
        for num, expected in [
            (555, 888),
            (9, 8),
            (1101057, 8808050),
            (123456, 820000),
            (123, 820),
            (111, 888),
        ]:
            output = self.maxDiff(num)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
