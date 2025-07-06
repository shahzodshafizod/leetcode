from typing import List
import unittest

# https://leetcode.com/problems/find-closest-node-to-given-two-nodes/

# python3 -m unittest graphs/2359-find-closest-node-to-given-two-nodes.py


class Solution(unittest.TestCase):
    # # Approach #1: Depth-First Search
    # # Time: O(3n) = O(n)
    # # Space: O(2n) = O(n)
    # def closestMeetingNode(self, edges: List[int], node1: int, node2: int) -> int:
    #     n = len(edges)

    #     dist1, dist = [n+1] * n, 0
    #     while node1 != -1 and dist1[node1] > dist:
    #         dist1[node1] = dist
    #         node1 = edges[node1]
    #         dist += 1

    #     dist2, dist = [n+1] * n, 0
    #     while node2 != -1 and dist2[node2] > dist:
    #         dist2[node2] = dist
    #         node2 = edges[node2]
    #         dist += 1

    #     node, dist = -1, n+1
    #     for i in range(n):
    #         d = max(dist1[i], dist2[i])
    #         if d < dist:
    #             dist = d
    #             node = i
    #     return node

    # Approach #2: Depth-First Search (Optimized)
    # Time: O(n)
    # Space: O(n)
    def closestMeetingNode(self, edges: List[int], node1: int, node2: int) -> int:
        class Node:
            def __init__(self, node: int, mask: int):
                self.node = node
                self.mask = mask

        n = len(edges)
        seen = [0] * (n + 1)
        n1, n2 = Node(node1, 1), Node(node2, 2)
        while n1.node != -1 or n2.node != -1:
            if n1.node > n2.node:
                n1, n2 = n2, n1

            if seen[n1.node + 1] & n1.mask:
                n1.node = -1
            if seen[n2.node + 1] & n2.mask:
                n2.node = -1

            seen[n1.node + 1] |= n1.mask
            seen[n2.node + 1] |= n2.mask

            if seen[n1.node + 1] == 3:
                return n1.node
            if seen[n2.node + 1] == 3:
                return n2.node

            if n1.node != -1:
                n1.node = edges[n1.node]
            if n2.node != -1:
                n2.node = edges[n2.node]
        return -1

    def test(self):
        for edges, node1, node2, expected in [
            ([2, 2, 3, -1], 0, 1, 2),
            ([1, 2, -1], 0, 2, 2),
            ([-1, -1], 0, 0, 0),
            ([2, 0, 0], 2, 0, 0),
        ]:
            output = self.closestMeetingNode(edges, node1, node2)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
