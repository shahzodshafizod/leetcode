from collections import defaultdict
import heapq
from typing import List
import unittest

# https://leetcode.com/problems/sliding-window-median/

# python3 -m unittest slidingwindows/0480-sliding-window-median.py


class Solution(unittest.TestCase):
    # # Approach: Brute-Force
    # def medianSlidingWindow(self, nums: List[int], k: int) -> List[float]:
    #     n = len(nums)
    #     median = [None] * (n-k+1)
    #     for end in range(k, n+1):
    #         # DANGEROUS ARES##################
    #         window = nums[end-k:end].copy()
    #         window.sort()
    #         # #################################
    #         median[end-k] = (window[k//2-1+k&1]+window[k//2])/2
    #     return median

    # Approach: Two Heaps (Two Priority Queues)
    # Time: O(n log k)
    # Space: O(2k) = O(k)
    def medianSlidingWindow(self, nums: List[int], k: int) -> List[float]:
        max_heap = []  # all numbers less than or equal to median
        min_heap = []  # all numbers greater than median
        for idx in range(k):
            heapq.heappush(max_heap, -nums[idx])
            heapq.heappush(min_heap, -heapq.heappop(max_heap))
            if len(min_heap) > len(max_heap):
                heapq.heappush(max_heap, -heapq.heappop(min_heap))

        medians = []
        if k & 1 == 1:
            medians.append(-max_heap[0])
        else:
            medians.append((min_heap[0] - max_heap[0]) / 2)

        heap_dict = defaultdict(int)
        for idx in range(k, len(nums)):
            outbound = nums[idx - k]
            heap_dict[outbound] += 1

            # "-1" means that the outbound element is in max_heap,
            # if we add new element to max_heap, heaps remain balanced
            # "1" means that the outbout element is in min_heap,
            # if we add new element to min_heap, heaps remain balanced
            balance = -1 if outbound <= medians[-1] else 1

            if nums[idx] <= medians[-1]:
                balance += 1
                heapq.heappush(max_heap, -nums[idx])
            else:
                balance -= 1
                heapq.heappush(min_heap, nums[idx])

            if balance < 0:
                # the outbound element was in max_heap,
                # but we added new element to min_heap
                heapq.heappush(max_heap, -heapq.heappop(min_heap))
            elif balance > 0:
                # the outbound element was in min_heap,
                # but we added new element to max_heap
                heapq.heappush(min_heap, -heapq.heappop(max_heap))

            while max_heap and heap_dict[-max_heap[0]] > 0:
                heap_dict[-max_heap[0]] -= 1
                heapq.heappop(max_heap)

            while min_heap and heap_dict[min_heap[0]] > 0:
                heap_dict[min_heap[0]] -= 1
                heapq.heappop(min_heap)

            if k & 1 == 1:
                medians.append(-max_heap[0])
            else:
                medians.append((min_heap[0] - max_heap[0]) / 2)

        return medians

    def test(self):
        for nums, k, expected in [
            ([1, 3, -1, -3, 5, 3, 6, 7], 3, [1.00000, -1.00000, -1.00000, 3.00000, 5.00000, 6.00000]),
            ([1, 2, 3, 4, 2, 3, 1, 4, 2], 3, [2.00000, 3.00000, 3.00000, 3.00000, 2.00000, 3.00000, 2.00000]),
        ]:
            output = self.medianSlidingWindow(nums, k)
            self.assertListEqual(expected, output, f"expected: {expected}, output: {output}")
