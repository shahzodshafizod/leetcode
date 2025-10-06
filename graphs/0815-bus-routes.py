from collections import defaultdict, deque
from typing import List, Dict, Set
import unittest

# https://leetcode.com/problems/bus-routes/

# python3 -m unittest graphs/0815-bus-routes.py


class Solution(unittest.TestCase):
    # Approach: Breadth-First Search
    # Time: O(n*m), n=len(routes), m=len(routes[i])
    # Space: O(n*m)
    def numBusesToDestination(self, routes: List[List[int]], source: int, target: int) -> int:
        graph: Dict[int, List[int]] = defaultdict(list)  # stop -> [bus]
        for bus, stops in enumerate(routes):
            for stop in stops:
                graph[stop].append(bus)
        queue = deque([(source, 0)])  # (stop, the number of buses taken from the source until this stop)
        seen_buses = [False] * len(routes)
        seen_stops: Set[int] = set([source])
        while queue:
            curr_stop, cnt = queue.popleft()
            if curr_stop == target:
                return cnt
            for bus in graph[curr_stop]:
                if seen_buses[bus]:
                    continue
                for next_stop in routes[bus]:
                    if next_stop not in seen_stops:
                        queue.append((next_stop, cnt + 1))
                        seen_stops.add(next_stop)
                seen_buses[bus] = True
        return -1

    def test(self):
        for routes, source, target, expected in [
            ([[1, 2, 7], [3, 6, 7]], 1, 6, 2),
            ([[7, 12], [4, 5, 15], [6], [15, 19], [9, 12, 13]], 15, 12, -1),
        ]:
            output = self.numBusesToDestination(routes, source, target)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
