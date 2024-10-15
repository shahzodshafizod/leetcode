from typing import List
import unittest

# https://leetcode.com/problems/cherry-pickup/

# python3 -m unittest dp/0741-cherry-pickup.py

class Solution(unittest.TestCase):
    # # Approach: Dynamic Programming (Top-Down)
    # # Time: O(N^4)
    # # Space: O(N^4)
    # def cherryPickup(self, grid: List[List[int]]) -> int:
    #     n = len(grid)
    #     memo = [[[[None] * n for _ in range(n)] for _ in range(n)] for _ in range(n)]
    #     def dfs(r1: int, c1: int, r2: int, c2: int) -> int:
    #         if (r1 == n or c1 == n or r2 == n or c2 == n or
    #             grid[r1][c1] == -1 or grid[r2][c2] == -1):
    #             # this -ve should make result of other instances equal to <= 0
    #             return -300 # or -3*(50+49)+1
    #         if r1==c1==n-1:
    #             # (r1,c1) & (r2,c2)
    #             # both come to the end
    #             # at the same time
    #             return grid[r2][c2]
    #         if memo[r1][c1][r2][c2] != None:
    #             return memo[r1][c1][r2][c2]
    #         count = grid[r1][c1] + (grid[r2][c2] if c1 != c2 else 0)
    #         # person1 -> down, person2 -> down
    #         down_down = dfs(r1+1, c1, r2+1, c2)
    #         # person1 -> down, person2 -> right
    #         down_right = dfs(r1+1, c1, r2, c2+1)
    #         # person1 -> right, person2 -> down
    #         right_down = dfs(r1, c1+1, r2+1, c2)
    #         # person1 -> right, person2 -> right
    #         right_right = dfs(r1, c1+1, r2, c2+1)
    #         count += max(down_down, down_right, right_down, right_right)
    #         memo[r1][c1][r2][c2] = count
    #         return count
    #     return max(0, dfs(0, 0, 0, 0))

    # # Approach: Dynamic Programming (Top-Down)
    # # Time: O(N^3)
    # # Space: O(N^3)
    # def cherryPickup(self, grid: List[List[int]]) -> int:
    #     # (r1+c1) = (r2+c2)
    #     # c2 = r1+c1-r2
    #     n = len(grid)
    #     memo = [[[None] * n for _ in range(n)] for _ in range(n)]
    #     def dfs(r1: int, c1: int, r2: int) -> int:
    #         c2 = r1+c1-r2
    #         if (r1 == n or c1 == n or r2 == n or c2 == n or
    #             grid[r1][c1] == -1 or grid[r2][c2] == -1):
    #             # this -ve should make result of other instances equal to <= 0
    #             return -300 # or -3*(50+49)+1
    #         if r1==c1==n-1:
    #             # (r1,c1) & (r2,c2)
    #             # both come to the end
    #             # at the same time
    #             return grid[r2][c2]
    #         if memo[r1][c1][r2] != None:
    #             return memo[r1][c1][r2]
    #         count = grid[r1][c1] + (grid[r2][c2] if c1 != c2 else 0)
    #         down_down = dfs(r1+1, c1, r2+1)  # c2
    #         down_right = dfs(r1+1, c1, r2)   # c2+1
    #         right_down = dfs(r1, c1+1, r2+1) # c2
    #         right_right = dfs(r1, c1+1, r2)  # c2+1
    #         count += max(down_down, down_right, right_down, right_right)
    #         memo[r1][c1][r2] = count
    #         return count
    #     return max(0, dfs(0, 0, 0))

    # Approach: Dynamic Programming (Bottom-Up)
    # Time: O(N^3)
    # Space: O(N^2)
    def cherryPickup(self, grid: List[List[int]]) -> int:
        n = len(grid)
        dp = [[0] * n for _ in range(n)]
        dp[0][0] = grid[0][0]
        for t in range(1, (n<<1) - 1): # 2*n-1: # of steps to reach (n-1,n-1)
            for i in range(n-1, -1, -1):
                for j in range(n-1, -1, -1):
                    if (t-i < 0 or t-i >= n or t-j < 0 or t-j >= n or
                        grid[i][t-i] == -1 or grid[j][t-j] == -1):
                        dp[i][j] = -1
                        continue
                    if i > 0: dp[i][j] = max(dp[i][j], dp[i-1][j])
                    if j > 0: dp[i][j] = max(dp[i][j], dp[i][j-1])
                    if i > 0 and j > 0: dp[i][j] = max(dp[i][j], dp[i-1][j-1])
                    if dp[i][j] >= 0: dp[i][j] += grid[i][t-i] + (grid[j][t-j] if i != j else 0)
        return max(0, dp[n-1][n-1])

    def test(self) -> None:
        for grid, expected in [
            ([[0,1,-1],[1,0,-1],[1,1,1]], 5),
		    ([[1,1,-1],[1,-1,1],[-1,1,1]], 0),
            ([[1,1,1,0,1],[0,0,0,0,0],[0,0,0,0,0],[0,0,0,0,0],[1,0,1,1,1]], 8),
            ([[1,1,1,0,0],[0,0,1,0,1],[1,0,1,0,0],[0,0,1,0,0],[0,0,1,1,1]], 11),
            ([[0,1,1,0,0],[1,1,1,1,0],[-1,1,1,1,-1],[0,1,1,1,0],[1,0,-1,0,0]], 11),
            ([[1,1,1,1,0,0,0],[0,0,0,1,0,0,0],[0,0,0,1,0,0,1],[1,0,0,1,0,0,0],[0,0,0,1,0,0,0],[0,0,0,1,0,0,0],[0,0,0,1,1,1,1]], 15),
        ]:
            output = self.cherryPickup(grid)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
