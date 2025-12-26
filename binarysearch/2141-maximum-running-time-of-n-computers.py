from typing import List
import unittest

# https://leetcode.com/problems/maximum-running-time-of-n-computers/

# python3 -m unittest binarysearch/2141-maximum-running-time-of-n-computers.py


class Solution(unittest.TestCase):
    # # Approach 1: Brute Force (Would Time Limit Exceed)
    # # Try every possible time from 1 to sum(batteries) and check if we can
    # # run all n computers for that duration.
    # # Time Complexity: O(sum(batteries) * m) where m = len(batteries)
    # # Space Complexity: O(1)
    # def maxRunTime(self, n: int, batteries: List[int]) -> int:
    #     total_sum = sum(batteries)

    #     # Try every possible time from 1 to total_sum // n
    #     max_time = 0
    #     for time in range(1, total_sum // n + 1):
    #         # Check if we can run all n computers for 'time' minutes
    #         available_power = 0
    #         for battery in batteries:
    #             # Each battery can contribute at most 'time' minutes
    #             available_power += min(battery, time)

    #         # If we have enough power to run n computers for 'time' minutes
    #         if available_power >= n * time:
    #             max_time = time

    #     return max_time

    # Approach 2: Binary Search + Greedy (Optimal)
    # The key insight: if we can run all n computers for time t,
    # we can also run them for any time < t (monotonic property).
    # This allows us to use binary search.
    #
    # For a given time t, we can run all n computers if:
    # sum(min(battery[i], t)) >= n * t
    #
    # Explanation: Each battery can contribute at most min(battery[i], t).
    # If total contribution >= n * t, we can distribute power to run
    # all n computers for t minutes (batteries can be swapped).
    #
    # # Time Complexity: O(m * log(sum/n)) where m = len(batteries)
    # # Space Complexity: O(1)
    # def maxRunTime(self, n: int, batteries: List[int]) -> int:
    #     total_sum = sum(batteries)

    #     # Binary search on the answer
    #     # Lower bound: 1 minute
    #     # Upper bound: total sum divided by n computers
    #     left, right = 1, total_sum // n

    #     while left < right:
    #         # Use right-biased mid to find maximum
    #         mid = right - (right - left) // 2

    #         # Check if we can run all n computers for 'mid' minutes
    #         available_power = sum(min(battery, mid) for battery in batteries)

    #         # If we have enough power, try for longer time
    #         if available_power >= n * mid:
    #             left = mid
    #         else:
    #             right = mid - 1

    #     return left

    # Alternative implementation with clearer binary search template
    def maxRunTime(self, n: int, batteries: List[int]) -> int:
        total_sum = sum(batteries)
        left, right = 1, total_sum // n
        result = 0

        while left <= right:
            mid = left + (right - left) // 2

            # Calculate total available power for 'mid' minutes
            power_sum = sum(min(battery, mid) for battery in batteries)

            if power_sum >= n * mid:
                result = mid
                left = mid + 1  # Try for longer time
            else:
                right = mid - 1  # Reduce time

        return result

    def test(self):
        for n, batteries, expected in [
            (2, [3, 3, 3], 4),
            (2, [1, 1, 1, 1], 2),
            (1, [1, 1, 1], 3),
            (3, [10, 10, 3, 5], 8),
            (2, [1, 2, 3, 4, 5], 7),
        ]:
            output = self.maxRunTime(n, batteries)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
