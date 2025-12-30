from typing import List
import unittest

# https://leetcode.com/problems/apple-redistribution-into-boxes/

# python3 -m unittest greedy/3074-apple-redistribution-into-boxes.py


class Solution(unittest.TestCase):
    # Approach: Greedy + Sorting
    # Calculate total apples, sort boxes by capacity (largest first)
    # Greedily select boxes until all apples fit
    # N = len(apple), M = len(capacity)
    # Time: O(N + M log M) - sum apples O(N), sort capacities O(M log M), iterate O(M)
    # Space: O(1) - sort in-place (or O(M) if considering sort space)
    def minimumBoxes(self, apple: List[int], capacity: List[int]) -> int:
        # Calculate total apples to distribute
        total_apples = sum(apple)

        # Sort capacities in descending order (largest boxes first)
        capacity.sort(reverse=True)

        # Greedily select boxes
        boxes_needed = 0
        current_capacity = 0

        for box_capacity in capacity:
            current_capacity += box_capacity
            boxes_needed += 1

            # If we can fit all apples, we're done
            if current_capacity >= total_apples:
                return boxes_needed

        # This shouldn't happen given problem constraints
        return boxes_needed

    def test(self):
        for apple, capacity, expected in [
            ([1, 3, 2], [4, 3, 1, 5, 2], 2),
            ([5, 5, 5], [2, 4, 2, 7], 4),
            ([1], [1], 1),
            ([10], [5, 6], 2),
            ([1, 1, 1, 1, 1], [10], 1),
        ]:
            output = self.minimumBoxes(apple, capacity)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
