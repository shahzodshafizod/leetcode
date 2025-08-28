from typing import List
import unittest
from sortedcontainers import SortedList

# https://leetcode.com/problems/distribute-elements-into-two-arrays-ii/

# python3 -m unittest fenwicks/3072-distribute-elements-into-two-arrays-ii.py


class Solution(unittest.TestCase):
    # # Approach #1: Brute-Force
    # # Time: O(nn)
    # # Space: O(n)
    # def resultArray(self, nums: List[int]) -> List[int]:
    #     def greaterCount(arr: List[int], val: int) -> int:
    #         return sum(1 for x in arr if x > val)

    #     arr1: List[int] = [nums[0]]
    #     arr2: List[int] = [nums[1]]
    #     for num in nums[2:]:
    #         gc1 = greaterCount(arr1, num)
    #         gc2 = greaterCount(arr2, num)
    #         if gc1 > gc2 or gc1 == gc2 and len(arr1) <= len(arr2):
    #             arr1.append(num)
    #         else:
    #             arr2.append(num)
    #     return arr1 + arr2

    # Approach #2: Sorted List
    # Time: O(n log n)
    # Space: O(n)
    def resultArray(self, nums: List[int]) -> List[int]:
        def greaterCount(sl: SortedList[int], val: int) -> int:
            return len(sl) - sl.bisect_right(val)

        arr1: List[int] = [nums[0]]
        sl1: SortedList[int] = SortedList(arr1)
        arr2: List[int] = [nums[1]]
        sl2: SortedList[int] = SortedList(arr2)
        for num in nums[2:]:
            g1 = greaterCount(sl1, num)
            g2 = greaterCount(sl2, num)
            if g1 > g2:
                arr1.append(num)
                sl1.add(num)
            elif g2 > g1:
                arr2.append(num)
                sl2.add(num)
            elif len(arr1) > len(arr2):
                arr2.append(num)
                sl2.add(num)
            else:
                arr1.append(num)
                sl1.add(num)
        return arr1 + arr2

    # # Approach #3: Binary Indexed Tree
    # # Time: O(n log n)
    # # Space: O(n)
    # def resultArray(self, nums: List[int]) -> List[int]:
    #     class BIT:
    #         def __init__(self, size: int):
    #             self.tree = [0] * (size + 1)
    #             self.size = size

    #         def update(self, idx: int):
    #             idx += 1  # indices in BIT are 1-based
    #             while idx <= self.size:
    #                 self.tree[idx] += 1
    #                 idx += idx & -idx

    #         def smallerCount(self, idx: int) -> int:
    #             # returns sum of frequences between [0 and idx]
    #             idx += 1  # indices in BIT are 1-based
    #             count = 1
    #             while idx > 0:
    #                 count += self.tree[idx]
    #                 idx -= idx & -idx
    #             return count

    #         def query(self, idx: int) -> int:
    #             return self.smallerCount(self.size - 1) - self.smallerCount(idx)

    #     def greaterCount(bit: BIT, idx: int) -> int:
    #         return bit.query(idx)

    #     indices = {x: i for i, x in enumerate(sorted(set(nums)))}
    #     size = len(indices)

    #     bit1, bit2 = BIT(size), BIT(size)
    #     bit1.update(indices[nums[0]])
    #     bit2.update(indices[nums[1]])

    #     arr1, arr2 = [nums[0]], [nums[1]]

    #     for i in range(2, len(nums)):
    #         idx = indices[nums[i]]
    #         gc1 = greaterCount(bit1, idx)
    #         gc2 = greaterCount(bit2, idx)

    #         if gc1 > gc2 or gc1 == gc2 and len(arr1) <= len(arr2):
    #             arr1.append(nums[i])
    #             bit1.update(idx)
    #         else:
    #             arr2.append(nums[i])
    #             bit2.update(idx)

    #     return arr1 + arr2

    def test(self):
        for nums, expected in [
            ([2, 1, 3, 3], [2, 3, 1, 3]),
            ([5, 14, 3, 1, 2], [5, 3, 1, 2, 14]),
            ([3, 3, 3, 3], [3, 3, 3, 3]),
        ]:
            output = self.resultArray(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
