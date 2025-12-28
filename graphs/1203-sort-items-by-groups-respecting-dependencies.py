from typing import List, Tuple
from collections import deque
import unittest

# https://leetcode.com/problems/sort-items-by-groups-respecting-dependencies/

# python3 -m unittest graphs/1203-sort-items-by-groups-respecting-dependencies.py


class Solution(unittest.TestCase):
    # Approach: Dual Topological Sort with Global Item Ordering
    # Sort all items globally, then sort groups, then merge
    # This avoids repeated per-group sorting overhead
    # N = number of items, M = number of groups, E = total edges
    # Time: O(N + M + E) - single topo sort for items + groups + bucketing
    # Space: O(N + M) for graphs, indegrees, and result buckets
    def sortItems(self, n: int, m: int, group: List[int], beforeItems: List[List[int]]) -> List[int]:
        for i in range(n):
            if group[i] == -1:
                group[i] = m
                m += 1

        item_graph: List[List[int]] = [[] for _ in range(n)]
        item_indeg: List[int] = [0] * n

        group_graph: List[List[int]] = [[] for _ in range(m)]
        group_indeg = [0] * m

        for i in range(n):
            for prev in beforeItems[i]:
                item_graph[prev].append(i)
                item_indeg[i] += 1

                if group[prev] != group[i]:  # cross-group dependency
                    group_graph[group[prev]].append(group[i])
                    group_indeg[group[i]] += 1

        # Topo sort function
        def topo_sort(graph: List[List[int]], indeg: List[int], total: int) -> List[int]:
            q = deque([i for i in range(total) if indeg[i] == 0])
            result: List[int] = []
            while q:
                x = q.popleft()
                result.append(x)
                for y in graph[x]:
                    indeg[y] -= 1
                    if indeg[y] == 0:
                        q.append(y)
            return result if len(result) == total else []

        # 1) Sort groups
        group_order = topo_sort(group_graph, group_indeg[:], m)
        if not group_order:
            return []

        # 2) Sort items
        item_order = topo_sort(item_graph, item_indeg[:], n)
        if not item_order:
            return []

        # Put items into buckets *by group* following item_order
        buckets: List[List[int]] = [[] for _ in range(m)]
        for item in item_order:
            buckets[group[item]].append(item)

        # Follow group ordering to build final answer
        ans: List[int] = []
        for g in group_order:
            ans.extend(buckets[g])

        return ans

    def test(self):
        test_cases: List[Tuple[int, int, List[int], List[List[int]], List[int]]] = [
            (8, 2, [-1, -1, 1, 0, 0, 1, 0, -1], [[], [6], [5], [6], [3, 6], [], [], []], [6, 3, 4, 5, 2, 0, 7, 1]),
            (8, 2, [-1, -1, 1, 0, 0, 1, 0, -1], [[], [6], [5], [6], [3], [], [4], []], []),
        ]
        for n, m, group, beforeItems, expected in test_cases:
            output = self.sortItems(n, m, group, beforeItems)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
