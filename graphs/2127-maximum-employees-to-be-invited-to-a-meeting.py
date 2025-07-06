from collections import deque
from typing import List
import unittest

# https://leetcode.com/problems/maximum-employees-to-be-invited-to-a-meeting/

# python3 -m unittest graphs/2127-maximum-employees-to-be-invited-to-a-meeting.py


class Solution(unittest.TestCase):
    # Approach: Topological Sort to Reduce Non-Cyclic Nodes
    # Time: O(n)
    # Space: O(n)
    def maximumInvitations(self, favorite: List[int]) -> int:
        n = len(favorite)
        indegree = [0] * n
        for empl in range(n):
            indegree[favorite[empl]] += 1
        queue = deque()
        for empl in range(n):
            if indegree[empl] == 0:
                queue.append(empl)
        depth = [1] * n
        while queue:
            curr_empl = queue.popleft()
            next_empl = favorite[curr_empl]
            depth[next_empl] = max(depth[next_empl], depth[curr_empl] + 1)
            indegree[next_empl] -= 1
            if indegree[next_empl] == 0:
                queue.append(next_empl)
        longest_cycle = 0
        two_len_cycle = 0
        for empl in range(n):
            if indegree[empl] == 0:
                continue
            cycle_length = 0
            tmp_empl = empl
            while indegree[tmp_empl] != 0:
                indegree[tmp_empl] -= 1
                cycle_length += 1
                tmp_empl = favorite[tmp_empl]
            if cycle_length == 2:
                two_len_cycle += depth[empl] + depth[favorite[empl]]
            else:
                longest_cycle = max(longest_cycle, cycle_length)
        return max(longest_cycle, two_len_cycle)

    def test(self):
        for favorite, expected in [
            ([2, 2, 1, 2], 3),
            ([1, 2, 0], 3),
            ([3, 0, 1, 4, 1], 4),
        ]:
            output = self.maximumInvitations(favorite)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
