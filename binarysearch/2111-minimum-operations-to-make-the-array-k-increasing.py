from typing import List
import unittest

# https://leetcode.com/problems/minimum-operations-to-make-the-array-k-increasing/

# python3 -m unittest binarysearch/2111-minimum-operations-to-make-the-array-k-increasing.py


class Solution(unittest.TestCase):
    # # Approach: Brute-Force (Dynamic Programming for each subsequence)
    # # Time: O(k * (n/k)^2) = O(n^2 / k)
    # # Space: O(n/k)
    # def kIncreasing(self, arr: List[int], k: int) -> int:
    #     def minOpsForNonDecreasing(subseq: List[int]) -> int:
    #         """Find minimum operations to make subsequence non-decreasing using DP"""
    #         n = len(subseq)
    #         if n <= 1:
    #             return 0
    
    #         # dp[i] = length of longest non-decreasing subsequence ending at index i
    #         dp = [1] * n
    
    #         for i in range(1, n):
    #             for j in range(i):
    #                 if subseq[j] <= subseq[i]:
    #                     dp[i] = max(dp[i], dp[j] + 1)
    
    #         # Maximum length of non-decreasing subsequence
    #         max_len = max(dp)
    #         # Minimum operations = total elements - elements we can keep
    #         return n - max_len
    
    #     total_ops = 0
    #     # Split array into k subsequences based on i % k
    #     for remainder in range(k):
    #         subseq: List[int] = []
    #         idx = remainder
    #         while idx < len(arr):
    #             subseq.append(arr[idx])
    #             idx += k
    #         total_ops += minOpsForNonDecreasing(subseq)
    
    #     return total_ops

    # Approach: Binary Search + Longest Non-Decreasing Subsequence (LIS variant)
    # Time: O(n log(n/k)) - for each of k subsequences of length n/k, we do binary search
    # Space: O(n/k) - for storing the LIS array
    def kIncreasing(self, arr: List[int], k: int) -> int:
        def longestNonDecreasingSubsequence(subseq: List[int]) -> int:
            """Find length of longest non-decreasing subsequence using binary search"""
            if not subseq:
                return 0

            # lis[i] stores the smallest tail element for non-decreasing subseq of length i+1
            lis: List[int] = []

            for num in subseq:
                # Binary search for the position to insert/replace
                # We want rightmost position where lis[pos] <= num
                left, right = 0, len(lis)
                while left < right:
                    mid = (left + right) // 2
                    if lis[mid] > num:
                        right = mid
                    else:
                        left = mid + 1

                # If we're at the end, append
                if right == len(lis):
                    lis.append(num)
                else:
                    # Replace element at position right
                    lis[right] = num

            return len(lis)

        total_ops = 0
        # Split array into k subsequences based on i % k
        for remainder in range(k):
            subseq: List[int] = []
            idx = remainder
            while idx < len(arr):
                subseq.append(arr[idx])
                idx += k

            # Minimum operations = total - longest non-decreasing subsequence
            total_ops += len(subseq) - longestNonDecreasingSubsequence(subseq)

        return total_ops

    def test(self):
        for arr, k, expected in [
            ([5, 4, 3, 2, 1], 1, 4),
            ([4, 1, 5, 2, 6, 2], 2, 0),
            ([4, 1, 5, 2, 6, 2], 3, 2),
            ([12, 6, 12, 6, 14, 2, 13, 17, 3, 8, 11, 7, 4, 11, 18, 8, 8, 3], 1, 12),
        ]:
            output = self.kIncreasing(arr, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
