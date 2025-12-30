from typing import List, Set
import unittest

# https://leetcode.com/problems/remove-invalid-parentheses/

# python3 -m unittest backtracking/0301-remove-invalid-parentheses.py


class Solution(unittest.TestCase):
    # # Approach 1: BFS (Breadth-First Search)
    # # Generate all possible strings by removing one parenthesis at a time
    # # Use BFS to ensure we find the minimum number of removals
    # # Time: O(2^n) where n is the length of string (each char can be kept or removed)
    # # Space: O(2^n) for storing all possible strings in the queue
    # def removeInvalidParentheses(self, s: str) -> List[str]:
    #     def is_valid(string: str) -> bool:
    #         count = 0
    #         for char in string:
    #             if char == "(":
    #                 count += 1
    #             elif char == ")":
    #                 count -= 1
    #                 if count < 0:
    #                     return False
    #         return count == 0

    #     result: List[str] = []
    #     visited: Set[str] = set()
    #     queue = deque([s])
    #     visited.add(s)
    #     found = False

    #     while queue:
    #         current = queue.popleft()

    #         if is_valid(current):
    #             result.append(current)
    #             found = True

    #         # If we found valid strings at this level, don't go deeper
    #         if found:
    #             continue

    #         # Try removing each parenthesis
    #         for i in range(len(current)):
    #             if current[i] not in ("(", ")"):
    #                 continue

    #             next_str = current[:i] + current[i + 1 :]
    #             if next_str not in visited:
    #                 visited.add(next_str)
    #                 queue.append(next_str)

    #     return result if result else [""]

    # Approach 2: Backtracking (Optimized)
    # First calculate how many '(' and ')' need to be removed
    # Then use backtracking to generate all valid combinations
    # Time: O(2^n) where n is the length of string
    # Space: O(n) for recursion depth
    def removeInvalidParentheses(self, s: str) -> List[str]:
        # Step 1: Count how many left and right parentheses to remove
        left_to_remove = 0
        right_to_remove = 0

        for char in s:
            if char == "(":
                left_to_remove += 1
            elif char == ")":
                if left_to_remove > 0:
                    left_to_remove -= 1
                else:
                    right_to_remove += 1

        result: Set[str] = set()

        def backtrack(index: int, left_count: int, right_count: int, left_rem: int, right_rem: int, expr: str):
            # index: current position in string
            # left_count: number of '(' in current expression
            # right_count: number of ')' in current expression
            # left_rem: number of '(' still need to remove
            # right_rem: number of ')' still need to remove
            # expr: current expression being built

            # Base case: reached end of string
            if index == len(s):
                # Valid if we removed all invalid parentheses and counts match
                if left_rem == 0 and right_rem == 0 and left_count == right_count:
                    result.add(expr)
                return

            char = s[index]

            # Option 1: Remove current character (if it's a parenthesis and we need to remove)
            if char == "(" and left_rem > 0:
                backtrack(index + 1, left_count, right_count, left_rem - 1, right_rem, expr)
            elif char == ")" and right_rem > 0:
                backtrack(index + 1, left_count, right_count, left_rem, right_rem - 1, expr)

            # Option 2: Keep current character
            if char == "(":
                backtrack(index + 1, left_count + 1, right_count, left_rem, right_rem, expr + char)
            elif char == ")":
                # Only add ')' if it doesn't make expression invalid
                if right_count < left_count:
                    backtrack(index + 1, left_count, right_count + 1, left_rem, right_rem, expr + char)
            else:
                # Regular character, always keep
                backtrack(index + 1, left_count, right_count, left_rem, right_rem, expr + char)

        backtrack(0, 0, 0, left_to_remove, right_to_remove, "")
        return list(result) if result else [""]

    def test(self):
        for s, expected in [
            ("()())()", ["(())()", "()()()"]),
            ("(a)())()", ["(a())()", "(a)()()"]),
            (")(", [""]),
            ("", [""]),
            ("()", ["()"]),
            ("()())", ["(())", "()()"]),
            ("())", ["()"]),
            ("(()", ["()"]),
        ]:
            output = self.removeInvalidParentheses(s)
            expected.sort()
            output.sort()
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
