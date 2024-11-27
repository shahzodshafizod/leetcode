from typing import List
import unittest

# https://leetcode.com/problems/plus-one/

# python3 -m unittest maths/0066-plus-one.py

class Solution(unittest.TestCase):
    def plusOne(self, digits: List[int]) -> List[int]:
        carry = 1
        for idx in range(len(digits)-1, -1, -1):
            carry, digits[idx] = divmod(digits[idx] + carry, 10)
            if carry == 0: break
        return digits if not carry else [carry] + digits

    def testPlusOne(self) -> None:
        for digits, expected in [
            ([1,2,3], [1,2,4]),
            ([4,3,2,1], [4,3,2,2]),
            ([9], [1,0]),
            ([9,9,9,9], [1,0,0,0,0]),
            ([9,8,7,6,5,4,3,2,1,0], [9,8,7,6,5,4,3,2,1,1]),
            ([6,1,4,5,3,9,0,1,9,5,1,8,6,7,0,5,5,4,3], [6,1,4,5,3,9,0,1,9,5,1,8,6,7,0,5,5,4,4]),
            ([9,9], [1,0,0]),
        ]:
            output = self.plusOne(digits)
            self.assertListEqual(expected, output, f"expected: {expected}, output: {output}")
