from typing import List
import unittest

# https://leetcode.com/problems/teemo-attacking/

# python3 -m unittest arrays/0495-teemo-attacking.py


class Solution(unittest.TestCase):
    # # Approach 1: Brute Force - Mark All Poisoned Seconds
    # # Create a set or boolean array to track poisoned seconds.
    # # For each attack at time t, mark seconds [t, t+duration-1] as poisoned.
    # # Time Complexity: O(n * duration) - could be very large
    # # Space Complexity: O(max_time) - need to store all poisoned seconds
    # # This would TLE for large inputs.
    # def findPoisonedDuration(self, timeSeries: List[int], duration: int) -> int:
    #     if not timeSeries:
    #         return 0

    #     # Use a set to track all poisoned seconds
    #     poisoned: set[int] = set()

    #     for attack_time in timeSeries:
    #         for second in range(attack_time, attack_time + duration):
    #             poisoned.add(second)

    #     return len(poisoned)

    # # Approach 2: Optimal - Calculate Overlapping Intervals
    # # Key insight: For each attack, we add either:
    # # - The full duration (if no overlap with next attack)
    # # - The gap until next attack (if there's overlap)
    # #
    # # Time Complexity: O(n) where n = len(timeSeries)
    # # Space Complexity: O(1)
    # def findPoisonedDuration(self, timeSeries: List[int], duration: int) -> int:
    #     if not timeSeries:
    #         return 0

    #     total_time = 0

    #     # Process each attack except the last
    #     for i in range(len(timeSeries) - 1):
    #         # Calculate gap between current and next attack
    #         gap = timeSeries[i + 1] - timeSeries[i]

    #         # Add either full duration or gap (whichever is smaller)
    #         total_time += min(gap, duration)

    #     # The last attack always lasts full duration
    #     total_time += duration

    #     return total_time

    # Alternative implementation tracking poison end time
    def findPoisonedDuration(self, timeSeries: List[int], duration: int) -> int:
        if not timeSeries:
            return 0

        total_time = 0
        poison_end = 0  # Track when poison effect ends

        for attack_time in timeSeries:
            if attack_time >= poison_end:
                # No overlap - add full duration
                total_time += duration
            else:
                # Overlap - only add additional time
                total_time += attack_time + duration - poison_end

            # Update poison end time
            poison_end = attack_time + duration

        return total_time

    def test(self):
        for timeSeries, duration, expected in [
            ([1, 4], 2, 4),
            ([1, 2], 2, 3),
            ([1, 2, 3, 4, 5], 5, 9),
            ([1], 10, 10),
            ([1, 2, 3, 4, 5, 6, 7, 8, 9], 1, 9),
        ]:
            output = self.findPoisonedDuration(timeSeries, duration)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
