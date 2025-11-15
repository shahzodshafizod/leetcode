import unittest

# https://leetcode.com/problems/count-all-possible-routes/

# python3 -m unittest dp/1575-count-all-possible-routes.py


class Solution(unittest.TestCase):
    # # Approach #1: Top-Down Dynamic Programming
    # # Time: O(n^2 * fuel)
    # # Space: O(n * fuel)
    # def countRoutes(self, locations: list[int], start: int, finish: int, fuel: int) -> int:
    #     n = len(locations)
    #     memo: dict[tuple[int, int], int] = {}

    #     def dfs(city: int, fuel: int) -> int:
    #         state = (city, fuel)
    #         if state in memo:
    #             return memo[state]
    #         cnt = 1 if city == finish else 0
    #         for next_city in range(n):
    #             if next_city == city:
    #                 continue

    #             remaining_fuel = fuel - abs(locations[city] - locations[next_city])
    #             if remaining_fuel >= 0:
    #                 cnt += dfs(next_city, remaining_fuel)

    #         memo[state] = cnt
    #         return cnt

    #     return dfs(start, fuel) % (10**9 + 7)

    # Approach #2: Bottom-Up Dynamic Programming
    # Time: O(n^2 * fuel)
    # Space: O(n * fuel)
    def countRoutes(self, locations: list[int], start: int, finish: int, fuel: int) -> int:
        n = len(locations)
        dp = [[0] * (fuel + 1) for _ in range(n)]
        dp[start][fuel] = 1
        for f in range(fuel, -1, -1):
            for city in range(n):
                if dp[city][f] == 0:
                    continue
                for next_city in range(n):
                    if next_city == city:
                        continue
                    remaining_fuel = f - abs(locations[city] - locations[next_city])
                    if remaining_fuel >= 0:
                        dp[next_city][remaining_fuel] += dp[city][f]
        return sum(dp[finish]) % (10**9 + 7)

    def test(self):
        for locations, start, finish, fuel, expected in [
            ([2, 3, 6, 8, 4], 1, 3, 5, 4),
            ([4, 3, 1], 1, 0, 6, 5),
            ([5, 2, 1], 0, 2, 3, 0),
            ([2, 1, 5], 0, 0, 3, 2),
            ([1, 2, 3], 0, 2, 40, 615088286),
        ]:
            output = self.countRoutes(locations, start, finish, fuel)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
