from collections import deque
from typing import List
import unittest

# https://leetcode.com/problems/divide-nodes-into-the-maximum-number-of-groups/

# python3 -m unittest graphs/2493-divide-nodes-into-the-maximum-number-of-groups.py

class Solution(unittest.TestCase):
    # Approach: Union-Find + Breadth-First Search
    # Time: O(v*(v+e)), v=# of vertices, e=# of edges
    # Space: O(v+e)
    def magnificentSets(self, n: int, edges: List[List[int]]) -> int: 
        parent = [-1] * n
        depth = [0] * n
        def find(node: int) -> int:
            while parent[node] != -1:
                node = parent[node]
            return node
        def union(x: int, y: int):
            px = find(x)
            py = find(y)
            if px == py:
                return
            if depth[px] < depth[py]:
                px, py = py, px
            parent[py] = px
            if depth[px] == depth[py]:
                depth[px] += 1
        graph = [[] for _ in range(n)]
        for edge in edges:
            a, b = edge[0]-1, edge[1]-1
            graph[a].append(b)
            graph[b].append(a)
            union(a, b)
        def bfs(src: int) -> int:
            queue = deque([src])
            node_layer = [-1] * n
            node_layer[src] = 0
            layer = 0
            while queue:
                for _ in range(len(queue)):
                    curr = queue.popleft()
                    for next in graph[curr]:
                        if node_layer[next] == -1:
                            node_layer[next] = layer + 1
                            queue.append(next)
                        elif node_layer[next] == layer:
                            return -1
                layer += 1
            return layer
        groups = {}
        for node in range(n):
            group_cnt = bfs(node)
            if group_cnt == -1:
                return -1
            root = find(node)
            groups[root] = max(groups.get(root, 0), group_cnt)
        return sum(groups.values())

    def test(self):
        for n, edges, expected in [
            (2, [[1,2]], 2),
            (30, [[16,8],[6,5]], 30),
            (3, [[1,2],[2,3],[3,1]], -1),
            (6, [[1,2],[1,4],[1,5],[2,6],[2,3],[4,6]], 4),
            (8, [[5,3],[5,8],[7,2],[5,4],[6,4],[6,3],[5,2],[6,2],[1,8],[6,8],[7,4],[7,3],[7,8],[1,4]], 4),
            (24, [[2,13],[7,3],[5,3],[21,1],[5,1],[4,13],[21,19],[7,13],[15,3],[21,22],[17,19],[23,22],[14,13]], 19),
            (14, [[1,4],[3,4],[9,8],[12,8],[11,8],[11,5],[6,4],[3,5],[1,8],[13,4],[9,5],[9,4],[12,5],[7,5],[2,4],[3,8],[1,5],[12,4],[11,4],[10,4],[14,5],[14,8]], 5),
        ]:
            output = self.magnificentSets(n, edges)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
