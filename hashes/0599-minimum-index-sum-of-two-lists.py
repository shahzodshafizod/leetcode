import unittest
from typing import List

# https://leetcode.com/problems/minimum-index-sum-of-two-lists/

# python3 -m unittest hashes/0599-minimum-index-sum-of-two-lists.py


class Solution(unittest.TestCase):
    # # Approach #1: Brute-Force (Check all pairs)
    # # Time: O(n * m) - two nested loops
    # # Space: O(1) - no extra space besides output
    # def findRestaurant(self, list1: List[str], list2: List[str]) -> List[str]:
    #     min_sum = len(list1) + len(list2)
    #     result = []
    
    #     # Check all pairs
    #     for i, str1 in enumerate(list1):
    #         for j, str2 in enumerate(list2):
    #             if str1 == str2:
    #                 index_sum = i + j
    #                 if index_sum < min_sum:
    #                     min_sum = index_sum
    #                     result = [str1]
    #                 elif index_sum == min_sum:
    #                     result.append(str1)
    
    #     return result

    # Approach #2: Hash Map (Optimized)
    # Time: O(n + m) - single pass through both lists
    # Space: O(n) - hash map for list1
    def findRestaurant(self, list1: List[str], list2: List[str]) -> List[str]:
        # Build hash map for list1: string -> index
        index_map = {restaurant: i for i, restaurant in enumerate(list1)}

        min_sum = len(list1) + len(list2)
        result = []

        # Iterate through list2 and find common strings
        for j, restaurant in enumerate(list2):
            if restaurant in index_map:
                i = index_map[restaurant]
                index_sum = i + j

                if index_sum < min_sum:
                    # Found new minimum, reset result
                    min_sum = index_sum
                    result = [restaurant]
                elif index_sum == min_sum:
                    # Same minimum, add to result
                    result.append(restaurant)

        return result

    def test(self):
        for list1, list2, expected in [
            (
                ["Shogun", "Tapioca Express", "Burger King", "KFC"],
                ["Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"],
                ["Shogun"],
            ),
            (
                ["Shogun", "Tapioca Express", "Burger King", "KFC"],
                ["KFC", "Shogun", "Burger King"],
                ["Shogun"],
            ),
            (
                ["happy", "sad", "good"],
                ["sad", "happy", "good"],
                ["sad", "happy"],
            ),
        ]:
            output = self.findRestaurant(list1, list2)
            # Sort both for comparison since order doesn't matter
            self.assertEqual(sorted(expected), sorted(output), f"expected: {expected}, output: {output}")
