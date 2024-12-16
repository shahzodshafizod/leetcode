from bisect import bisect_right, bisect_left
from typing import List
import unittest

# https://leetcode.com/problems/maximum-profit-in-job-scheduling/

# python3 -m unittest dp/1235-maximum-profit-in-job-scheduling.py

class Solution(unittest.TestCase):
    # # Approach: Top-Down Dynamic Programming + Binary Search
    # # Time: O(n⋅logn)
    # # Space: O(n)
    # def jobScheduling(self, startTime: List[int], endTime: List[int], profit: List[int]) -> int:
    #     jobs = sorted(zip(startTime, endTime, profit))
    #     n = len(jobs)
    #     memo = [None] * n
    #     def dp(idx: int) -> int:
    #         if idx == n: return 0
    #         if memo[idx] != None:
    #             return memo[idx]
    #         next = bisect_right(jobs, (jobs[idx][1],-1,-1), lo=idx+1, hi=n)
    #         memo[idx] = max(
    #             dp(idx+1), # skip this job
    #             jobs[idx][2] + dp(next), # take this job
    #         )
    #         return memo[idx]
    #     return dp(0)
    
	# # Approach: Bottom-Up Dynamic Programming + Binary Search
    # # Time: O(n⋅logn)
    # # Space: O(n)
    # def jobScheduling(self, startTime: List[int], endTime: List[int], profit: List[int]) -> int:
    #     # sort by starting time
    #     jobs = sorted(zip(startTime, endTime, profit))
    #     n = len(jobs)
    #     dp = [0] * (n+1)
    #     for idx in range(n-1,-1,-1):
    #         next = bisect_right(jobs, (jobs[idx][1],-1,-1), lo=idx+1, hi=n)
    #         dp[idx] = max(
    #             dp[idx+1], # skip this job
    #             jobs[idx][2] + dp[next], # take this job
    #     	)
    #     return dp[0]

	# Approach: Bottom-Up Dynamic Programming + Binary Search
    # Time: O(n⋅logn)
    # Space: O(n)
    def jobScheduling(self, startTime: List[int], endTime: List[int], profit: List[int]) -> int:
		# sort by ending time
        jobs = sorted(zip(endTime, startTime, profit))
        n = len(jobs)
        dp = [0] * (n+1)
        for idx, (_, start, profit) in enumerate(jobs):
            prev = bisect_left(jobs, (start+1,-1,-1), hi=idx)
            dp[idx+1] = max(
                dp[idx], # skip this job
                profit + dp[prev], # take this job
			)
        return dp[n]

    def test(self):
        for startTime, endTime, profit, expected in [
            ([1,1,1], [2,3,4], [5,6,4], 6),
            ([7779], [13973], [4658], 4658),
            ([1,2,3,3], [3,4,5,6], [50,10,40,70], 120),
            ([3559,5372], [8592,11236], [8968,2073], 8968),
            ([1,2,3,4,6], [3,5,10,6,9], [20,20,100,70,60], 150),
            ([497,9813,5338,2256,9062], [7676,18960,14385,9924,12218], [2712,9026,603,5918,9964], 12676),
            ([6,24,45,27,13,43,47,36,14,11,11,12], [31,27,48,46,44,46,50,49,24,42,13,27], [14,4,16,12,20,3,18,6,9,1,2,8], 45),
            ([8882,7030,7842,8505,8586,2257,8970,1669,1124,3159], [18708,10088,7955,9997,9173,4828,13306,5053,1380,8825], [9933,967,5417,1835,2746,3901,773,2248,5062,3712], 24313),
            # ([96,815,776,701,4773,3771,6725,8841,4289,4992,9852,6641,6065,3366,6415,7684,9429,7620,8880,6509], [3816,7818,8658,3359,6845,12795,12943,13524,8099,9911,16365,7667,10043,9038,7822,14342,13581,12347,10690,11189], [8628,8418,9036,7717,1828,1498,5657,1156,4749,2528,8675,4130,3791,6585,655,3686,8462,9630,2263,9595], 22977),
        ]:
            output = self.jobScheduling(startTime, endTime, profit)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
