from typing import List
import unittest

# https://leetcode.com/problems/binary-watch/

# python3 -m unittest backtracking/0401-binary-watch.py


class Solution(unittest.TestCase):
    # # Approach 1: Brute Force - Try All Possible Times
    # # Iterate through all possible hours (0-11) and minutes (0-59).
    # # For each combination, count the number of bits set.
    # # If it matches turnedOn, add to result.
    # # Time Complexity: O(12 * 60) = O(720) = O(1)
    # # Space Complexity: O(1) - not counting output
    # def readBinaryWatch(self, turnedOn: int) -> List[str]:
    #     result: List[str] = []

    #     # Try all possible hours and minutes
    #     for hour in range(12):
    #         for minute in range(60):
    #             # Count total set bits
    #             if bin(hour).count('1') + bin(minute).count('1') == turnedOn:
    #                 # Format: h:mm
    #                 result.append(f"{hour}:{minute:02d}")

    #     return result

    # Approach 2: Bit Manipulation (Same as brute force, just cleaner)
    # Count set bits in hour and minute using bit_count()
    # Time Complexity: O(1) - fixed iterations
    # Space Complexity: O(1)
    def readBinaryWatch(self, turnedOn: int) -> List[str]:
        result: List[str] = []

        for hour in range(12):
            for minute in range(60):
                # Count set bits using bit_count (Python 3.10+) or bin().count('1')
                if hour.bit_count() + minute.bit_count() == turnedOn:
                    result.append(f"{hour}:{minute:02d}")

        return result

    # # Approach 3: Backtracking
    # # Generate all combinations of LED positions
    # # Time Complexity: O(C(10, turnedOn))
    # # Space Complexity: O(turnedOn) - recursion depth
    # def readBinaryWatch(self, turnedOn: int) -> List[str]:
    #     result: List[str] = []

    #     # LED positions: first 4 are hours [8,4,2,1], last 6 are minutes [32,16,8,4,2,1]
    #     positions = [8, 4, 2, 1, 32, 16, 8, 4, 2, 1]

    #     def backtrack(index: int, count: int, hours: int, minutes: int):
    #         # Pruning
    #         if hours >= 12 or minutes >= 60:
    #             return

    #         # Found valid combination
    #         if count == turnedOn:
    #             result.append(f"{hours}:{minutes:02d}")
    #             return

    #         # Try turning on each remaining LED
    #         for i in range(index, len(positions)):
    #             if i < 4:
    #                 # Hour LED
    #                 backtrack(i + 1, count + 1, hours + positions[i], minutes)
    #             else:
    #                 # Minute LED
    #                 backtrack(i + 1, count + 1, hours, minutes + positions[i])

    #     backtrack(0, 0, 0, 0)
    #     return result

    def test(self):
        test_cases: list[tuple[int, list[str]]] = [
            (1, ["0:01", "0:02", "0:04", "0:08", "0:16", "0:32", "1:00", "2:00", "4:00", "8:00"]),
            (9, []),
            (0, ["0:00"]),
        ]
        for turnedOn, expected in test_cases:
            output = self.readBinaryWatch(turnedOn)
            # Sort both lists to compare (order doesn't matter)
            output.sort()
            expected.sort()
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
