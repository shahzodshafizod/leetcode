from typing import List
import unittest

# https://leetcode.com/problems/minimum-difficulty-of-a-job-schedule/

# python3 -m unittest dp/1335-minimum-difficulty-of-a-job-schedule.py

class Solution(unittest.TestCase):
    # # Approach #1: Top-Down Dynamic Programming
    # # Time: O(N^2 * D)
    # # Space: O(N*D)
    # def minDifficulty(self, jobDifficulty: List[int], d: int) -> int:
    #     MAX = 300_001
    #     n = len(jobDifficulty)
    #     if n < d: return -1
    #     max_right = [0] * (n+1)
    #     for idx in range(n-1, -1, -1):
    #         max_right[idx] = max(max_right[idx+1], jobDifficulty[idx])
    #     dp = [[None] * (d+1) for _ in range(n)]
    #     def dfs(idx: int, day: int) -> int:
    #         if day == 1:
    #             # only one day left, we have to do all the remaining jobs
    #             return max_right[idx]
    #         if dp[idx][day] != None:
    #             return dp[idx][day]
    #         max_diff, min_diff = 0, MAX
    #         for j in range(idx, n-day+1):
    #             max_diff = max(max_diff, jobDifficulty[j])
    #             min_diff = min(min_diff, max_diff + dfs(j+1, day-1))
    #         dp[idx][day] = min_diff
    #         return min_diff
    #     return dfs(0, d)

    # # Approach #2: Bottom-Up Dynamic Programming
    # # Time: O(N^2 * D)
    # # Space: O(N*D)
    # def minDifficulty(self, jobDifficulty: List[int], d: int) -> int:
    #     n = len(jobDifficulty)
    #     if n < d:
    #         # not enough tasks to split into each day
    #         return -1
    #     dp = [[0] * d for _ in range(n)]
    #     dp[0][0] = jobDifficulty[0]
    #     for idx in range(1, n):
    #         # day=1-1, only one day left, we have to do all the remaining jobs
    #         # fill out first column based on max(jobDiffculty[0:i])
    #         dp[idx][0] = max(jobDifficulty[idx], dp[idx-1][0])
    #     for day in range(1, d):
    #         for idx in range(day, n):
    #             max_diff = 0
    #             dp[idx][day] = 300_001
    #             for j in range(idx, day-1, -1):
    #                 max_diff = max(max_diff, jobDifficulty[j])
    #                 dp[idx][day] = min(dp[idx][day], max_diff + dp[j-1][day-1])
    #     return dp[n-1][d-1]

    # # Approach #3: Bottom-Up Dynamic Programming (Space Optimized)
    # # Time: O(N^2 * D)
    # # Space: O(N)
    # def minDifficulty(self, jobDifficulty: List[int], d: int) -> int:
    #     n = len(jobDifficulty)
    #     # not enough tasks to split into each day
    #     if n < d: return -1
    #     prev, curr = [0] * n, [0] * n
    #     curr[0] = jobDifficulty[0]
    #     for idx in range(1, n):
    #         # day=1-1, only one day left, we have to do all the remaining jobs
    #         # fill out first column based on max(jobDiffculty[0:i])
    #         curr[idx] = max(jobDifficulty[idx], curr[idx-1])
    #     for day in range(1, d):
    #         prev, curr = curr, prev
    #         for idx in range(day, n):
    #             max_diff = 0
    #             curr[idx] = 300_001
    #             for j in range(idx, day-1, -1):
    #                 max_diff = max(max_diff, jobDifficulty[j])
    #                 curr[idx] = min(curr[idx], max_diff + prev[j-1])
    #     return curr[n-1]

    # Approach #4: Stack
    # Time: O(N*D)
    # Space: O(N)
    def minDifficulty(self, jobDifficulty: List[int], d: int) -> int:
        n = len(jobDifficulty)
        if n < d: return -1
        prev = [0] * n
        curr = [300_001] * n
        for day in range(d):
            prev, curr = curr, prev
            stack = []
            for idx in range(day, n):
                curr[idx] = prev[idx-1]+jobDifficulty[idx] if idx else jobDifficulty[idx]
                while stack and jobDifficulty[stack[-1]] <= jobDifficulty[idx]:
                    j = stack.pop()
                    curr[idx] = min(curr[idx], curr[j] - jobDifficulty[j] + jobDifficulty[idx])
                if stack:
                    curr[idx] = min(curr[idx], curr[stack[-1]])
                stack.append(idx)
        return curr[n-1]

    def test(self) -> None:
        for jobDifficulty, d, expected in [
            ([6,5,4,3,2,1], 2, 7),
            ([9,9,9], 4, -1),
            ([1,1,1], 3, 3),
            ([1], 1, 1),
            ([186,398,479,206,885,423,805,112,925,656,16,932,740,292,671,360], 4, 1803),
            # ([182,112,52,22,74,167,282,126,179,8,125,271,260,119,228,294,214,131,180,247,20,295,41,175,118,55,27,251,21,136,250,257,106,97,272,75,49,16,9,63,218,53,244,289,211,70,246,268,51,56,287,240,267,189,238,296,144,64,26,105,201,283,181,193,274,35,82,73,156,212,232,249,122,134,43,298,124,1,159,46,280,204,270,217,299,31,23,151,5,83,15,6,96,203,197,103,128,292,107,77,264,87,195,173,13,196,261,239,225,109,286,108,45,162,157,113,142,184,290,176,213,153,72,199,48,163,161,133,104,198,276,130,220,230,224,226,207,242,281,117,259,235,25,94,265,223,19,85,168,285,91,186,65,54,98,2,3,293,28,30,58,0,165,132,44,234,255,227,248,185,4,237,33,93,50,14,39,149,164,253,216,262,80,208,174,279,202,111,135,37,241,42,66,150,277,205,278,252,231,160,258,81,273,266,57,158,95,219,139,178,138,24,148,206,84,89,256,17,145,221,275,200,68,194,154,146,59,61,210,127,288,11,291,123,114,140,18,155,183,166,147,60,101,116,172,137,88,191,71,169,236,215,62,297,110,284,47,36,79,120,100,254,121,40,170,152,86,90,141,99,143,192,187,78,190,34,76,92,263,67,222,233,10,245,115,102,38,188,32,269,171,12,229,29,177,209,243,69,7,129], 10, 1006),
            # ([210,118,91,249,70,82,131,62,230,24,6,291,264,185,234,254,8,7,188,81,241,119,61,101,0,214,275,266,267,43,26,235,69,223,212,190,240,183,106,147,122,28,113,297,260,50,149,220,86,194,293,1,262,25,52,157,172,123,92,137,243,294,161,196,191,225,63,253,12,136,163,141,78,114,247,177,186,233,132,182,135,58,289,144,115,173,5,104,227,134,57,280,245,33,265,110,209,189,53,167,218,154,120,36,90,127,29,231,145,108,32,88,180,9,125,17,109,187,21,202,107,34,77,37,20,236,128,268,44,96,140,228,205,124,232,272,171,23,213,46,282,111,18,164,263,279,290,273,30,60,75,98,158,219,146,15,259,274,55,176,103,73,206,79,93,130,237,42,47,204,277,197,84,226,45,85,112,74,39,238,71,217,244,179,11,4,257,166,99,251,169,117,22,181,72,159,211,192,193,38,261,14,207,35,105,199,87,208,285,67,121,102,242,246,59,292,295,198,68,40,174,201,139,296,152,221,31,299,27,168,283,239,65,271,156,269,200,49,298,151,288,143,80,252,203,116,256,138,56,170,97,195,150,133,129,287,76,142,54,83,19,270,95,255,184,284,215,3,162,16,41,278,160,248,148,229,89,281,48,126,66,286,10,100,216,153,222,155,175,94,165,250,2,258,51,13,178,276,64,224], 10, 1117),
            # ([186,141,297,63,236,288,91,125,175,209,201,129,45,23,215,256,260,42,289,132,197,187,142,93,158,229,34,258,29,15,293,101,255,156,203,105,224,180,25,210,184,99,252,145,80,24,261,243,66,39,6,35,291,220,26,57,190,36,273,55,153,237,177,250,264,189,265,152,263,74,157,107,50,92,124,67,110,198,48,280,131,235,267,16,100,223,213,88,77,207,179,14,271,193,216,71,98,221,166,86,94,165,168,269,13,58,217,116,173,128,47,160,27,279,163,52,8,185,114,17,147,211,136,272,161,59,278,64,137,53,172,191,121,228,287,245,20,140,44,167,169,218,115,89,122,206,233,162,176,78,150,113,178,104,195,87,242,270,123,159,112,108,284,222,38,69,10,134,274,75,139,192,54,296,7,90,294,202,298,133,12,241,30,1,72,9,79,119,51,146,295,170,84,281,299,246,283,28,49,151,85,212,32,253,268,21,164,181,219,73,232,37,149,174,227,120,199,205,117,138,143,231,275,40,249,276,135,106,196,103,248,65,60,43,62,76,194,31,230,2,266,259,244,148,83,118,97,61,154,292,200,247,0,109,11,4,226,5,82,102,225,22,240,171,127,204,282,208,19,257,46,95,251,239,183,130,286,18,262,290,238,41,188,33,111,234,254,126,277,182,285,56,81,96,3,214,68,70,144,155], 8, 879),
            # ([135,223,23,39,94,150,141,247,38,207,15,275,230,259,279,218,85,44,53,299,105,102,70,180,127,156,184,175,101,290,30,278,13,55,168,7,29,255,145,42,108,273,3,1,282,136,206,202,163,242,190,277,134,199,0,173,260,235,10,76,83,217,96,165,74,64,21,174,47,246,181,81,51,291,32,225,267,151,35,114,111,189,97,59,122,69,36,22,245,228,248,27,179,208,182,234,215,293,196,158,152,295,224,205,87,221,2,103,187,57,191,244,211,219,148,166,204,89,128,193,95,197,240,78,46,54,33,155,162,169,41,171,25,164,200,80,276,72,109,146,73,287,229,71,116,52,93,120,91,236,271,233,222,63,84,88,160,281,113,12,132,249,49,283,100,220,209,159,60,117,18,198,241,86,62,212,149,28,37,107,177,121,142,253,4,153,34,176,192,137,147,261,167,186,262,178,8,98,19,280,140,216,61,11,79,210,118,77,285,119,157,75,143,92,26,264,112,45,250,43,139,66,56,231,288,272,99,237,24,131,268,126,161,31,194,20,263,214,185,292,289,125,296,251,188,58,130,172,195,67,138,65,170,6,110,258,104,203,266,48,232,274,123,50,286,297,154,238,257,90,17,298,252,213,269,239,82,14,129,106,254,265,16,294,144,183,201,270,115,284,256,9,40,243,227,226,5,133,124,68], 9, 1007),
            # ([174,151,119,140,150,6,205,257,222,293,237,156,123,41,261,211,244,288,233,285,299,7,167,181,36,12,134,50,21,157,281,49,44,66,76,173,108,129,115,197,109,229,126,208,245,255,185,75,133,241,104,249,284,227,149,54,2,253,47,204,265,110,143,96,37,99,42,269,61,130,178,224,252,180,136,283,146,291,46,187,282,232,242,69,212,70,127,221,43,9,111,278,92,273,84,225,59,80,145,162,103,289,296,8,210,189,16,120,55,228,195,48,13,97,34,93,15,53,58,262,192,73,23,238,132,196,183,160,179,65,226,106,165,223,112,240,141,256,168,60,137,11,51,182,19,0,82,166,188,63,1,71,280,68,190,144,138,114,297,234,62,214,248,38,260,186,279,128,88,57,267,85,17,216,163,290,139,40,203,268,217,125,201,83,105,20,3,171,236,220,142,124,86,275,87,200,117,33,81,131,27,98,100,89,22,32,231,122,101,235,77,72,250,251,193,219,91,5,276,207,164,198,28,39,169,14,199,247,64,243,67,271,94,239,298,264,177,107,209,153,102,295,118,29,24,4,184,266,18,294,154,254,287,215,213,148,155,258,277,90,194,263,35,270,26,286,10,45,246,176,56,191,135,202,31,147,113,292,272,78,206,170,25,159,175,230,95,259,218,161,79,52,152,274,158,74,30,121,116,172], 1, 299),
        ]:
            output = self.minDifficulty(jobDifficulty, d)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
