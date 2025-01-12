import unittest

# https://leetcode.com/problems/check-if-a-parentheses-string-can-be-valid/

# python3 -m unittest stacks/2116-check-if-a-parentheses-string-can-be-valid.py

class Solution(unittest.TestCase):
    def canBeValid(self, s: str, locked: str) -> bool:
        n = len(s)
        if n&1 == 1: return False
        # check closing brackets
        locked_cnt, unlocked_cnt = 0, 0
        for idx in range(n):
            if locked[idx] == '0':
                unlocked_cnt += 1
            elif s[idx] == '(':
                locked_cnt += 1
            elif locked_cnt > 0:
                locked_cnt -= 1
            elif unlocked_cnt > 0:
                unlocked_cnt -= 1
            else:
                return False
        # check opening brackets
        locked_cnt, unlocked_cnt = 0, 0
        for idx in range(n-1,-1,-1):
            if locked[idx] == '0':
                unlocked_cnt += 1
            elif s[idx] == ')':
                locked_cnt += 1
            elif locked_cnt > 0:
                locked_cnt -= 1
            elif unlocked_cnt > 0:
                unlocked_cnt -= 1
            else:
                return False
        return True

    def test(self):
        for s, locked, expected in [
            ("))()))", "010100", True),
            ("()()", "0000", True),
            (")", "0", False),
            ("())()))()(()(((())(()()))))((((()())(())", "1011101100010001001011000000110010100101", True),
        ]:
            output = self.canBeValid(s, locked)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
