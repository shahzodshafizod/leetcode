from typing import List
import unittest

# https://leetcode.com/problems/baseball-game/

# python3 -m unittest stacks/0682-baseball-game.py


class Solution(unittest.TestCase):
    # # Approach #1: Brute Force - List with operations
    # # Use a list to store all scores, perform operations on the list
    # # For each operation, append/remove from list, then sum at the end
    # # N = number of operations
    # # Time: O(N) - process each operation once
    # # Space: O(N) - list stores all valid scores
    # def calPoints(self, operations: List[str]) -> int:
    #     record: List[int] = []
    
    #     for op in operations:
    #         if op == "+":
    #             # Sum of previous two scores
    #             record.append(record[-1] + record[-2])
    #         elif op == "D":
    #             # Double of previous score
    #             record.append(record[-1] * 2)
    #         elif op == "C":
    #             # Remove previous score
    #             record.pop()
    #         else:
    #             # It's a number
    #             record.append(int(op))
    
    #     return sum(record)

    # Approach #2: Optimized - Stack simulation
    # Use stack for natural push/pop operations
    # Same logic but clearer intent with stack operations
    # N = number of operations
    # Time: O(N) - process each operation once, final sum is O(N)
    # Space: O(N) - stack stores all valid scores
    def calPoints(self, operations: List[str]) -> int:
        stack: List[int] = []

        for op in operations:
            if op == "+":
                # Sum of previous two scores
                stack.append(stack[-1] + stack[-2])
            elif op == "D":
                # Double of previous score
                stack.append(stack[-1] * 2)
            elif op == "C":
                # Remove previous score
                stack.pop()
            else:
                # It's a number (can be negative)
                stack.append(int(op))

        return sum(stack)

    def test(self):
        for operations, expected in [
            (["5", "2", "C", "D", "+"], 30),
            (["5", "-2", "4", "C", "D", "9", "+", "+"], 27),
            (["1", "C"], 0),
            (["1"], 1),
            (["1", "2", "+", "D", "C"], 6),
            (["-60", "D", "-36", "30", "13", "C", "C", "-33", "53", "79"], -117),
        ]:
            output = self.calPoints(operations)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
