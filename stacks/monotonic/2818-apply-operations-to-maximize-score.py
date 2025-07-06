from math import sqrt
import heapq
from typing import List
import unittest

# https://leetcode.com/problems/apply-operations-to-maximize-score/

# python3 -m unittest stacks/monotonic/2818-apply-operations-to-maximize-score.py


class Solution(unittest.TestCase):
    # Approach: Monotonic Stack & Priority Queue (Optimized)
    # Time: O(n*sqrt(max) + klogn)
    # Space: O(4n) = O(n)
    def maximumScore(self, nums: List[int], k: int) -> int:
        n = len(nums)
        # step 1: calculate prime score: O(n*sqrt(max))
        scores = [0] * n
        max_heap = [None] * n
        for idx, num in enumerate(nums):
            max_heap[idx] = (-num, idx)
            for f in range(2, int(sqrt(num)) + 1):
                if num % f == 0:
                    scores[idx] += 1
                    while num % f == 0:
                        num //= f
            if num > 1:  # the number itself if prime
                scores[idx] += 1
        # step 2: find the previous and next greater or equal element: O(n)
        left, right = [-1] * n, [n] * n
        stack = []
        for idx in range(n):
            while stack and scores[stack[-1]] < scores[idx]:
                right[stack.pop()] = idx
            if stack:
                left[idx] = stack[-1]
            stack.append(idx)
        # step 3: Greedily picking k contributions: O(klogn)
        score = 1
        MOD = 10**9 + 7
        heapq.heapify(max_heap)
        while k > 0:
            num, idx = heapq.heappop(max_heap)
            num = -num
            count = min((idx - left[idx]) * (right[idx] - idx), k)
            k -= count
            score = (score * pow(num, count, MOD)) % MOD
        return score

    def test(self):
        for nums, k, expected in [
            ([8, 3, 9, 3, 8], 2, 81),
            ([19, 12, 14, 6, 10, 18], 3, 4788),
            ([6, 8, 16, 14, 9], 3, 3136),
            ([1], 1, 1),
            ([1, 7, 11, 1, 5], 14, 823751938),
            ([3289, 2832, 14858, 22011], 6, 256720975),
            ([6, 1, 13, 10, 1, 17, 6], 27, 630596200),
            ([3289, 2832, 14858, 22011], 6, 256720975),
            ([12, 5, 1, 6, 9, 1, 17, 14], 12, 62996359),
        ]:
            output = self.maximumScore(nums, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
