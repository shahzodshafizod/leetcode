from typing import List, Tuple
import unittest
import heapq

# https://leetcode.com/problems/minimum-cost-to-reach-destination-in-time/

# python3 -m unittest dp/1928-minimum-cost-to-reach-destination-in-time.py


class Solution(unittest.TestCase):
    # # Approach #1: Brute-Force (DFS with memoization)
    # # Time: O(n * maxTime * E) where E is number of edges
    # # Space: O(n * maxTime) - memoization cache
    # def minCost(self, maxTime: int, edges: List[List[int]], passingFees: List[int]) -> int:
    #     n = len(passingFees)
    
    #     # Build adjacency list
    #     graph: List[List[Tuple[int, int]]] = [[] for _ in range(n)]
    #     for u, v, time in edges:
    #         graph[u].append((v, time))
    #         graph[v].append((u, time))
    
    #     # memo[city][time] = minimum cost to reach city with exactly time spent
    #     MAX = 1 << 31
    #     memo: Dict[Tuple[int, int], int] = {}
    #     def dfs(city: int, time_spent: int) -> int:
    #         # Base cases
    #         if city == n - 1:
    #             return passingFees[city]
    
    #         if time_spent > maxTime:
    #             return MAX
    
    #         if (city, time_spent) in memo:
    #             return memo[(city, time_spent)]
    
    #         min_cost = MAX
    #         for neighbor, travel_time in graph[city]:
    #             if time_spent + travel_time <= maxTime:
    #                 cost = dfs(neighbor, time_spent + travel_time)
    #                 if cost != MAX:
    #                     min_cost = min(min_cost, passingFees[city] + cost)
    
    #         memo[(city, time_spent)] = min_cost
    #         return min_cost
    
    #     result = dfs(0, 0)
    #     return result if result != MAX else -1

    # Approach #2: Dynamic Programming (Dijkstra-like with time constraint)
    # Time: O((n * maxTime) * log(n * maxTime)) - priority queue operations
    # Space: O(n * maxTime) - DP table and priority queue
    def minCost(self, maxTime: int, edges: List[List[int]], passingFees: List[int]) -> int:
        n = len(passingFees)

        # Build adjacency list
        graph: List[List[Tuple[int, int]]] = [[] for _ in range(n)]
        for u, v, time in edges:
            graph[u].append((v, time))
            graph[v].append((u, time))

        # dp[city][time] = minimum cost to reach city with time spent
        # We track minimum cost for each (city, time) state
        dp = [[float('inf')] * (maxTime + 1) for _ in range(n)]
        dp[0][0] = passingFees[0]

        # Priority queue: (cost, city, time)
        pq = [(passingFees[0], 0, 0)]

        while pq:
            cost, city, time = heapq.heappop(pq)

            # If we reached destination
            if city == n - 1:
                return cost

            # Skip if we found a better path already
            if cost > dp[city][time]:
                continue

            # Explore neighbors
            for neighbor, travel_time in graph[city]:
                new_time = time + travel_time
                if new_time <= maxTime:
                    new_cost = cost + passingFees[neighbor]

                    # Only proceed if this is better than previous best for this state
                    if new_cost < dp[neighbor][new_time]:
                        dp[neighbor][new_time] = new_cost
                        heapq.heappush(pq, (new_cost, neighbor, new_time))

        return -1

    def test(self):
        for maxTime, edges, passingFees, expected in [
            (
                30,
                [[0, 1, 10], [1, 2, 10], [2, 5, 10], [0, 3, 1], [3, 4, 10], [4, 5, 15]],
                [5, 1, 2, 20, 20, 3],
                11,
            ),
            (
                29,
                [[0, 1, 10], [1, 2, 10], [2, 5, 10], [0, 3, 1], [3, 4, 10], [4, 5, 15]],
                [5, 1, 2, 20, 20, 3],
                48,
            ),
            (
                25,
                [[0, 1, 10], [1, 2, 10], [2, 5, 10], [0, 3, 1], [3, 4, 10], [4, 5, 15]],
                [5, 1, 2, 20, 20, 3],
                -1,
            ),
        ]:
            output = self.minCost(maxTime, edges, passingFees)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
