from typing import List
import unittest

# https://leetcode.com/problems/minimum-cost-to-make-array-equal/

# python3 -m unittest prefixsums/2448-minimum-cost-to-make-array-equal.py


class Solution(unittest.TestCase):
    # # Approach: Brute-Force (Time Limit Exceeded)
    # # Time: O(N*C), N=len(nums), C=max(cost[i])
    # # Space: O(1)
    # def minCost(self, nums: List[int], cost: List[int]) -> int:
    #     return min(sum(abs(nums[idx]-target)*cost[idx] for idx in range(len(nums))) for target in nums)

    # Approach: Sorting + Prefix Sum
    # Time: O(nlogn + c), n=len(nums), c=max(cost[i])
    # Space: O(n)
    def minCost(self, nums: List[int], cost: List[int]) -> int:
        n = len(nums)
        indices = sorted(range(n), key=lambda idx: nums[idx])
        postcost = sum(cost)
        postmult = sum(n * c for n, c in zip(nums, cost))
        precost, premult = 0, 0
        min_total = float("inf")
        for idx in indices:
            precost += cost[idx]
            postcost -= cost[idx]
            premult += nums[idx] * cost[idx]
            postmult -= nums[idx] * cost[idx]
            total = nums[idx] * precost - premult
            total += postmult - nums[idx] * postcost
            min_total = min(min_total, total)
        return min_total

    def test(self):
        for nums, cost, expected in [
            ([1, 3, 5, 2], [2, 3, 1, 14], 8),
            ([2, 2, 2, 2, 2], [4, 2, 8, 1, 3], 0),
            # ([735103,366367,132236,133334,808160,113001,49051,735598,686615,665317,999793,426087,587000,649989,509946,743518], [724182,447415,723725,902336,600863,287644,13836,665183,448859,917248,397790,898215,790754,320604,468575,825614], 1907611126748), # pylint: disable=line-too-long
        ]:
            output = self.minCost(nums, cost)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
