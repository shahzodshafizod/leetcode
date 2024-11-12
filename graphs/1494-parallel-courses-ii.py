from collections import defaultdict, deque
from functools import lru_cache
from itertools import combinations
from typing import List
import unittest

# https://leetcode.com/problems/parallel-courses-ii/

# python3 -m unittest graphs/1494-parallel-courses-ii.py

# TODO: Solve this problem by yourself

class Solution(unittest.TestCase):
    # # Top-Down Dynamic Programming
    # def minNumberOfSemesters(self, n: int, relations: List[List[int]], k: int) -> int:
    #     @lru_cache(None) # caching for faster lookups
    #     def dfs(in_degrees):
    #         # all the nodes that *can* be taken now, following both the properties
    #         nodes = [i for i in range(n) if in_degrees[i] == 0]
    #         # if all the nodes have already been considered, we have taken all the courses
    #         if len(nodes) == 0: return 0
    #         ans = float("inf")
    #         # enumerating all the possible combinations
    #         for k_nodes in combinations(nodes, min(k, len(nodes))):
    #             new_in_degrees = list(in_degrees)
    #             # updating what would happen to new_mask and new_in_degrees 
    #             # if we considered the nodes in k_nodes
    #             for node in k_nodes:
    #                 # since we know the bit is set, we un-set this bit, to mark it "considered"
    #                 new_in_degrees[node] -= 1
    #                 # updating each of the in-degrees, since the "parents" have been taken away
    #                 for child in graph[node]:
    #                     new_in_degrees[child] -= 1
    #             # the heart of recursion
    #             # note the +1!
    #             ans = min(ans, 1+dfs(tuple(new_in_degrees)))
    #         return ans
    #     # saving n and k for later use
    #     in_degrees = [0]*n
    #     # graph layout remains the same, although the in_degrees change. 
    #     # This allows us to keep graph as self.graph
    #     # instead of passing it over and over.
    #     graph = defaultdict(list)
    #     for prev_course, next_course in relations:
    #         prev_course -= 1
    #         next_course -= 1
    #         # remember, its 0-indexed now!
    #         in_degrees[next_course] += 1
    #         graph[prev_course].append(next_course)
    #     # start with all the bits set
    #     return dfs(tuple(in_degrees))

    # Bottom-Up Dynamic Programming
    def minNumberOfSemesters(self, n: int, relations: List[List[int]], k: int) -> int:
        seen = [False] * (1<<n)
        preq = [0]*n
        for prev, next in relations:
            preq[next-1] += 1 << (prev-1)
        TARGET = (1<<n)-1 # zero-based index
        queue = deque([(0,0)])
        result = n
        while queue:
            state, semesters = queue.popleft()
            if state == TARGET:
                result = semesters
                break
            nextCourses = [] # new courses to study
            for course in range(n):
                if preq[course] & state != preq[course]: # not enough prereq
                    continue
                if state & (1<<course) > 0: # already studies this course
                    continue
                nextCourses.append(course)
            for k_courses in combinations(nextCourses, min(k, len(nextCourses))):
                newState = state
                for course in k_courses:
                    newState |= 1 << course
                if not seen[newState]:
                    seen[newState] = True
                    queue.append((newState, semesters+1))
        return result

    def test(self) -> None:
        for n, relations, k, expected in [
            (11, [], 2, 6),
            (4, [[2,1],[3,1],[1,4]], 2, 3),
            (5, [[2,1],[3,1],[4,1],[1,5]], 2, 4),
            (12, [[1,2],[1,3],[7,5],[7,6],[4,8],[8,9],[9,10],[10,11],[11,12]], 2, 6),
            (13, [[12,8],[2,4],[3,7],[6,8],[11,8],[9,4],[9,7],[12,4],[11,4],[6,4],[1,4],[10,7],[10,4],[1,7],[1,8],[2,7],[8,4],[10,8],[12,7],[5,4],[3,4],[11,7],[7,4],[13,4],[9,8],[13,8]], 9, 3),
        ]:
            output = self.minNumberOfSemesters(n, relations, k)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
