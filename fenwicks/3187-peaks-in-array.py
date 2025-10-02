from typing import List
import unittest

# https://leetcode.com/problems/peaks-in-array/

# python3 -m unittest fenwicks/3187-peaks-in-array.py


class Solution(unittest.TestCase):
    def countOfPeaks(self, nums: List[int], queries: List[List[int]]) -> List[int]:
        n = len(nums)
        bitree = [0] * (n + 1)

        def update(index: int, value: int):
            index += 1
            while index <= n:
                bitree[index] += value
                index += index & -index

        def query(index: int) -> int:
            total = 0
            index += 1
            while index > 0:
                total += bitree[index]
                index -= index & -index
            return total

        def is_peak(i: int) -> bool:
            return 0 < i < n - 1 and nums[i - 1] < nums[i] > nums[i + 1]

        for i in range(n):
            if is_peak(i):
                update(i, 1)

        counts: List[int] = []
        for t, first, second in queries:
            if t == 1:
                left, right = first, second
                counts.append(query(right - 1) - query(left) if left < right else 0)
            else:
                index, value = first, second
                before = [is_peak(index - 1), is_peak(index), is_peak(index + 1)]
                nums[index] = value
                for i in range(len(before)):
                    curr = is_peak(index + i - 1)
                    if before[i] and not curr:
                        update(index + i - 1, -1)
                    elif not before[i] and curr:
                        update(index + i - 1, 1)

        return counts

    def test(self):
        for nums, queries, expected in [
            ([3, 1, 4, 2, 5], [[2, 3, 4], [1, 0, 4]], [0]),
            ([4, 1, 4, 2, 1, 5], [[2, 2, 4], [1, 0, 2], [1, 0, 4]], [0, 1]),
            ([4, 9, 4, 10, 7], [[2, 3, 2], [2, 1, 3], [1, 2, 3]], [0]),
            ([5, 4, 8, 6], [[1, 2, 2], [1, 1, 2], [2, 1, 6]], [0, 0]),
        ]:
            output = self.countOfPeaks(nums, queries)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
