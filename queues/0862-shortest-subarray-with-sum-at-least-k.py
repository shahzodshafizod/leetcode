from collections import deque
from typing import List
import unittest

# https://leetcode.com/problems/shortest-subarray-with-sum-at-least-k/

# python3 -m unittest queues/0862-shortest-subarray-with-sum-at-least-k.py


class Solution(unittest.TestCase):
    # # Approach #1: Sliding Window + Prefix Sum (Time Limit Exceeded)
    # # Time: O(nn)
    # # Space: O(1)
    # def shortestSubarray(self, nums: List[int], k: int) -> int:
    #     n = len(nums)
    #     presum = [0] * (n+1)
    #     for idx in range(1, n+1):
    #         presum[idx] = nums[idx-1] + presum[idx-1]
    #     for size in range(1, n+1):
    #         for end in range(size, n+1):
    #             if presum[end] - presum[end-size] >= k:
    #                 return size
    #     return -1

    # # Approach #2: Heap (Priority Queue) + Prefix Sum
    # # Time: O(nlogn)
    # # Space: O(n)
    # def shortestSubarray(self, nums: List[int], k: int) -> int:
    #     length = n+1
    #     presum, n = 0, len(nums)
    #     pq = []
    #     for idx in range(n):
    #         presum += nums[idx]
    #         if presum >= k:
    #             length = min(length, idx+1)
    #         while pq and presum - pq[0][0] >= k:
    #             length = min(length, idx-heapq.heappop(pq)[1])
    #         heapq.heappush(pq, (presum, idx))
    #     return -1 if length > n else length

    # Approach #3: Deque (Monotonic Queue)
    # Time: O(n)
    # Space: O(n)
    def shortestSubarray(self, nums: List[int], k: int) -> int:
        n = len(nums)
        presum = [0] * (n + 1)
        for idx in range(n):
            presum[idx + 1] = presum[idx] + nums[idx]
        queue = deque([0])
        length = n + 1
        for end in range(1, n + 1):
            # Find the minimum valid window ending at 'end'
            while queue and presum[end] - presum[queue[0]] >= k:
                length = min(length, end - queue.popleft())
            # Validate the Monotonic Increasing Queue
            while queue and presum[end] <= presum[queue[-1]]:
                queue.pop()
            queue.append(end)
        return -1 if length > n else length

    def test(self):
        for nums, k, expected in [
            ([1], 1, 1),
            ([1, 2], 4, -1),
            ([2, -1, 2], 3, 3),
            ([48, 99, 37, 4, -31], 140, 2),
            ([44, -25, 75, -50, -38, -42, -32, -6, -40, -46, -47], 19, 1),
            ([10, -5, 200], 150, 1),
            ([2, -1, 1, 3], 4, 2),
        ]:
            output = self.shortestSubarray(nums, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
