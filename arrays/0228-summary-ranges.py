from typing import List, Tuple
import unittest

# https://leetcode.com/problems/summary-ranges/

# python3 -m unittest arrays/0228-summary-ranges.py


class Solution(unittest.TestCase):
    # Approach: Two pointers for range tracking
    # Track start of each range, extend while consecutive.
    # Time: O(n) - single pass
    # Space: O(1) - excluding result
    def summaryRanges(self, nums: List[int]) -> List[str]:
        # ranges = []
        # for num in nums:
        #     if not ranges or ranges[-1][1]+1 < num:
        #         ranges.append([[] for _ in range(2)])
        #         ranges[-1][0] = num
        #     ranges[-1][1] = num
        # return [str(s) if s==e else f"{s}->{e}" for s, e in ranges]

        result: List[str] = []
        start = 0
        for i in range(len(nums)):
            # Check if end of range (last element or gap detected)
            if i == len(nums) - 1 or nums[i] + 1 != nums[i + 1]:
                if start == i:
                    result.append(str(nums[start]))
                else:
                    result.append(f"{nums[start]}->{nums[i]}")
                start = i + 1
        return result

    def test(self):
        test_cases: List[Tuple[List[int], List[str]]] = [
            ([0, 1, 2, 4, 5, 7], ["0->2", "4->5", "7"]),
            ([0, 2, 3, 4, 6, 8, 9], ["0", "2->4", "6", "8->9"]),
            ([], []),
            ([1], ["1"]),
        ]
        for nums, expected in test_cases:
            output = self.summaryRanges(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
