import unittest

# https://leetcode.com/problems/find-subarray-with-bitwise-or-closest-to-k/

# python3 -m unittest bits/3171-find-subarray-with-bitwise-or-closest-to-k.py


class Solution(unittest.TestCase):
    # # Approach #1: Brute-Force
    # # Time: O(nn)
    # # Space: O(1)
    # def minimumDifference(self, nums: list[int], k: int) -> int:
    #     n = len(nums)
    #     min_diff = abs(nums[0] - k)
    #     for i in range(n):
    #         current_or = 0
    #         for j in range(i, n):
    #             current_or |= nums[j]
    #             min_diff = min(min_diff, abs(current_or - k))
    #     return min_diff

    # Approach #2: Sliding OR Window with Early Break
    # Time: O(n * log MAX_VAL) ≈ O(n * 30) = O(n)
    # Space: O(log MAX_VAL) ≈ O(30) = O(1)
    def minimumDifference(self, nums: list[int], k: int) -> int:
        min_diff: int = abs(nums[0] - k)
        ors: list[int] = []
        for num in nums:
            ors.append(0)
            for i in range(len(ors) - 1, -1, -1):
                if ors[i] == ors[i] | num:
                    # ORing num with any previous OR values
                    # (which are subsets of ors[i]) won't change them either
                    break
                ors[i] |= num
                min_diff = min(min_diff, abs(k - ors[i]))
        return min_diff

    def test(self):
        for nums, k, expected in [
            ([1, 2, 4, 5], 3, 0),
            ([1, 3, 1, 3], 2, 1),
            ([1], 10, 9),
            ([1, 2, 1, 2], 4, 1),
            ([3, 6, 9], 5, 1),
        ]:
            output = self.minimumDifference(nums, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
