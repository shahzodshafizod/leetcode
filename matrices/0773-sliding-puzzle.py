from collections import deque
from typing import List
import unittest

# https://leetcode.com/problems/sliding-puzzle/

# python3 -m unittest matrices/0773-sliding-puzzle.py

class Solution(unittest.TestCase):
    # # Approach: Depth-First Search
    # # Time: O((m⋅n)!×(m⋅n)^2)
    # # Space: O((m⋅n)!)
    # def slidingPuzzle(self, board: List[List[int]]) -> int:
    #     neighbors = [[1,3],[0,2,4],[1,5],[0,4],[1,3,5],[2,4]]
    #     dist = {}
    #     def dfs(state: List[str], curr: int, moves: int) -> None:
    #         key = "".join(state)
    #         if key in dist and dist[key] <= moves:
    #             return
    #         dist[key] = moves
    #         for next in neighbors[curr]:
    #             state[curr], state[next] = state[next], state[curr]
    #             dfs(state, next, moves+1)
    #             state[curr], state[next] = state[next], state[curr]
    #     state = [str(val) for row in board for val in row]
    #     dfs(state, state.index('0'), 0)
    #     return dist.get("123450", -1)

    # Approach: Breadth-First Search
    # Time: O((m⋅n)!×(m⋅n))
    # Space: O((m⋅n)!)
    def slidingPuzzle(self, board: List[List[int]]) -> int:
        neighbors = [[1,3],[0,2,4],[1,5],[0,4],[1,3,5],[2,4]]
        state = "".join([str(val) for row in board for val in row])
        queue = deque([(state, state.index("0"), 0)])
        visited = set([state])
        while queue:
            state, curr, length = queue.popleft()
            if state == "123450":
                return length
            b = list(state)
            for next in neighbors[curr]:
                b[curr], b[next] = b[next], b[curr]
                new_state = "".join(b)
                if new_state not in visited:
                    queue.append((new_state, next, length+1))
                    visited.add(new_state)
                b[curr], b[next] = b[next], b[curr]
        return -1

    def test(self):
        for board, expected in [
            ([[1,2,3],[4,0,5]], 1),
            ([[1,2,3],[5,4,0]], -1),
            ([[4,1,2],[5,0,3]], 5),
            ([[1,3,4],[0,2,5]], 14),
            ([[3,2,4],[1,5,0]], 14),
            ([[2,4,1],[5,3,0]], 12),
            ([[0,5,2],[4,3,1]], 15),
            ([[4,2,0],[5,1,3]], 7),
            ([[3,0,1],[4,2,5]], -1),
            ([[3,2,1],[4,0,5]], -1),
            ([[1,2,3],[4,5,0]], 0),
        ]:
            output = self.slidingPuzzle(board)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
