import unittest

# https://leetcode.com/problems/minimum-operations-to-make-array-equal-to-target/

# python3 -m unittest greedy/3229-minimum-operations-to-make-array-equal-to-target.py


class Solution(unittest.TestCase):
    # # Approach #1: Brute-force (Simulation-based)
    # # Time: O(n * max_diff) where max_diff is the maximum difference
    # # Space: O(n) for creating a copy of nums array
    # def minimumOperations(self, nums: list[int], target: list[int]) -> int:
    #     n = len(nums)
    #     operations = 0
    #     while nums != target:
    #         i = 0
    #         while i < n and nums[i] == target[i]:
    #             i += 1
    #         diff = target[i] - nums[i]
    #         increment = 1 if diff > 0 else -1
    #         start, end = i, i
    #         while end < n and (target[end] - nums[end]) * increment > 0:
    #             end += 1
    #         for k in range(start, end):
    #             nums[k] += increment
    #         operations += 1
    #     return operations

    # Approach #2: Greedy
    # Time: O(n) - single pass through the array
    # Space: O(1) - only using constant extra space
    def minimumOperations(self, nums: list[int], target: list[int]) -> int:
        n = len(nums)
        curr_diff = target[0] - nums[0]
        operations = abs(curr_diff)

        for i in range(1, n):
            prev_diff = curr_diff
            curr_diff = target[i] - nums[i]

            if prev_diff * curr_diff >= 0:  # if signs are same
                # Same direction: add only the increase in magnitude
                operations += max(0, abs(curr_diff) - abs(prev_diff))
            else:
                # Different directions: need full magnitude
                operations += abs(curr_diff)

        return operations

    def test(self):
        for nums, target, expected in [
            ([3, 5, 1, 2], [4, 6, 2, 4], 2),
            ([1, 3, 2], [2, 1, 4], 5),
            ([3, 5, 1, 2], [4, 6, 2, 4], 2),
            ([1, 3, 2], [2, 1, 4], 5),
            ([1, 1, 1, 1], [1, 1, 1, 1], 0),
            ([5, 5, 5, 5], [3, 3, 3, 3], 2),
            ([1, 2, 3, 4, 5], [5, 4, 3, 2, 1], 8),
            ([10], [20], 10),
            ([100, 200, 300], [150, 150, 150], 200),
            ([1, 5, 10], [10, 5, 1], 18),
        ]:
            output = self.minimumOperations(nums[:], target[:])
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
