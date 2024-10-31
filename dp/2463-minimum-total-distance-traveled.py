from typing import List
import unittest

# https://leetcode.com/problems/minimum-total-distance-traveled/

# python3 -m unittest dp/2463-minimum-total-distance-traveled.py

class Solution(unittest.TestCase):
    # # Approach: Top-Down Dynamic Programming
    # # Time: O(RF), R=len(robot), F=sum(factory[i][1])
    # # Space: O(RF)
    # def minimumTotalDistance(self, robot: List[int], factory: List[List[int]]) -> int:
    #     robot.sort()
    #     factory.sort()
    #     factories = []
    #     for pos, limit in factory:
    #         factories.extend([pos]*limit)
    #     rlen, flen = len(robot), len(factories)
    #     memo = [[None] * flen for _ in range(rlen)]
    #     def dp(rid: int, fid: int) -> int:
    #         if rid == rlen: return 0
    #         if fid == flen: return float("inf")
    #         if memo[rid][fid] != None:
    #             return memo[rid][fid]
    #         distance = abs(robot[rid] - factories[fid])
    #         cost = min(
    #             distance + dp(rid+1, fid+1), # include factory
    #             dp(rid, fid+1), # skip factory
    #         )
    #         memo[rid][fid] = cost
    #         return cost
    #     return dp(0, 0)

    # Approach: Bottom-Up Dynamic Programming
    # Time: O(RF), R=len(robot), F=sum(factory[i][1])
    # Space: O(RF)
    def minimumTotalDistance(self, robot: List[int], factory: List[List[int]]) -> int:
        robot.sort()
        factory.sort()
        factories = []
        for factory in factory:
            factories.extend([factory[0]] * factory[1])
        rlen, flen = len(robot), len(factories)
        dp = [[0] * (flen+1) for _ in range(rlen+1)]
        for rid in range(rlen-1,-1,-1):
            dp[rid][flen] = float("inf") # No factories left
            for fid in range(flen-1,-1,-1):
                distance = abs(robot[rid]-factories[fid])
                dp[rid][fid] = min(
                    distance+dp[rid+1][fid+1], # include factory
                    dp[rid][fid+1], # skip factory
                )
        return dp[0][0]

    def test(self) -> None:
        for robot, factory, expected in [
            ([7], [[0,1]], 7),
		    ([1,-1], [[-2,1],[2,1]], 2),
            ([0,4,6], [[2,2],[6,2]], 4),
            ([-40,-14,-8,1,3,5,39], [[-34,5],[28,2],[-12,3]], 92),
            ([9,11,99,101], [[10,1],[7,1],[14,1],[100,1],[96,1],[103,1]], 6),
            ([-551,-510,-344,-264,-242,-50,202,185,700,947], [[-79,5],[-534,5]], 3172),
            ([403625544,670355988,886437985,224430896,126139936,-477101480,-868159607,-293937930], [[333473422,7],[912209329,7],[468372740,7],[-765827269,4],[155827122,4],[635462096,2],[-300275936,2],[-115627659,0]], 509199280),
            ([9486709,305615257,214323605,282129380,763907021,-662831631,628054452,-132239126,50488558,381544523,-656107497,810414334,421675516,-304916551,571202823,478460906,-972398628,325714139,-86452967,660743346], [[-755430217,18],[382914340,2],[977509386,4],[94299927,9],[32194684,16],[-371001457,2],[-426296769,13],[-284404215,8],[-421288725,0],[-893030428,2],[291827872,17],[-766616824,8],[-730172656,17],[-722387876,1],[510570520,20],[756326049,7],[-242350340,14],[6585224,19],[-173457765,15]], 925405949),
        ]:
            output = self.minimumTotalDistance(robot, factory)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
