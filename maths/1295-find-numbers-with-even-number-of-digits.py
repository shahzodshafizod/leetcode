import unittest
from typing import List

# https://leetcode.com/problems/find-numbers-with-even-number-of-digits/

# python3 -m unittest maths/1295-find-numbers-with-even-number-of-digits.py

class Solution(unittest.TestCase):
    # # Approach: Extract Digits
    # # Time: O(NxM), N=len(nums), M=max(len(nums[i]))
    # # Space: O(1)
    # def findNumbers(self, nums: List[int]) -> int:
    #     count = 0
    #     for num in nums:
    #         digits = 0
    #         while num > 0:
    #             num //= 10
    #             digits += 1
    #         count += 1 if digits&1==0 else 0
    #     return count

    # Approach: Convert to String
    # Time: O(NxM), N=len(nums), M=max(len(nums[i]))
    # Space: O(M)
    def findNumbers(self, nums: List[int]) -> int:
        return sum(len(str(num))&1==0 for num in nums)

    def testFindNumbers(self) -> None:
        for nums, expected in [
            ([12,345,2,6,7896], 2),
            ([555,901,482,1771], 1),
            ([437,315,322,431,686,264,442], 0),
        ]:
            output = self.findNumbers(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
