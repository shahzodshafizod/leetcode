import heapq
from typing import List
import unittest

# https://leetcode.com/problems/minimize-deviation-in-array/

# python3 -m unittest heaps/1675-minimize-deviation-in-array.py

class Solution(unittest.TestCase):
    # # Approach: Heap (Priority Queue)
    # # Time: O(n x logm x logn), n=len(nums), m=max(nums[i])
    # # Space: O(n)
    # def minimumDeviation(self, nums: List[int]) -> int:
    #     n = len(nums)
    #     max_heap = [0] * n
    #     largest = 0
    #     for idx in range(n): # O(n)
    #         mn = mx = nums[idx]
    #         while mn&1 == 0: # O(logm)
    #             mn //= 2
    #         if mx&1:
    #             mx *= 2
    #         max_heap[idx] = (mn, mx)
    #         largest = max(largest, mn)
    #     heapq.heapify(max_heap)
    #     deviation = int(1e9)
    #     while len(max_heap) == n: # O(logm)
    #         mn, mx = heapq.heappop(max_heap) # O(nlogn)
    #         deviation = min(deviation, largest-mn)
    #         if mn < mx:
    #             heapq.heappush(max_heap, (mn*2, mx)) # O(nlogn)
    #             largest = max(largest, mn*2)
    #     return deviation

    # Approach: Heap (Priority Queue) (Optimized)
    # Time: O(n + logn x logm), n=len(nums), m=max(nums[i])
    # Space: O(n)
    def minimumDeviation(self, nums: List[int]) -> int:
        n = len(nums)
        max_heap = [0]*n
        smallest = int(1e9)
        for idx, num in enumerate(nums): # O(n)
            if num&1:
                num *= 2
            max_heap[idx] = -num
            smallest = min(smallest, num)
        heapq.heapify(max_heap) # O(n)
        deviation = int(1e9)
        while len(max_heap) == n: # O(logm)
            num = -max_heap[0] # O(logn)
            deviation = min(deviation, num-smallest)
            if num&1 == 0:
                num //= 2
                smallest = min(smallest, num)
                heapq.heapreplace(max_heap, -num) # O(logn)
            else:
                heapq.heappop(max_heap)
        return deviation

    def test(self):
        for nums, expected in [
            ([1,2,3,4], 1),
            ([4,1,5,20,3], 3),
            ([2,10,8], 3),
            ([2,2,3,2], 1),
            ([3,5], 1),
        ]:
            output = self.minimumDeviation(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
