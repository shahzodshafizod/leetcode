from collections import defaultdict
import random
import unittest
from typing import List, Tuple, Any, Dict, Set

# from collections import defaultdict

# https://leetcode.com/problems/insert-delete-getrandom-o1-duplicates-allowed/

# python3 -m unittest design/0381-insert-delete-getrandom-o1-duplicates-allowed.py


# # Approach #1: Brute Force (List only)
# # Time: insert O(1), remove O(n), getRandom O(1)
# # Space: O(n)
# class RandomizedCollection:
#     def __init__(self):
#         self.nums: List[int] = []

#     def insert(self, val: int) -> bool:
#         exists = val in self.nums
#         self.nums.append(val)
#         return not exists

#     def remove(self, val: int) -> bool:
#         if val not in self.nums:
#             return False
#         self.nums.remove(val)  # O(n) - finds and removes first occurrence
#         return True

#     def getRandom(self) -> int:
#         return random.choice(self.nums)


# Approach #2: ArrayList + HashMap with Set of Indices
# Time: insert O(1), remove O(1), getRandom O(1)
# Space: O(n)
class RandomizedCollection:
    def __init__(self):
        self.nums: List[int] = []  # ArrayList to store values
        self.indices: Dict[int, Set[int]] = defaultdict(set)  # value -> set of indices

    def insert(self, val: int) -> bool:
        """
        Insert val into the collection.
        Returns True if val was not present, False otherwise.
        """
        exists = val in self.indices
        self.indices[val].add(len(self.nums))
        self.nums.append(val)
        return not exists

    def remove(self, val: int) -> bool:
        """
        Remove one occurrence of val from the collection.
        Returns True if val was present, False otherwise.
        """
        if val not in self.indices or len(self.indices[val]) == 0:
            return False

        # Get any index of val (we'll use the first one from the set)
        removeIdx = self.indices[val].pop()
        lastIdx = len(self.nums) - 1
        lastVal = self.nums[lastIdx]

        # Move last element to the removed position
        self.nums[removeIdx] = lastVal

        # Update indices for the last element (only if removeIdx != lastIdx)
        self.indices[lastVal].add(removeIdx)
        self.indices[lastVal].discard(lastIdx)

        # Remove last element
        self.nums.pop()

        # Clean up empty sets
        if len(self.indices[val]) == 0:
            del self.indices[val]

        return True

    def getRandom(self) -> int:
        """
        Get a random element from the collection.
        Probability is proportional to the number of occurrences.
        """
        return random.choice(self.nums)


class Solution(unittest.TestCase):
    def test(self):
        test_cases: List[Tuple[List[str], List[List[int]], List[Any]]] = [
            (
                ["RandomizedCollection", "insert", "insert", "insert", "getRandom", "remove", "getRandom"],
                [[], [1], [1], [2], [], [1], []],
                [None, True, False, True, None, True, None],
            ),
            (
                ["RandomizedCollection", "insert", "remove", "insert", "remove", "getRandom", "getRandom", "getRandom", "getRandom"],
                [[], [1], [2], [2], [1], [], [], [], []],
                [None, True, False, True, True, None, None, None, None],
            ),
            (
                ["RandomizedCollection", "insert", "insert", "insert", "insert", "remove", "remove", "remove", "remove"],
                [[], [10], [10], [20], [20], [10], [10], [20], [20]],
                [None, True, False, True, False, True, True, True, True],
            ),
            (
                ["RandomizedCollection", "insert", "insert", "remove", "insert"],
                [[], [0], [0], [0], [0]],
                [None, True, False, True, False],
            ),
        ]
        for commands, values, expected in test_cases:
            randomCol: RandomizedCollection | None = None
            for idx, command in enumerate(commands):
                output = None
                match command:
                    case "RandomizedCollection":
                        randomCol = RandomizedCollection()
                    case "insert":
                        assert randomCol is not None
                        output = randomCol.insert(values[idx][0])
                    case "remove":
                        assert randomCol is not None
                        output = randomCol.remove(values[idx][0])
                    case "getRandom":
                        assert randomCol is not None
                        output = randomCol.getRandom()
                        # For getRandom, we just verify it returns a value
                        # We can't check exact value due to randomness
                        continue
                    case _:
                        pass
                if expected[idx] is not None:
                    self.assertEqual(expected[idx], output, f"expected: {expected[idx]}, output: {output}")


# Your RandomizedCollection object will be instantiated and called as such:
# obj = RandomizedCollection()
# param_1 = obj.insert(val)
# param_2 = obj.remove(val)
# param_3 = obj.getRandom()
