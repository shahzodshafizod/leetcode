from typing import List
from collections import defaultdict
import unittest
from sortedcontainers import SortedList

# https://leetcode.com/problems/find-x-sum-of-all-k-long-subarrays-ii/

# python3 -m unittest slidingwindows/3321-find-x-sum-of-all-k-long-subarrays-ii.py


class Solution(unittest.TestCase):
    # # Brute-force solution
    # def findXSum(self, nums: List[int], k: int, x: int) -> List[int]:
    #     return [sum(num * cnt for num, cnt in sorted(Counter(nums[i : i + k]).items(), key=lambda item: (item[1], item[0]), reverse=True)[:x]) for i in range(len(nums) - k + 1)]

    # Approach #2: Two Heaps
    # Time: O(n log k)
    # Space: O(k)
    def findXSum(self, nums: List[int], k: int, x: int) -> List[int]:
        fm: dict[int, int] = defaultdict(int)
        top: SortedList[tuple[int, int]] = SortedList()
        low: SortedList[tuple[int, int]] = SortedList()
        curr = 0

        def change(num: int, count: int):
            nonlocal curr

            # remove previous state if present
            prev = (fm[num], num)
            if num in fm:  # means exists in top or low
                if prev in low:
                    low.discard(prev)
                else:
                    top.discard(prev)
                    curr -= fm[num] * num

            # add new state
            fm[num] += count
            if fm[num]:
                low.add((fm[num], num))

            # rebalance: if top is not full
            while low and len(top) < x:
                freq, num = low.pop(-1)
                curr += freq * num
                top.add((freq, num))

            # rebalance: swap, if 'low' has greater freqs than some 'top' elements
            while low and low[-1] > top[0]:
                lfreq, lnum = low.pop(-1)
                tfreq, tnum = top.pop(0)
                curr = curr - tfreq * tnum + lfreq * lnum
                low.add((tfreq, tnum))
                top.add((lfreq, lnum))

        n = len(nums)
        answer: List[int] = []
        for i in range(n):
            change(nums[i], 1)
            if i >= k - 1:
                answer.append(curr)
                change(nums[i - k + 1], -1)

        return answer

    def test(self):
        for nums, k, x, expected in [
            ([1, 1, 2, 2, 3, 4, 2, 3], 6, 2, [6, 10, 12]),
            ([3, 8, 7, 8, 7, 5], 2, 2, [11, 15, 15, 15, 12]),
        ]:
            output = self.findXSum(nums, k, x)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
