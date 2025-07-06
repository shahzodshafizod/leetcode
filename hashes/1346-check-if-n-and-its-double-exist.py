from typing import List
import unittest

# https://leetcode.com/problems/check-if-n-and-its-double-exist/

# python3 -m unittest hashes/1346-check-if-n-and-its-double-exist.py


class Solution(unittest.TestCase):
    # # Approach #1: Brute-Force
    # # Time: O(nn)
    # # Space: O(1)
    # def checkIfExist(self, arr: List[int]) -> bool:
    #     for i in range(len(arr)):
    #         for j in range(len(arr)):
    #             if i != j and arr[i] == 2*arr[j]:
    #                 return True
    #     return False

    # # Approach #2: Sorting + Binary Search
    # # Time: O(nlogn)
    # # Space: O(n), In Python, the sort() method sorts a list using
    # # the Timsort algorithm which is a combination of Merge Sort
    # # and Insertion Sort and has a space complexity of O(n)
    # def checkIfExist(self, arr: List[int]) -> bool:
    #     arr.sort()
    #     n = len(arr)
    #     for idx, num in enumerate(arr):
    #         num += num # its double
    #         left, right = 0, n-1
    #         while left <= right:
    #             mid = (left + right) // 2
    #             if arr[mid] < num:
    #                 left = mid+1
    #             elif arr[mid] > num:
    #                 right = mid-1
    #             elif mid != idx:
    #                 return True
    #             else:
    #                 break
    #     return False

    # Approach #3: Hash Set
    # Time: O(n)
    # Space: O(n)
    def checkIfExist(self, arr: List[int]) -> bool:
        seen = set()
        for num in arr:
            if num * 2 in seen or num / 2 in seen:
                return True
            seen.add(num)
        return False

    def testCheckIfExist(self) -> None:
        for arr, expected in [
            ([0], False),
            ([0, 0], True),
            ([10, 2, 5, 3], True),
            ([7, 1, 14, 11], True),
            ([3, 1, 7, 11], False),
            ([-16, -13, 8], False),
            ([7, 15, 3, 4, 30], True),
            ([0, 2, -7, 11, 4, 18], True),
            ([10, 2, 7, 3, 0, 0, -13], True),
            ([-2, 0, 10, -19, 4, 6, -8], False),
            # ([357,-53,277,-706,980,826,93,-352,-669,989,-193,920,209,-574,-389,221,383,352,665,873,759,-480,-64,-103,-721,-623,-642,-680,20,-168,528,-336,-656,264,581,-714,-458,721,815,106,328,476,351,325,-954,890,-174,635,95,-443,338,907,-648,113,-278,498,532,-778,95,-487,-909,-642,774,296,417,-132,-951,857,-867,321,-960,180,108,-984,-54,103,703,-118,-252,235,577,-703,842,-638,-888,-981,-246,484,202,328,661,447,-831,946,-888,-749,-702], True),
            # ([2,3,5,7,11,13,17,19,23,29,31,37,41,43,47,53,59,61,67,71,73,79,83,89,97,101,103,107,109,113,127,131,137,139,149,151,157,163,167,173,179,181,191,193,197,199,211,223,227,229,233,239,241,251,257,263,269,271,277,281,283,293,307,311,313,317,331,337,347,349,353,359,367,373,379,383,389,397,401,409,419,421,431,433,439,443,449,457,461,463,467,479,487,491,499,503,509,521,523,541,547,557,563,569,571,577,587,593,599,601,607,613,617,619,631,641,643,647,653,659,661,673,677,683,691,701,709,719,727,733,739,743,751,757,761,769,773,787,797,809,811,821,823,827,829,839,853,857,859,863,877,881,883,887,907,911,919,929,937,941,947,953,967,971,977,983,991,997], False),
        ]:
            output = self.checkIfExist(arr)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
