from typing import List
import unittest

# https://leetcode.com/problems/range-sum-query-immutable/

# python3 -m unittest design/0303-range-sum-query-immutable.py

class NumArray:
    def __init__(self, nums: List[int]):
        n = len(nums)
        self.presum = [0] * (n+1)
        for idx in range(n):
            self.presum[idx+1] += self.presum[idx] + nums[idx]

    def sumRange(self, left: int, right: int) -> int:
        return self.presum[right+1] - self.presum[left]

class TestNumArray(unittest.TestCase):
    def test(self):
        for commands, values, expected in [
            (
                ["NumArray","sumRange","sumRange","sumRange"],
                [[[-2,0,3,-5,2,-1]],[0,2],[2,5],[0,5]],
                [None,1,-1,-3],
            ),
        ]:
            for idx, command in enumerate(commands):
                output = None
                match command:
                    case "NumArray":
                        numArray = NumArray(values[idx][0])
                    case "sumRange":
                        output = numArray.sumRange(values[idx][0], values[idx][1])
                self.assertEqual(expected[idx], output, f"expected: {expected[idx]}, output: {output}")

# Your NumArray object will be instantiated and called as such:
# obj = NumArray(nums)
# param_1 = obj.sumRange(left,right)