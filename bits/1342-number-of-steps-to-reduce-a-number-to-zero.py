import unittest

# https://leetcode.com/problems/number-of-steps-to-reduce-a-number-to-zero/

# python3 -m unittest bits/1342-number-of-steps-to-reduce-a-number-to-zero.py


class Solution(unittest.TestCase):
    def numberOfSteps(self, num: int) -> int:
        steps = 0
        while num:
            if num & 1:
                num ^= 1
            else:
                num >>= 1
            steps += 1
        return steps
        # if num == 0:
        #     return 0
        # steps = 0
        # while num:
        #     steps += 1 + (num & 1)
        #     num >>= 1
        # return steps - 1

    def test(self):
        for num, expected in [
            (14, 6),
            (8, 4),
            (123, 12),
            (0, 0),
        ]:
            output = self.numberOfSteps(num)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
