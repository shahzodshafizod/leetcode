import unittest

# https://leetcode.com/problems/find-the-punishment-number-of-an-integer/

# python3 -m unittest backtracking/2698-find-the-punishment-number-of-an-integer.py

# Number | Squared | Partition | Summation
# 1      | 1       | 1         | 1
# 9      | 81      | 8+1       | 9
# 10     | 100     | 10+0      | 10
# 36     | 1296    | 1+29+6    | 36

# Multiple ways exist to split the digits of a squared integer,
# leading to different summations.


class Solution(unittest.TestCase):
    # Approach: Recursion of Strings
    # Time: O(nn)
    # Space: O(n)
    def punishmentNumber(self, n: int) -> int:
        def partition(target: int, idx: int, s: str):
            if idx == len(s) or target < 0:
                return target == 0
            num = 0
            for c in s[idx:]:
                num = num * 10 + ord(c) - ord('0')
                if partition(target - num, idx + 1, s):
                    return True
                idx += 1
            return False

        panishment = 0
        for num in range(1, n + 1):
            if partition(num, 0, str(num * num)):
                panishment += num * num
        return panishment

    def test(self):
        for n, expected in [
            (10, 182),
            (37, 1478),
        ]:
            output = self.punishmentNumber(n)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
