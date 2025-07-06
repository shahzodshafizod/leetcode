from collections import deque
import heapq  # pylint: disable=unused-import
from typing import List
import unittest

# https://leetcode.com/problems/sliding-window-maximum/

# python3 -m unittest slidingwindows/0239-sliding-window-maximum.py


class Solution(unittest.TestCase):
    # # Approach #1: Brute-Force
    # # Time: O(n*k)
    # # Space: O(1)
    # def maxSlidingWindow(self, nums: List[int], k: int) -> List[int]:
    #     maxes = []
    #     for end in range(k, len(nums)+1):
    #         maximum = int(-1e4)
    #         for idx in range(end-k, end):
    #             maximum = max(maximum, nums[idx])
    #         maxes.append(maximum)
    #     return maxes

    # # Approach #2: Priority Queue
    # # Time: O(N Log N)
    # # Space: O(N)
    # def maxSlidingWindow(self, nums: List[int], k: int) -> List[int]:
    #     max_heap = []
    #     maxes = []
    #     for idx, num in enumerate(nums):
    #         heapq.heappush(max_heap, (-num, idx))
    #         while max_heap and idx-max_heap[0][1] >= k:
    #             heapq.heappop(max_heap)
    #         if idx+1 >= k:
    #             maxes.append(-max_heap[0][0])
    #     return maxes

    # Approach #3: Monotonic Decreasing Queue
    # Time: O(N)
    # Space: O(k)
    def maxSlidingWindow(self, nums: List[int], k: int) -> List[int]:
        queue = deque()
        maxes = []
        for idx, num in enumerate(nums):
            if queue and idx - queue[0] >= k:
                queue.popleft()
            while queue and nums[queue[-1]] < num:
                queue.pop()
            queue.append(idx)
            if idx + 1 >= k:
                maxes.append(nums[queue[0]])
        return maxes

    def test(self):
        for nums, k, expected in [
            ([1, 3, -1, -3, 5, 3, 6, 7], 3, [3, 3, 5, 5, 6, 7]),
            ([1], 1, [1]),
            ([2, 1, 3, 4, 6, 3, 8, 9, 10, 12, 56], 4, [4, 6, 6, 8, 9, 10, 12, 56]),
            ([1, 3, -1, -3, 5, 3, 6, 7], 3, [3, 3, 5, 5, 6, 7]),
            ([1], 1, [1]),
            ([-7, -8, 7, 5, 7, 1, 6, 0], 4, [7, 7, 7, 7, 7]),
            ([9, 10, 9, -7, -4, -8, 2, -6], 5, [10, 10, 9, 2]),  # tricky test case
            ([1, -1], 1, [1, -1]),
            ([7, 2, 4], 2, [7, 4]),
            # ([], 0, []), # invalid test case due to constraints: 1 <= k <= nums.length
        ]:
            output = self.maxSlidingWindow(nums, k)
            self.assertListEqual(expected, output, f"expected: {expected}, output: {output}")
