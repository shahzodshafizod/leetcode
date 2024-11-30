from collections import defaultdict
from typing import List
import unittest

# https://leetcode.com/problems/valid-arrangement-of-pairs/

# python3 -m unittest graphs/2097-valid-arrangement-of-pairs.py

"""
Eulerian paths have a couple of conditions:
	- In an undirected graph, either all nodes have an even degree,
		or exactly two have an odd degree.
	- In a directed graph (which is what we have here), we need to check if:
		- Each node’s degree matches its degree.
		- Or, exactly one node has one more outgoing edge
			(degree = degree + 1), which indicates our starting point.
"""

class Solution(unittest.TestCase):
    # # Approach 1: Eulerian Path (Recursive)
    # # Time: O(V+E), V=# of vertices(nodes), E=# of edges
    # # Space: O(V+E)
    # def validArrangement(self, pairs: List[List[int]]) -> List[List[int]]:
    #     graph = defaultdict(lambda: [])
    #     degree = defaultdict(int)
    #     for start, end in pairs:
    #         graph[start].append(end)
    #         degree[start] += 1
    #         degree[end] -= 1
    #     start = pairs[0][0]
    #     for node in degree.keys():
    #         if degree[node] == 1:
    #             start = node
    #             break
    #     order = []
    #     def visit(node: int) -> None:
    #         while graph[node]:
    #             visit(graph[node].pop())
    #         order.append(node)
    #     visit(start)
    #     return [[order[i+1], order[i]] for i in range(len(order)-2,-1,-1)]

    # Approach 2: Hierholzer's Algorithm (Iterative)
    # Time: O(V+E), V=# of vertices(nodes), E=# of edges
    # Space: O(V+E)
    def validArrangement(self, pairs: List[List[int]]) -> List[List[int]]:
        graph = defaultdict(lambda: [])
        degree = defaultdict(int)
        for start, end in pairs:
            graph[start].append(end)
            degree[start] += 1
            degree[end] -= 1
        start = pairs[0][0]
        for node in degree.keys():
            if degree[node] == 1:
                start = node
                break
        stack = [start]
        order = []
        while stack:
            node = stack[-1]
            if graph[node]:
                stack.append(graph[node].pop())
            else:
                order.append(stack.pop())
        return [[order[i], order[i-1]] for i in range(len(order)-1,0,-1)]

    def test(self) -> None:
        for pairs, expected in [
            ([[1,3],[3,2],[2,1]], [[1,3],[3,2],[2,1]]),
            ([[1,2],[1,3],[2,1]], [[1,2],[2,1],[1,3]]),
            ([[5,1],[4,5],[11,9],[9,4]], [[11,9],[9,4],[4,5],[5,1]]),
            ([[299,1],[1,2],[1,3],[2,1],[3,1]], [[299,1],[1,3],[3,1],[1,2],[2,1]]),
            ([[1,3],[3,9],[9,4],[4,1],[1,4],[4,6],[6,3],[3,1],[1,456]], [[1,4],[4,6],[6,3],[3,1],[1,3],[3,9],[9,4],[4,1],[1,456]]),
            # ([[299,1],[1,2],[1,3],[3,5],[5,8],[8,3],[3,7],[7,6],[6,9],[9,3],[3,6],[6,7],[7,3],[3,8],[8,5],[5,3],[3,1],[2,1]], [[299,1],[1,3],[3,8],[8,5],[5,3],[3,6],[6,7],[7,3],[3,7],[7,6],[6,9],[9,3],[3,5],[5,8],[8,3],[3,1],[1,2],[2,1]]),
            # ([[5,13],[10,6],[11,3],[15,19],[16,19],[1,10],[19,11],[4,16],[19,9],[5,11],[5,6],[13,5],[13,9],[9,15],[11,16],[6,9],[9,13],[3,1],[16,5],[6,5]], [[4,16],[16,5],[5,6],[6,5],[5,11],[11,16],[16,19],[19,9],[9,13],[13,5],[5,13],[13,9],[9,15],[15,19],[19,11],[11,3],[3,1],[1,10],[10,6],[6,9]]),
        ]:
            output = self.validArrangement(pairs)
            self.assertListEqual(expected, output, f"expected: {expected}, output: {output}")
