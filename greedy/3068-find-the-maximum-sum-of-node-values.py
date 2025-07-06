from typing import List
import unittest

# https://leetcode.com/problems/find-the-maximum-sum-of-node-values/

# python3 -m unittest greedy/3068-find-the-maximum-sum-of-node-values.py


class Solution(unittest.TestCase):
    # # Approach #1: Brute Force
    # # Time: O(???)
    # # Space: O(1)
    # def maximumValueSum(self, nums: List[int], k: int, edges: List[List[int]]) -> int:
    #     exit = False
    #     while not exit:
    #         exit = True
    #         for u,v in edges:
    #             if (nums[u]^k)+(nums[v]^k) > nums[u]+nums[v]:
    #                 exit = False
    #                 nums[u] ^= k
    #                 nums[v] ^= k
    #     return sum(nums)

    # # Approach #2: Top-Down Dynamic Programming (Memoization)
    # # Time: O(n)
    # # Space: O(n)
    # def maximumValueSum(self, nums: List[int], k: int, edges: List[List[int]]) -> int:
    #     n = len(nums)
    #     def dp(idx: int, even: int) -> int:
    #         if idx == n:
    #             return 0 if even == 1 else -float("inf")
    #         no_xor_done = nums[idx] + dp(idx+1, even)
    #         xor_done = (nums[idx] ^ k) + dp(idx+1, even^1)
    #         return max(no_xor_done, xor_done)
    #     return dp(0, 1)

    # # Approach #3: Bottom-Up Dynamic Programming (Tabulation)
    # # Time: O(n)
    # # Space: O(n)
    # def maximumValueSum(self, nums: List[int], k: int, edges: List[List[int]]) -> int:
    #     n = len(nums)
    #     dp = [[0,0] for _ in range(n+1)]
    #     dp[n][0] = -float("inf")
    #     for idx in range(n-1,-1,-1):
    #         for even in [0,1]:
    #             xored = dp[idx+1][even^1] + (nums[idx] ^ k)
    #             not_xored = dp[idx+1][even] + nums[idx]
    #             dp[idx][even] = max(xored, not_xored)
    #     return dp[0][1]

    # # Approach #4: Greedy (Sorting based approach)
    # # Time: O(nlogn)
    # # Space: O(n)
    # def maximumValueSum(self, nums: List[int], k: int, edges: List[List[int]]) -> int:
    #     deltas = sorted([(num^k)-num for num in nums])
    #     res = sum(nums)
    #     while len(deltas) > 1 and deltas[-2]+deltas[-1] > 0:
    #         res += deltas.pop() + deltas.pop()
    #     return res

    # # Approach #5: Greedy (Sorting based approach)
    # # Time: O(nlogn)
    # # Space: O(1)
    # def maximumValueSum(self, nums: List[int], k: int, edges: List[List[int]]) -> int:
    #     n, res = len(nums), 0
    #     for idx in range(n):
    #         res += nums[idx]
    #         nums[idx] = (nums[idx] ^ k) - nums[idx]
    #     nums.sort()
    #     while len(nums) > 1 and nums[-2]+nums[-1] > 0:
    #         res += nums.pop() + nums.pop()
    #     return res

    # Approach #6: Greedy (Finding local maxima and minima)
    # Time: O(n)
    # Space: O(1)
    def maximumValueSum(
        self, nums: List[int], k: int, edges: List[List[int]]  # pylint: disable=unused-argument
    ) -> int:
        # 10**9 = 0b111011100110101100101000000000
        min_delta = 1 << 30
        total, cnt = 0, 0
        for num in nums:
            total += num
            delta = (num ^ k) - num
            if delta > 0:
                total += delta
                cnt += 1
                min_delta = min(min_delta, delta)
            else:
                min_delta = min(min_delta, -delta)
        if cnt & 1:
            total -= min_delta
        return total

    def test(self):
        for nums, k, edges, expected in [
            ([1, 2, 1], 3, [[0, 1], [0, 2]], 6),
            ([2, 3], 7, [[0, 1]], 9),
            ([7, 7, 7, 7, 7, 7], 3, [[0, 1], [0, 2], [0, 3], [0, 4], [0, 5]], 42),
            ([24, 78, 1, 97, 44], 6, [[0, 2], [1, 2], [4, 2], [3, 4]], 260),
        ]:
            output = self.maximumValueSum(nums, k, edges)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
