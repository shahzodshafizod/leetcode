from typing import List
import unittest

# https://leetcode.com/problems/create-maximum-number/

# python3 -m unittest stacks/0321-create-maximum-number.py


class Solution(unittest.TestCase):
    # Approach: Greedy + Monotonic Stack + Merge
    # This problem can be broken down into 3 sub-problems:
    # 1. Find max subsequence of length k from a single array (monotonic stack)
    # 2. Merge two arrays to get lexicographically largest result
    # 3. Try all possible splits: i digits from nums1, k-i from nums2
    #
    # Time: O(k * (m+n)^2) where m=len(nums1), n=len(nums2)
    #   - O(k) iterations for different splits
    #   - O(m+n) for maxSubsequence for each array
    #   - O((m+n)^2) for merge comparison in worst case
    # Space: O(k) for storing subsequences and result
    def maxNumber(self, nums1: List[int], nums2: List[int], k: int) -> List[int]:
        def max_subsequence(nums: List[int], length: int) -> List[int]:
            """
            Get the maximum subsequence of given length from nums
            while maintaining relative order using monotonic stack
            """
            if length == 0:
                return []

            stack: List[int] = []
            to_drop = len(nums) - length  # How many elements we can/must drop

            for num in nums:
                # Pop smaller elements from stack if we still have drops left
                while stack and stack[-1] < num and to_drop > 0:
                    stack.pop()
                    to_drop -= 1
                stack.append(num)

            # Return only the first 'length' elements
            # (we might have more if we didn't drop enough)
            return stack[:length]

        def merge(nums1: List[int], nums2: List[int]) -> List[int]:
            """
            Merge two arrays to get the lexicographically largest result
            """
            result: List[int] = []
            i, j = 0, 0

            while i < len(nums1) and j < len(nums2):
                # Compare remaining subarrays lexicographically
                # Choose the one that gives larger result
                if nums1[i:] > nums2[j:]:
                    result.append(nums1[i])
                    i += 1
                else:
                    result.append(nums2[j])
                    j += 1

            # Append remaining elements
            result.extend(nums1[i:])
            result.extend(nums2[j:])

            return result

        m, n = len(nums1), len(nums2)
        max_result = []

        # Try all possible splits: i from nums1, k-i from nums2
        for i in range(k + 1):
            j = k - i

            # Check if this split is valid
            if i > m or j > n:
                continue

            # Get max subsequence from each array
            sub1 = max_subsequence(nums1, i)
            sub2 = max_subsequence(nums2, j)

            # Merge and compare with current max
            merged = merge(sub1, sub2)
            if merged > max_result:
                max_result = merged

        return max_result

    def test(self):
        for nums1, nums2, k, expected in [
            ([3, 4, 6, 5], [9, 1, 2, 5, 8, 3], 5, [9, 8, 6, 5, 3]),
            ([6, 7], [6, 0, 4], 5, [6, 7, 6, 0, 4]),
            ([3, 9], [8, 9], 3, [9, 8, 9]),
            ([2, 5, 6, 4, 4, 0], [7, 3, 8, 0, 6, 5, 7, 6, 2], 15, [7, 3, 8, 2, 5, 6, 4, 4, 0, 6, 5, 7, 6, 2, 0]),
        ]:
            output = self.maxNumber(nums1, nums2, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
