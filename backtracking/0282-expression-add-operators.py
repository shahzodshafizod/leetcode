from typing import List, Tuple
import unittest

# https://leetcode.com/problems/expression-add-operators/

# python3 -m unittest backtracking/0282-expression-add-operators.py


class Solution(unittest.TestCase):
    # # Approach 1: Brute Force - Generate all possible expressions and evaluate each
    # # We try every possible way to insert operators and multi-digit numbers
    # # Time: O(4^n * n) where n is the length of num
    # #   - At each position, we can choose to add an operator or continue building a number
    # #   - Roughly 4 choices at each step: +num, -num, *num, or extend current number
    # #   - Each expression evaluation takes O(n) time
    # # Space: O(n) for recursion depth and expression string
    # def addOperators(self, num: str, target: int) -> List[str]:
    #     result: List[str] = []
    #     n = len(num)

    #     def evaluate(expr: str) -> int:
    #         # Simple evaluation by processing left to right with operator precedence
    #         # We need to handle * having higher precedence than + and -
    #         tokens: List[str] = []
    #         current_num = ""

    #         i = 0
    #         while i < len(expr):
    #             if expr[i].isdigit() or (expr[i] == "-" and i == 0):
    #                 current_num += expr[i]
    #             else:
    #                 tokens.append(current_num)
    #                 tokens.append(expr[i])
    #                 current_num = ""
    #             i += 1
    #         if current_num:
    #             tokens.append(current_num)

    #         # First pass: handle multiplication
    #         i = 0
    #         while i < len(tokens):
    #             if i > 0 and tokens[i] == "*":
    #                 left = int(tokens[i - 1])
    #                 right = int(tokens[i + 1])
    #                 tokens = tokens[: i - 1] + [str(left * right)] + tokens[i + 2 :]
    #             else:
    #                 i += 1

    #         # Second pass: handle addition and subtraction
    #         result = int(tokens[0])
    #         i = 1
    #         while i < len(tokens):
    #             op = tokens[i]
    #             num = int(tokens[i + 1])
    #             if op == "+":
    #                 result += num
    #             elif op == "-":
    #                 result -= num
    #             i += 2

    #         return result

    #     def backtrack(index: int, expr: str):
    #         if index == n:
    #             try:
    #                 if evaluate(expr) == target:
    #                     result.append(expr)
    #             except:
    #                 pass
    #             return

    #         # Try all possible lengths of the next number (1 to n-index digits)
    #         for length in range(1, n - index + 1):
    #             num_str = num[index : index + length]

    #             # Skip numbers with leading zeros (except "0" itself)
    #             if len(num_str) > 1 and num_str[0] == "0":
    #                 break

    #             if index == 0:
    #                 # First number, no operator needed
    #                 backtrack(index + length, num_str)
    #             else:
    #                 # Try each operator
    #                 backtrack(index + length, expr + "+" + num_str)
    #                 backtrack(index + length, expr + "-" + num_str)
    #                 backtrack(index + length, expr + "*" + num_str)

    #     backtrack(0, "")
    #     return result

    # Approach 2: Optimized Backtracking with evaluation during construction
    # Instead of evaluating at the end, we maintain the current value and last operand
    # This allows us to handle multiplication correctly by reversing the effect of the last operation
    # Time: O(4^n) where n is the length of num
    #   - At each position, we try 4 possibilities (extend number or add +, -, *)
    #   - No need for separate evaluation step
    # Space: O(n) for recursion depth and expression string
    def addOperators(self, num: str, target: int) -> List[str]:
        result: List[str] = []
        n = len(num)

        def backtrack(index: int, expr: str, curr_val: int, last_operand: int):
            # curr_val: current evaluation result
            # last_operand: the value of the last number we added/subtracted
            #               (needed to handle multiplication correctly)

            if index == n:
                if curr_val == target:
                    result.append(expr)
                return

            # Try all possible lengths of the next number (1 to n-index digits)
            for length in range(1, n - index + 1):
                num_str = num[index : index + length]

                # Skip numbers with leading zeros (except "0" itself)
                if len(num_str) > 1 and num_str[0] == "0":
                    break

                current_num = int(num_str)

                if index == 0:
                    # First number, no operator needed
                    backtrack(index + length, num_str, current_num, current_num)
                else:
                    # Addition: simply add to current value
                    backtrack(
                        index + length,
                        expr + "+" + num_str,
                        curr_val + current_num,
                        current_num,
                    )

                    # Subtraction: subtract from current value
                    backtrack(
                        index + length,
                        expr + "-" + num_str,
                        curr_val - current_num,
                        -current_num,
                    )

                    # Multiplication: reverse the last operation and apply multiplication
                    # Example: if we had "1+2" (curr_val=3, last_operand=2) and multiply by 3:
                    # We need: 1+2*3 = 1+6 = 7
                    # So: curr_val - last_operand + (last_operand * current_num)
                    #     = 3 - 2 + (2 * 3) = 1 + 6 = 7
                    backtrack(
                        index + length,
                        expr + "*" + num_str,
                        curr_val - last_operand + last_operand * current_num,
                        last_operand * current_num,
                    )

        backtrack(0, "", 0, 0)
        return result

    def test(self):
        test_cases: List[Tuple[str, int, List[str]]] = [
            ("123", 6, ["1*2*3", "1+2+3"]),
            ("232", 8, ["2*3+2", "2+3*2"]),
            ("3456237490", 9191, []),
            ("105", 5, ["1*0+5", "10-5"]),
            ("00", 0, ["0*0", "0+0", "0-0"]),
        ]
        for num, target, expected in test_cases:
            output = self.addOperators(num, target)
            expected.sort()
            output.sort()
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
