from collections import defaultdict
from typing import List, Optional, Dict, Tuple, Any
import unittest
from pkg.tree import TreeNode, create_tree

# https://leetcode.com/problems/height-of-binary-tree-after-subtree-removal-queries/

# python3 -m unittest trees/2458-height-of-binary-tree-after-subtree-removal-queries.py


class Solution(unittest.TestCase):
    # # Approach: Brute-Force
    # # Time: O(nm), n=# of tree nodes, m=len(queries)
    # # Space: O(nm)
    # def treeQueries(self, root: Optional[TreeNode], queries: List[int]) -> List[int]:
    #     def get_height(node: Optional[TreeNode], level: int, skip: int) -> int:
    #         if not node: return level
    #         if node.val == skip: return level-1
    #         return max(
    #             get_height(node.left, level+1, skip),
    #             get_height(node.right, level+1, skip),
    #         )
    #     return [get_height(root, -1, queries[idx]) for idx in range(len(queries))]

    # # Approach: DFS + BFS
    # # Time: O(n+m), n=# of tree nodes, m=len(queries)
    # # Space: O(n+m)
    # def treeQueries(self, root: Optional[TreeNode], queries: List[int]) -> List[int]:
    #     height = defaultdict(int)
    #     def dfs(node: Optional[TreeNode]) -> int:
    #         if not node: return -1
    #         height[node.val] = 1 + max(dfs(node.left), dfs(node.right))
    #         return height[node.val]
    #     def bfs() -> None:
    #         queue = deque([root])
    #         level, max1, max2 = 0, -1, -1
    #         while queue:
    #             nextmax1, nextmax2 = level, level
    #             for _ in range(len(queue)):
    #                 node = queue.popleft()
    #                 height[node.val] += level
    #                 height[node.val] = max2 if height[node.val] == max1 else max1
    #                 for next in [node.left, node.right]:
    #                     if next:
    #                         queue.append(next)
    #                         nextmax = level + 1 + height[next.val]
    #                         if nextmax > nextmax1:
    #                             nextmax1, nextmax2 = nextmax, nextmax1
    #                         elif nextmax > nextmax2:
    #                             nextmax2 = nextmax
    #             max1, max2 = nextmax1, nextmax2
    #             level += 1
    #     dfs(root) # set down-side height
    #     bfs() # set tree height if remove current node
    #     return [height[queries[idx]] for idx in range(len(queries))]

    # Approach: Depth-First Search
    # Time: O(n+m), n=# of tree nodes, m=len(queries)
    # Space: O(n+m)
    def treeQueries(self, root: Optional[TreeNode], queries: List[int]) -> List[int]:
        heights: Dict[int, int] = defaultdict(int)
        max_height = 0

        def dfs(node: Optional[TreeNode], height: int) -> None:
            if not node:
                return
            nonlocal max_height
            heights[node.val] = max(heights[node.val], max_height)
            max_height = max(max_height, height)
            dfs(node.left, height + 1)
            dfs(node.right, height + 1)
            node.left, node.right = node.right, node.left

        dfs(root, 0)
        max_height = 0  # reset for the second traversal
        dfs(root, 0)
        return [heights[queries[idx]] for idx in range(len(queries))]

    def test(self):
        test_cases: List[Tuple[List[Any], List[int], List[int]]] = [
            ([5, 8, 9, 2, 1, 3, 7, 4, 6], [3, 2, 4, 8], [3, 2, 3, 2]),
            ([1, None, 5, 3, None, 2, 4], [3, 5, 4, 2, 4], [1, 0, 3, 3, 3]),
            ([2, None, 5, 3, 1, None, 4], [1, 5, 5, 3, 4, 5], [3, 0, 0, 2, 2, 0]),
            ([1, 3, 4, 2, None, 6, 5, None, None, None, None, None, 7], [4], [2]),
            ([2, 1, 5, None, None, 3, 6, None, 4], [1, 5, 5, 6, 4, 5], [3, 1, 1, 3, 2, 1]),
            (
                [7, 2, 9, 1, 3, 8, 12, None, None, None, 6, None, None, 10, None, 4, None, None, 11, None, 5],
                [3, 8, 9, 10, 10, 5, 6, 12, 2, 6, 1, 11],
                [4, 5, 5, 5, 5, 4, 4, 5, 4, 4, 5, 5],
            ),
            (
                [1, 2, None, 3, None, 4, None, 5, None, 6, None, 7, None, 8, None, 9, None, 10, None, 11, None, 12],
                [2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12],
                [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10],
            ),
        ]
        for root, queries, expected in test_cases:
            root = create_tree(root)
            output = self.treeQueries(root, queries)
            self.assertListEqual(expected, output, f"expected: {expected}, output: {output}")
