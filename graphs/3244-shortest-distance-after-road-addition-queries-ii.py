from typing import List
import unittest

# https://leetcode.com/problems/shortest-distance-after-road-addition-queries-ii/

# python3 -m unittest graphs/3244-shortest-distance-after-road-addition-queries-ii.py

class Solution(unittest.TestCase):
    # # Approach: Union Find
    # # Time: O(n+q)
    # # Space: O(n)
    # def shortestDistanceAfterQueries(self, n: int, queries: List[List[int]]) -> List[int]:
    #     class UnionFind:
    #         def __init__(self, n: int) -> None:
    #             self.root = [x for x in range(n)]
    #             self.n = n
    #         def find(self, x: int) -> int:
    #             if self.root[x] != x:
    #                 self.root[x] = self.find(self.root[x])
    #             return self.root[x]
    #         def union(self, x: int, y: int) -> None:
    #             px = self.find(x)
    #             py = self.find(y)
    #             if px != py:
    #                 self.root[py] = px
    #                 self.n -= 1
    #         def size(self) -> int:
    #             return self.n
    #     uf = UnionFind(n)
    #     answer = [n-1] * len(queries)
    #     for idx, (src, dst) in enumerate(queries):
    #         src += 1
    #         while dst > src:
    #             uf.union(dst-1, dst)
    #             dst = uf.find(dst)
    #         answer[idx] = uf.size()-1
    #     return answer

    # Approach: Linked List
    # Time: O(n+q)
    # Space: O(n)
    def shortestDistanceAfterQueries(self, n: int, queries: List[List[int]]) -> List[int]:
        next = [node+1 for node in range(n)]
        distance = n-1
        answer = [None] * len(queries)
        for idx, (src, dst) in enumerate(queries):
            while next[src] < dst:
                next[src], src = dst, next[src]
                distance -= 1
            answer[idx] = distance
        return answer

    def test(self) -> None:
        for n, queries, expected in [
            (5, [[2,4],[0,2],[0,4]], [3,2,1]),
		    (4, [[0,3],[0,2]], [1,1]),
            (13, [[6,10],[1,12],[7,9]], [9,2,2]),
            (7, [[1,4],[1,3],[1,5]], [4,4,3]),
        ]:
            output = self.shortestDistanceAfterQueries(n, queries)
            self.assertListEqual(expected, output, f"expected: {expected}, output: {output}")
