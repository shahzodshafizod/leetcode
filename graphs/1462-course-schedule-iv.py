from typing import List
from collections import deque
import unittest

# https://leetcode.com/problems/course-schedule-iv/

# python3 -m unittest graphs/1462-course-schedule-iv.py


class Solution(unittest.TestCase):
    def checkIfPrerequisite(
        self, numCourses: int, prerequisites: List[List[int]], queries: List[List[int]]
    ) -> List[bool]:
        adj_list = [[] for _ in range(numCourses)]
        indegree = [0] * numCourses
        for src, dst in prerequisites:
            adj_list[src].append(dst)
            indegree[dst] += 1
        queue = deque()
        for course in range(numCourses):
            if indegree[course] == 0:
                queue.append(course)
        reach = [[False] * numCourses for _ in range(numCourses)]
        while queue:
            src = queue.popleft()
            for dst in adj_list[src]:
                reach[dst][src] = True
                for node in range(numCourses):
                    reach[dst][node] = reach[dst][node] or reach[src][node]
                indegree[dst] -= 1
                if indegree[dst] == 0:
                    queue.append(dst)
        return [reach[v][u] for u, v in queries]

    def test(self):
        for numCourses, prerequisites, queries, expected in [
            (2, [[1, 0]], [[0, 1], [1, 0]], [False, True]),
            (2, [], [[1, 0], [0, 1]], [False, False]),
            (3, [[1, 2], [1, 0], [2, 0]], [[1, 0], [1, 2]], [True, True]),
            (3, [[1, 0], [2, 0]], [[0, 1], [2, 0]], [False, True]),
            (5, [[0, 1], [1, 2], [2, 3], [3, 4]], [[0, 4], [4, 0], [1, 3], [3, 0]], [True, False, True, False]),
            (
                4,
                [[2, 3], [2, 1], [0, 3], [0, 1]],
                [[0, 1], [0, 3], [2, 3], [3, 0], [2, 0], [0, 2]],
                [True, True, True, False, False, False],
            ),
            (
                5,
                [[4, 3], [4, 1], [4, 0], [3, 2], [3, 1], [3, 0], [2, 1], [2, 0], [1, 0]],
                [[1, 4], [4, 2], [0, 1], [4, 0], [0, 2], [1, 3], [0, 1]],
                [False, True, False, True, False, False, False],
            ),
        ]:
            output = self.checkIfPrerequisite(numCourses, prerequisites, queries)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
