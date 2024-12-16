from typing import List
import unittest

# https://leetcode.com/problems/find-critical-and-pseudo-critical-edges-in-minimum-spanning-tree/

# python3 -m unittest unionfinds/1489-find-critical-and-pseudo-critical-edges-in-minimum-spanning-tree.py

class UnionFind:
    def __init__(self, n: int):
        self.size = n
        self.parent = [None] * n
    def Reset(self) -> None:
        for idx in range(self.size):
            self.parent[idx] = idx
        self.weight = 0
        self.count = self.size
    def Find(self, x: int) -> int:
        if self.parent[x] != x:
            self.parent[x] = self.Find(self.parent[x])
        return self.parent[x]
    def Union(self, x: int, y: int, weight: int) -> None:
        px = self.Find(x)
        py = self.Find(y)
        if px != py:
            self.parent[max(px, py)] = min(px, py)
            self.weight += weight
            self.count -= 1

class Solution(unittest.TestCase):
    # Approach: Kruskal's Algorithm
    # Time: O(E^2LogV), E=# of edges, V=# of vertices
    # Space: O(E+V)
    def findCriticalAndPseudoCriticalEdges(self, n: int, edges: List[List[int]]) -> List[List[int]]:
        for idx, edge in enumerate(edges):
            edge.append(idx) # (node1, node2, weight, idx)
        edges.sort(key=lambda edge: edge[2])
        # Kruskal's Algorithm to determine the MST (Minimum Spanning Tree)
        mst = UnionFind(n)
        def kruskal(skip: int, pick: int) -> int: # Time: O(ElogV)
            mst.Reset()
            if pick >= 0:
                node1, node2, weight, _ = edges[pick]
                mst.Union(node1, node2, weight)
            for idx, (node1,node2,weight,_) in enumerate(edges):
                if idx != skip:
                    mst.Union(node1, node2, weight)
            return mst.weight if mst.count == 1 else float("inf")
        mst_weight = kruskal(-1, -1)
        critical, pseudoc = [], [] # pseudocritical = non-critical
        for idx in range(len(edges)): # Time: O(E)
            if kruskal(idx, -1) > mst_weight: # without curr edge
                critical.append(edges[idx][3])
            elif kruskal(-1, idx) == mst_weight: # with curr edge
                pseudoc.append(edges[idx][3])
        return [critical, pseudoc]

    def test(self) -> None:
        for n, edges, expected in [
            (2, [[0,1,3]], [[0],[]]),
            (3, [[0,1,1],[0,2,2],[1,2,3]], [[0,1],[]]),
            (4, [[0,1,1],[0,2,2],[0,3,3]], [[0,1,2],[]]),
            (4, [[0,1,1],[1,2,1],[2,3,1],[0,3,1]], [[],[0,1,2,3]]),
            (5, [[0,1,2],[0,2,2],[2,3,2],[2,4,2]], [[0,1,2,3],[]]),
            (6, [[0,1,2],[0,2,5],[2,3,5],[1,4,4],[2,5,5],[4,5,2]], [[0,5,3,2],[1,4]]),
            (4, [[0,1,1],[0,3,1],[0,2,1],[1,2,1],[1,3,1],[2,3,1]], [[],[0,1,2,3,4,5]]),
            (5, [[0,1,1],[1,2,1],[2,3,2],[0,3,2],[0,4,3],[3,4,3],[1,4,6]], [[0,1],[2,3,4,5]]),
            (6, [[0,1,1],[1,2,1],[0,2,1],[2,3,4],[3,4,2],[3,5,2],[4,5,2]], [[3],[0,1,2,4,5,6]]),
            (8, [[0,1,1],[1,2,1],[2,3,1],[0,4,1],[1,5,1],[2,6,1],[3,7,1],[4,5,1],[6,7,1]], [[1],[0,2,3,4,5,6,7,8]]),
            (6, [[0,1,4],[0,2,3],[1,3,4],[3,4,1],[3,5,2],[0,5,1],[2,3,5],[2,4,6],[0,4,4],[2,5,2]], [[3,5,4,9],[0,2]]),
        ]:
            output = self.findCriticalAndPseudoCriticalEdges(n, edges)
            self.assertListEqual(expected, output, f"expected: {expected}, output: {output}")
