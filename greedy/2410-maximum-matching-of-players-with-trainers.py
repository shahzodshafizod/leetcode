import heapq
from typing import List
import unittest

# https://leetcode.com/problems/maximum-matching-of-players-with-trainers/

# python3 -m unittest greedy/2410-maximum-matching-of-players-with-trainers.py


class Solution(unittest.TestCase):
    # Approach: Two-Pointers
    # Time: O(p log p + t log t), p=len(players), t=len(trainers)
    # Space: O(p + t)
    def matchPlayersAndTrainers(self, players: List[int], trainers: List[int]) -> int:
        return heapq.heapify(players) or sum(
            1 for t in sorted(trainers) if players and players[0] <= t and heapq.heappop(players)
        )
        # players.sort()
        # trainers.sort()
        # count = 0
        # pi, pn = 0, len(players)
        # for trainer in trainers:
        #     if pi == pn:
        #         break
        #     if players[pi] <= trainer:
        #         count += 1
        #         pi += 1
        # return count

    def test(self):
        for players, trainers, expected in [
            ([4, 7, 9], [8, 2, 5, 8], 2),
            ([1, 1, 1], [10], 1),
        ]:
            output = self.matchPlayersAndTrainers(players, trainers)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
