from typing import List
import unittest

# https://leetcode.com/problems/final-prices-with-a-special-discount-in-a-shop/

# python3 -m unittest stacks/monotonic/1475-final-prices-with-a-special-discount-in-a-shop.py


class Solution(unittest.TestCase):
    # # Approach #1: Brute-Force
    # # Time: O(nn)
    # # Space: O(1)
    # def finalPrices(self, prices: List[int]) -> List[int]:
    #     n = len(prices)
    #     answer = [None] * n
    #     for idx in range(n):
    #         answer[idx] = prices[idx]
    #         for j in range(idx+1, n):
    #             if prices[j] <= prices[idx]:
    #                 answer[idx] -= prices[j]
    #                 break
    #     return answer

    # # Approach #2: Monotonic Increasing Stack + Binary Search
    # # Time: O(nlogn)
    # # Space: O(n)
    # def finalPrices(self, prices: List[int]) -> List[int]:
    #     stack = []
    #     n = len(prices)
    #     answer = [None] * n
    #     for idx in range(n-1, -1, -1):
    #         answer[idx] = prices[idx]
    #         left, right = 0, len(stack)-1
    #         while left <= right:
    #             mid = (left+right) // 2
    #             if stack[mid] > prices[idx]:
    #                 right = mid-1
    #             else:
    #                 left = mid+1
    #         if right >= 0:
    #             answer[idx] -= stack[right]
    #         while stack and stack[-1] >= prices[idx]:
    #             stack.pop()
    #         stack.append(prices[idx])
    #     return answer

    # Approach #3: Monotonic Increasing Stack
    # Time: O(n)
    # Space: O(n)
    def finalPrices(self, prices: List[int]) -> List[int]:
        stack = []
        n = len(prices)
        answer = [None] * n
        for idx in range(n):
            answer[idx] = prices[idx]
            while stack and answer[stack[-1]] >= prices[idx]:
                answer[stack.pop()] -= prices[idx]
            stack.append(idx)
        return answer

    def test(self):
        for prices, expected in [
            ([10, 1, 1, 6], [9, 0, 1, 6]),
            ([8, 4, 6, 2, 3], [4, 2, 4, 2, 3]),
            ([1, 2, 3, 4, 5], [1, 2, 3, 4, 5]),
            ([10, 2, 5, 2, 8], [8, 0, 3, 2, 8]),
            ([8, 7, 4, 2, 8, 1, 7, 7, 10, 1], [1, 3, 2, 1, 7, 0, 0, 6, 9, 1]),
        ]:
            output = self.finalPrices(prices)
            self.assertListEqual(expected, output, f"expected: {expected}, output: {output}")
