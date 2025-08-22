from typing import List
import unittest

# https://leetcode.com/problems/count-submatrices-with-all-ones/

# python3 -m unittest stacks/monotonic/1504-count-submatrices-with-all-ones.py


class Solution(unittest.TestCase):
    # # Approach #1: Brute-Force
    # # Time: O(nnxmm)
    # # Space: O(1)
    # def numSubmat(self, mat: List[List[int]]) -> int:
    #     count = 0
    #     m, n = len(mat), len(mat[0])
    #     for row in range(m):
    #         for col in range(n):
    #             bound = n
    #             for i in range(row, m):
    #                 j = col
    #                 while j < bound:
    #                     if mat[i][j]:
    #                         count += 1
    #                     else:
    #                         bound = j
    #                     j += 1
    #     return count

    # # Approach #2: Enumeration
    # # Time: O(mxnn)
    # # Space: O(n)
    # def numSubmat(self, mat: List[List[int]]) -> int:
    #     m, n = len(mat), len(mat[0])
    #     count = 0
    #     height = [0] * n
    #     for row in range(m):
    #         # calculate height for the current row
    #         for col in range(n):
    #             height[col] = height[col] + 1 if mat[row][col] else 0
    #         # for each ceil, count rectangles ending at (row, col)
    #         for col in range(n):
    #             if height[col] > 0:
    #                 min_h = height[col]
    #                 for k in range(col, -1, -1):
    #                     min_h = min(min_h, height[k])
    #                     if min_h == 0:
    #                         break
    #                     count += min_h
    #     return count

    # # Approach #3: Enumeration
    # # Time: O(mmxn)
    # # Space: O(1)
    # def numSubmat(self, mat: List[List[int]]) -> int:
    #     count = 0
    #     m, n = len(mat), len(mat[0])
    #     for row in range(m):
    #         for col in range(n):
    #             if mat[row][col]:
    #                 mat[row][col] += mat[row][col - 1] if col > 0 else 0
    #             cur = mat[row][col]
    #             for k in range(row, -1, -1):
    #                 cur = min(cur, mat[k][col])
    #                 if cur == 0:
    #                     break
    #                 count += cur
    #     return count

    # Approach #4: Monotonic Stack
    # Time: O(mxn)
    # Space: O(n)
    def numSubmat(self, mat: List[List[int]]) -> int:
        m, n = len(mat), len(mat[0])
        count = 0
        height = [0] * n
        for row in range(m):
            sub_res = [0] * n
            stack: List[int] = []
            for col in range(n):
                height[col] = height[col] + 1 if mat[row][col] else 0
                while stack and height[stack[-1]] >= height[col]:
                    stack.pop()
                if stack:
                    prev = stack[-1]
                    sub_res[col] = sub_res[prev] + height[col] * (col - prev)
                else:
                    sub_res[col] = height[col] * (col + 1)
                stack.append(col)
                count += sub_res[col]
        return count

    def test(self):
        for mat, expected in [
            ([[1, 0, 1], [1, 1, 0], [1, 1, 0]], 13),
            ([[0, 1, 1, 0], [0, 1, 1, 1], [1, 1, 1, 0]], 24),
        ]:
            output = self.numSubmat(mat)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
