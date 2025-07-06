from math import factorial

# from typing import List
import unittest

# https://leetcode.com/problems/permutation-sequence/

# python3 -m unittest maths/0060-permutation-sequence.py


class Solution(unittest.TestCase):
    # # Solved with the help of this problem:
    # # 31. Next Permutation (https://leetcode.com/problems/next-permutation/)
    # # Time: O(n x k)
    # # Space: O(n)
    # def getPermutation(self, n: int, k: int) -> str:
    #     def nextPermutation(nums: List[int]) -> None:
    #         n = len(nums)
    #         pivot = n-2
    #         while pivot >= 0 and nums[pivot] >= nums[pivot+1]:
    #             pivot -= 1
    #         if pivot >= 0:
    #             successor = n-1
    #             while nums[successor] < nums[pivot]:
    #                 successor -= 1
    #             nums[successor], nums[pivot] = nums[pivot], nums[successor]
    #         left, right = pivot+1, n-1
    #         while left < right and nums[left] > nums[right]:
    #             nums[left], nums[right] = nums[right], nums[left]
    #             left += 1
    #             right -= 1

    #     nums = [i for i in range(1, n+1)]
    #     for _ in range(k-1):
    #         nextPermutation(nums)
    #     return "".join(str(num) for num in nums)

    # # Time: O(n+k)
    # # Space: O(n)
    # def getPermutation(self, n: int, k: int) -> str:
    #     nums = list(range(1, n+1))
    #     fact = factorial(n)
    #     result = []
    #     for i in range(n, 0, -1):
    #         fact //= i
    #         idx = 0
    #         while k > fact:
    #             idx += 1
    #             k -= fact
    #         result.append(nums.pop(idx))
    #     return "".join(str(num) for num in result)

    # Time: O(n)
    # Space: O(n)
    def getPermutation(self, n: int, k: int) -> str:
        nums = list(range(1, n + 1))
        fact = factorial(n)
        permutations = []
        k -= 1  # to make it 0-based
        for i in range(n, 0, -1):
            fact //= i
            idx = k // fact
            permutations.append(str(nums.pop(idx)))
            k %= fact  # k -= idx * fact
        return "".join(permutations)

    def test(self):
        for n, k, expected in [
            (3, 3, "213"),
            (4, 9, "2314"),
            (3, 1, "123"),
            (4, 17, "3412"),
        ]:
            output = self.getPermutation(n, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
