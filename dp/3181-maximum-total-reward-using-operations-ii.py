from typing import List
import unittest

# https://leetcode.com/problems/maximum-total-reward-using-operations-ii/

# python3 -m unittest dp/3181-maximum-total-reward-using-operations-ii.py


class Solution(unittest.TestCase):
    # # Approach #1: Brute Force Backtracking
    # # Time: O(2^n) - exponential, TLE for large inputs
    # # Space: O(n) - recursion depth
    # def maxTotalReward(self, rewardValues: List[int]) -> int:
    #     rewards = sorted(set(rewardValues))

    #     def backtrack(index: int, current_sum: int) -> int:
    #         if index == len(rewards):
    #             return current_sum

    #         # Skip current reward
    #         max_reward = backtrack(index + 1, current_sum)

    #         # Take current reward if possible
    #         if current_sum < rewards[index]:
    #             max_reward = max(max_reward, backtrack(index + 1, current_sum + rewards[index]))

    #         return max_reward

    #     return backtrack(0, 0)

    # Approach #2: Dynamic Programming with Bitmask Optimization
    # Time: O(n log n + n * max_val) where max_val is maximum reward value
    # Space: O(max_val) for bitmask
    def maxTotalReward(self, rewardValues: List[int]) -> int:
        # Remove duplicates and sort
        rewards = sorted(set(rewardValues))

        # Bitmask: bit i is set if we can achieve sum i
        # Initially, we can achieve sum 0
        possible = 1  # binary: 00...001

        for reward in rewards:
            # We can only use this reward if current sum < reward
            # Get all possible sums less than reward
            # This is done by shifting right to remove sums >= reward
            smaller_sums = possible & ((1 << reward) - 1)

            # Add reward to all smaller sums
            # This creates new possible sums by shifting left by reward positions
            possible |= smaller_sums << reward

        # Find the highest bit set in possible (maximum sum)
        result = 0
        while possible > 0:
            result += 1
            possible >>= 1

        return result - 1  # Subtract 1 because we counted from bit 0

    def test(self):
        for rewardValues, expected in [
            ([1, 1, 3, 3], 4),  # Take 1, then 3: total = 4
            ([1, 6, 4, 3, 2], 11),  # Take 1, 2, 4: total = 7 or 1, 6: total = 7, optimal is 1, 4, 6 = 11
            ([5], 5),  # Take 5 (0 < 5): total = 5
            ([1], 1),  # Take 1
            ([1, 2], 3),  # Take 1, then 2: total = 3
            ([1, 2, 3], 5),  # Take 2, 3: total = 5 (can't take all 3 since 3 not < 3)
        ]:
            output = self.maxTotalReward(rewardValues)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
