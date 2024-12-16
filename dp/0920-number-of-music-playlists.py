import unittest

# https://leetcode.com/problems/number-of-music-playlists/

# python3 -m unittest dp/0920-number-of-music-playlists.py

class Solution(unittest.TestCase):
    # # Approach #1: Top-Down Dynamic Programming
    # # Time: O(gn), g=goal
    # # Space: O(gn)
    # def numMusicPlaylists(self, n: int, goal: int, k: int) -> int:
    #     MOD = int(1e9+7)
    #     memo = [[None] * (n+1) for _ in range(goal+1)]
    #     def dp(slots: int, songs: int) -> int:
    #         if slots == 0 and songs == 0: return 1
    #         if slots == 0 or songs == 0: return 0
    #         if memo[slots][songs] != None:
    #             return memo[slots][songs]
    #         # add a new song: n-songs+1 = # of choices
    #         count = (n-songs+1) * dp(slots-1, songs-1) % MOD
    #         # replay an old song: songs-k = # of choices
    #         if songs > k:
    #             count = (count + (songs-k) * dp(slots-1, songs)) % MOD
    #         memo[slots][songs] = count
    #         return count
    #     return dp(goal, n)

    # # Approach #2: Bottom-Up Dynamic Programming
    # # Time: O(gn), g=goal
    # # Space: O(gn)
    # def numMusicPlaylists(self, n: int, goal: int, k: int) -> int:
    #     MOD = int(1e9+7)
    #     dp = [[0] * (n+1) for _ in range(goal+1)]
    #     dp[0][0] = 1
    #     for slots in range(1, goal+1):
    #         for songs in range(1, min(slots,n)+1):
    #             # add a new song: n-songs+1 = # of choices
    #             dp[slots][songs] = (n-songs+1) * dp[slots-1][songs-1] % MOD
    #             # replay an old song: songs-k = # of choices
    #             if songs > k:
    #                 dp[slots][songs] = (
    #                     dp[slots][songs] +
    #                     (songs-k) * dp[slots-1][songs]
    #                 ) % MOD
    #     return dp[goal][n]

    # Approach #3: Bottom-Up Dynamic Programming (Space-Optimized)
    # Time: O(gn), g=goal
    # Space: O(n)
    def numMusicPlaylists(self, n: int, goal: int, k: int) -> int:
        MOD = int(1e9+7)
        prev, curr = [0]*(n+1), [0]*(n+1)
        curr[0] = 1
        for slots in range(1, goal+1):
            prev, curr = curr, prev
            curr[0] = 0
            for songs in range(1, min(slots,n)+1):
                # add a new song: n-songs+1 = # of choices
                curr[songs] = (n-songs+1) * prev[songs-1] % MOD
                # replay an old song: songs-k = # of choices
                if songs > k:
                    curr[songs] = (
                        curr[songs] +
                        (songs-k) * prev[songs]
                    ) % MOD
        return curr[n]

    def test(self):
        for n, goal, k, expected in [
            (3, 3, 1, 6),
            (2, 3, 0, 6),
            (2, 3, 1, 2),
            (2, 4, 0, 14),
            (16, 16, 4, 789741546),
        ]:
            output = self.numMusicPlaylists(n, goal, k)
            self.assertEqual(expected, output, "expected: {expected}, output: {output}")
