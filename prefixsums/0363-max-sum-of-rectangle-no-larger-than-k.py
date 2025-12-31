import bisect
import unittest
from typing import List

# https://leetcode.com/problems/max-sum-of-rectangle-no-larger-than-k/

# python3 -m unittest prefixsums/0363-max-sum-of-rectangle-no-larger-than-k.py


class Solution(unittest.TestCase):
    # # Approach #1: Brute Force with Prefix Sum
    # # Time: O(m²n²)
    # # Space: O(mn)
    # def maxSumSubmatrix(self, matrix: List[List[int]], k: int) -> int:
    #     m, n = len(matrix), len(matrix[0])

    #     # Build 2D prefix sum array
    #     # presum[i][j] = sum of rectangle from (0,0) to (i-1,j-1)
    #     presum = [[0] * (n + 1) for _ in range(m + 1)]
    #     for row in range(m):
    #         for col in range(n):
    #             presum[row + 1][col + 1] = matrix[row][col] + presum[row][col + 1] + presum[row + 1][col] - presum[row][col]

    #     maxSum = -(1 << 31)

    #     # Try all possible rectangles
    #     for top in range(m):
    #         for left in range(n):
    #             for bottom in range(top, m):
    #                 for right in range(left, n):
    #                     # Calculate sum of rectangle (top,left) to (bottom,right)
    #                     rectSum = presum[bottom + 1][right + 1] - presum[top][right + 1] - presum[bottom + 1][left] + presum[top][left]
    #                     if rectSum <= k:
    #                         maxSum = max(maxSum, rectSum)

    #     return maxSum

    # Approach #2: Prefix Sum + Binary Search (Ordered Set)
    # Time: O(n²m·log(m)) where n = cols, m = rows
    # Space: O(m)
    def maxSumSubmatrix(self, matrix: List[List[int]], k: int) -> int:
        def maxSumSubarray(nums: List[int], k: int) -> int:
            """
            Find maximum subarray sum <= k using prefix sum + binary search.

            For each position i, we want to find the smallest prefix sum
            that is >= (currentSum - k), so that (currentSum - prefixSum) <= k.

            We maintain a sorted list of all prefix sums seen so far.
            """
            maxSum = -(1 << 31)
            prefixSums = [0]  # Sorted list of prefix sums
            currentSum = 0

            for num in nums:
                currentSum += num

                # Find smallest prefix sum >= (currentSum - k)
                # This gives us the subarray sum = currentSum - prefixSum <= k
                target = currentSum - k
                idx = bisect.bisect_left(prefixSums, target)

                if idx < len(prefixSums):
                    maxSum = max(maxSum, currentSum - prefixSums[idx])

                # Insert current sum in sorted order
                bisect.insort(prefixSums, currentSum)

            return maxSum

        m, n = len(matrix), len(matrix[0])
        maxSum = -(1 << 31)

        # Optimize: iterate over smaller dimension
        # If rows >> cols, we iterate over rows instead
        if m > n:
            # Transpose: iterate over rows as column boundaries
            for left in range(m):
                rowSums = [0] * n
                for right in range(left, m):
                    # Compress columns [left, right] into 1D array
                    for col in range(n):
                        rowSums[col] += matrix[right][col]

                    # Find max subarray sum <= k in 1D array
                    maxSum = max(maxSum, maxSumSubarray(rowSums, k))
                    if maxSum == k:
                        return k
        else:
            # Standard: iterate over columns as boundaries
            for left in range(n):
                rowSums = [0] * m
                for right in range(left, n):
                    # Compress columns [left, right] into 1D array
                    for row in range(m):
                        rowSums[row] += matrix[row][right]

                    # Find max subarray sum <= k in 1D array
                    maxSum = max(maxSum, maxSumSubarray(rowSums, k))
                    if maxSum == k:
                        return k

        return maxSum

    def test(self):
        for matrix, k, expected in [
            ([[1, 0, 1], [0, -2, 3]], 2, 2),
            ([[2, 2, -1]], 3, 3),
            ([[5, -4, -3, 4], [-3, -4, 4, 5], [5, 1, 5, -4]], 8, 8),
            ([[2, 2, -1]], 0, -1),
        ]:
            output = self.maxSumSubmatrix(matrix, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
