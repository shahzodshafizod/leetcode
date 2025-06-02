from typing import List
import unittest

# https://leetcode.com/problems/candy/

# python3 -m unittest greedy/0135-candy.py

class Solution(unittest.TestCase):
    # # Approach #1: Sorting
    # # Time: O(nlogn)
    # # Space: O(n)
    # def candy(self, ratings: List[int]) -> int:
    #     n = len(ratings)
    #     cnts = [1] * n
    #     candies = 0
    #     for idx in sorted([idx for idx in range(n)], key=lambda x: ratings[x]):
    #         if idx > 0 and ratings[idx] > ratings[idx-1]:
    #             cnts[idx] = cnts[idx-1]+1
    #         if idx+1 < n and ratings[idx] > ratings[idx+1]:
    #             cnts[idx] = max(cnts[idx], cnts[idx+1]+1)
    #         candies += cnts[idx]
    #     return candies

    # Approach #2: Greedy
    # Time: O(n)
    # Space: O(n)
    def candy(self, ratings: List[int]) -> int:
        n = len(ratings)
        cnts = [1] * n
        for idx in range(1, n):
            if ratings[idx-1] < ratings[idx]:
                cnts[idx] = cnts[idx-1]+1
        candies = cnts[-1]
        for idx in range(n-2,-1,-1):
            if ratings[idx] > ratings[idx+1]:
                cnts[idx] = max(cnts[idx], cnts[idx+1]+1)
            candies += cnts[idx]
        return candies

    def test(self):
        for ratings, expected in [
            ([1,0,2], 5),
            ([1,2,2], 4),
            ([60,80,100,100,100,100,100], 10),
            ([100,80,70,60,70,80,90,100,90,80,70,60,60], 35),
            ([6,7,6,5,4,3,2,1,0,0,0,1,0], 42),
            # ([20000,20000,16001,16001,16002,16002,16003,16003,16004,16004,16005,16005,16006,16006,16007,16007,16008,16008,16009,16009,16010,16010,16011,16011,16012,16012,16013,16013,16014,16014,16015,16015,16016,16016,16017,16017,16018,16018,16019,16019,16020,16020,16021,16021,16022,16022,16023,16023,16024,16024], 74),
        ]:
            output = self.candy(ratings)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
