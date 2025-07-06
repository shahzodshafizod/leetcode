from collections import defaultdict
from typing import List
import unittest

# https://leetcode.com/problems/number-of-submatrices-that-sum-to-target/

# python3 -m unittest prefixsums/1074-number-of-submatrices-that-sum-to-target.py


class Solution(unittest.TestCase):
    # # Approach: Prefix Sum
    # # Time: O(nnmm)
    # # Space: O(m)
    # def numSubmatrixSumTarget(self, matrix: List[List[int]], target: int) -> int:
    #     m, n = len(matrix), len(matrix[0])
    #     # 560. Subarray Sum Equals K
    #     # from https://leetcode.com/problems/subarray-sum-equals-k/
    #     def subarraySum(nums: List[int], k: int) -> int:
    #         count, presum = 0, 0
    #         for idx in range(m):
    #             presum += nums[idx]
    #             subsum = presum
    #             for start in range(idx+1):
    #                 if subsum == k: count += 1
    #                 subsum -= nums[start]
    #         return count
    #     count = 0
    #     for start in range(n):
    #         presum = [0] * m
    #         for end in range(start, n):
    #             for row in range(m):
    #                 presum[row] += matrix[row][end]
    #             count += subarraySum(presum, target)
    #     return count

    # Approach: Prefix Sum + Hash Table
    # Time: O(nnm)
    # Space: O(m)
    def numSubmatrixSumTarget(self, matrix: List[List[int]], target: int) -> int:
        m, n = len(matrix), len(matrix[0])

        # 560. Subarray Sum Equals K
        # https://leetcode.com/problems/subarray-sum-equals-k/
        def subarraySum(nums: List[int], k: int) -> int:
            counter = defaultdict(int)
            counter[0] = 1
            count, presum = 0, 0
            for idx in range(m):
                presum += nums[idx]
                count += counter[presum - k]
                counter[presum] += 1
            return count

        count = 0
        for start in range(n):
            presum = [0] * m
            for end in range(start, n):
                for row in range(m):
                    presum[row] += matrix[row][end]
                count += subarraySum(presum, target)
        return count

    def test(self):
        for matrix, target, expected in [
            ([[0, 1, 0], [1, 1, 1], [0, 1, 0]], 0, 4),
            ([[1, -1], [-1, 1]], 0, 5),
            ([[904]], 0, 0),
        ]:
            output = self.numSubmatrixSumTarget(matrix, target)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
