import unittest
from typing import List

# https://leetcode.com/problems/distribute-candies/

# python3 -m unittest hashes/0575-distribute-candies.py


class Solution(unittest.TestCase):
    # # Approach #1: Brute-Force (Count unique types manually)
    # # Time: O(n^2) - checking if each candy type exists in unique list
    # # Space: O(n) - storing unique types
    # def distributeCandies(self, candyType: List[int]) -> int:
    #     # Find all unique candy types manually
    #     unique_types: List[int] = []
    #     for candy in candyType:
    #         is_unique = True
    #         for unique_candy in unique_types:
    #             if candy == unique_candy:
    #                 is_unique = False
    #                 break
    #         if is_unique:
    #             unique_types.append(candy)
    
    #     # Alice can eat at most n/2 candies
    #     max_eat = len(candyType) // 2
    #     # Return minimum of unique types and max she can eat
    #     return min(len(unique_types), max_eat)

    # Approach #2: Hash Set (Optimal)
    # Time: O(n) - single pass to build set
    # Space: O(n) - set stores unique candy types
    def distributeCandies(self, candyType: List[int]) -> int:
        # Count unique candy types using set
        unique_types = len(set(candyType))

        # Alice can eat at most n/2 candies
        max_eat = len(candyType) // 2

        # She wants to maximize variety, so return minimum of:
        # - number of unique types available
        # - maximum number she can eat (n/2)
        return min(unique_types, max_eat)

    def test(self):
        for candyType, expected in [
            ([1, 1, 2, 2, 3, 3], 3),
            ([1, 1, 2, 3], 2),
            ([6, 6, 6, 6], 1),
            ([1, 1, 1, 1, 2, 2, 2, 3, 3, 3], 3),
            ([100000, -100000, 100000, -100000], 2),
        ]:
            output = self.distributeCandies(candyType)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
