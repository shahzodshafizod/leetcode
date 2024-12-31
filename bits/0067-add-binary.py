import unittest

# https://leetcode.com/problems/add-binary/

# python3 -m unittest bits/0067-add-binary.py

class Solution(unittest.TestCase):
    def addBinary(self, a: str, b: str) -> str:
        bisum = ""
        aid = len(a)-1
        bid = len(b)-1
        carry = 0
        while aid >= 0 or bid >= 0:
            if aid >= 0:
                carry += ord(a[aid]) - ord('0')
                aid -= 1
            if bid >= 0:
                carry += ord(b[bid]) - ord('0')
                bid -= 1
            bisum += str(carry&1)
            carry >>= 1
        if carry != 0:
            bisum += str(carry)
        return bisum[::-1]

    def test(self):
        for a, b, expected in [
            ("11", "1", "100"),
            ("1010", "1011", "10101"),
        ]:
            output = self.addBinary(a, b)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
