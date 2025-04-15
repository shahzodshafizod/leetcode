from typing import List
import unittest

# https://leetcode.com/problems/count-of-smaller-numbers-after-self/

# python3 -m unittest fenwicks/0315-count-of-smaller-numbers-after-self.py

class Solution(unittest.TestCase):
    # # Approach: Merge (Sort) Count
    # # Time: O(nlogn)
    # # Space: O(n)
    # def countSmaller(self, nums: List[int]) -> List[int]:
    #     n = len(nums)
    #     indices = [idx for idx in range(n)]
    #     counts = [0] * n
    #     def merge(left: int, right: int):
    #         if right-left <= 1: return
    #         mid = (left+right)//2
    #         merge(left, mid)
    #         merge(mid, right)
    #         count = 0
    #         idx, left1, left2 = -1, left, mid
    #         merged = [0] * (right-left)
    #         while left1 < mid or left2 < right:
    #             idx += 1
    #             if left1 == mid or left2 < right and nums[indices[left2]] < nums[indices[left1]]:
    #                 merged[idx] = indices[left2]
    #                 left2 += 1
    #                 count += 1
    #             else:
    #                 merged[idx] = indices[left1]
    #                 counts[indices[left1]] += count
    #                 left1 += 1
    #         for idx in range(left, right):
    #             indices[idx] = merged[idx-left]
    #     merge(0, n)
    #     return counts

    # Approach: Binary Indexed Tree (Fenwick Tree)
    # Time: O(nlogn)
    # Space: O(n)
    def countSmaller(self, nums: List[int]) -> List[int]:
        small = min(nums)
        bitree = [0] * (max(nums)-small+2) # 1-indexed
        def update(num: int, delta: int):
            num += 1
            while num < len(bitree):
                bitree[num] += delta
                num += num & -num
        def query(num: int) -> int:
            total = 0
            while num > 0:
                total += bitree[num]
                num -= num & -num
            return total
        n = len(nums)
        for idx in range(n):
            num = nums[idx]-small
            update(num, 1)
        counts = [0] * n
        for idx in range(n):
            num = nums[idx]-small
            update(num, -1)
            counts[idx] = query(num)
        return counts

    def test(self):
        for nums, expected in [
            ([5,2,6,1], [2,1,1,0]),
            ([-1], [0]),
            ([-1,-1], [0,0]),
            ([9,11,15], [0,0,0]),
            ([15,11,9], [2,1,0]),
            ([2,0,1], [2,0,0]),
            # ([26,78,27,100,33,67,90,23,66,5,38,7,35,23,52,22,83,51,98,69,81,32,78,28,94,13,2,97,3,76,99,51,9,21,84,66,65,36,100,41], [10,27,10,35,12,22,28,8,19,2,12,2,9,6,12,5,17,9,19,12,14,6,12,5,12,3,0,10,0,7,8,4,0,0,4,3,2,0,1,0]),
        ]:
            output = self.countSmaller(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
