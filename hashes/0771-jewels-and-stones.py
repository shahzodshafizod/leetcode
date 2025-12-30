import unittest

# https://leetcode.com/problems/jewels-and-stones/

# python3 -m unittest hashes/0771-jewels-and-stones.py


class Solution(unittest.TestCase):
    # # Approach 1: Brute Force
    # # For each stone, check if it exists in jewels string
    # # Time: O(n * m) where n is length of stones, m is length of jewels
    # # Space: O(1) - no extra data structures
    # def numJewelsInStones(self, jewels: str, stones: str) -> int:
    #     count = 0
    #     # For each stone, check if it's a jewel
    #     for stone in stones:
    #         # Linear search through jewels
    #         if stone in jewels:
    #             count += 1
    #     return count

    # # Approach 2: Optimized using Set (Hash Table)
    # # Convert jewels to a set for O(1) lookup time
    # # Time: O(n + m) where n is length of jewels, m is length of stones
    # # Space: O(n) where n is length of jewels (for the set)
    # def numJewelsInStones(self, jewels: str, stones: str) -> int:
    #     # Create a set of jewel types for O(1) lookup
    #     jewel_set = set(jewels)

    #     # Count how many stones are jewels
    #     count = 0
    #     for stone in stones:
    #         if stone in jewel_set:
    #             count += 1

    #     return count

    # Alternative one-liner using sum and generator expression
    def numJewelsInStones(self, jewels: str, stones: str) -> int:
        jewel_set = set(jewels)
        return sum(stone in jewel_set for stone in stones)

    def test(self):
        for jewels, stones, expected in [
            ("aA", "aAAbbbb", 3),           # Example 1: 'a', 'A', 'A' are jewels
            ("z", "ZZ", 0),                 # Example 2: case sensitive, no matches
            ("abc", "aabbccdd", 6),         # All a, b, c are jewels
            ("A", "AaAaAa", 3),             # Only uppercase A is jewel
            ("", "stones", 0),              # No jewels
            ("jewels", "", 0),              # No stones
            ("aAbBcC", "aAaAbBbBcCcC", 12), # Multiple jewel types
            ("x", "xxxxXXXX", 4),           # Case sensitive 'x' only
        ]:
            output = self.numJewelsInStones(jewels, stones)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
