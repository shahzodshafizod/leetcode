import unittest
from typing import List

# https://leetcode.com/problems/maximum-fruits-harvested-after-at-most-k-steps/

# python3 -m unittest slidingwindows/2106-maximum-fruits-harvested-after-at-most-k-steps.py

class Solutin(unittest.TestCase):
    def maxTotalFruits(self, fruits: List[List[int]], startPos: int, k: int) -> int:
        # to find left, could be used binary search approach
        left, n = 0, len(fruits)
        while left < n and fruits[left][0] < startPos-k:
            left += 1 # skip fruits before the left bound
        total, max_total = 0, 0
        for right in range(left, n):
            if fruits[right][0] > startPos+k:
                break # skip fruits after the right bound
            total += fruits[right][1]
            while min(startPos-2*fruits[left][0] + fruits[right][0], 2*fruits[right][0] - fruits[left][0] - startPos) > k:
                total -= fruits[left][1]
                left += 1
            max_total = max(max_total, total)
        return max_total

    # def maxTotalFruits(self, fruits: List[List[int]], startPos: int, k: int) -> int:
    #     n = len(fruits)
    #     max_pos = startPos+k+1
    #     presum = [0] * (max_pos+1)
    #     idx = 0
    #     for pos in range(1, max_pos+1):
    #         if idx < n and pos-1 == fruits[idx][0]:
    #             presum[pos] = fruits[idx][1]
    #             idx += 1
    #         presum[pos] += presum[pos-1]
    #     harvest = 0
    #     # left-then-right
    #     for left in range(startPos-k, startPos+1):
    #         right = max(startPos, startPos + k - 2*(startPos-left))
    #         harvest = max(harvest, presum[right+1] - presum[(max(0, left))])
    #     # right-then-left
    #     for right in range(startPos+k, startPos-1, -1):
    #         left = min(startPos, startPos - (k - 2*(right-startPos)))
    #         harvest = max(harvest, presum[right+1] - presum[(max(0, left))])
    #     return harvest

    def testMaxTotalFruits(self) -> None:
        for fruits, startPos, k, expected in [
            ([[2,8],[6,3],[8,6]], 5, 4, 9),
            ([[0,9],[4,1],[5,7],[6,2],[7,4],[10,9]], 5, 4, 14),
            ([[0,3],[6,4],[8,5]], 3, 2, 0),
            ([[200000, 10000]], 200000, 0, 10000),
        ]:
            output = self.maxTotalFruits(fruits, startPos, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
