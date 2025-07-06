from typing import Optional, List
import unittest
from pkg.tree import TreeNode, create_tree

# https://leetcode.com/problems/find-largest-value-in-each-tree-row/

# python3 -m unittest trees/0515-find-largest-value-in-each-tree-row.py


class Solution(unittest.TestCase):
    # # Approach: Depth-First Search
    # # Time: O(n), n=# of nodes
    # # Space: O(h), h=height of tree
    # def largestValues(self, root: Optional[TreeNode]) -> List[int]:
    #     largest = []
    #     def dfs(node: Optional[TreeNode], level: int):
    #         if not node: return
    #         if len(largest) == level:
    #             largest.append(node.val)
    #         else:
    #             largest[level] = max(largest[level], node.val)
    #         dfs(node.left, level+1)
    #         dfs(node.right, level+1)
    #     dfs(root, 0)
    #     return largest

    # Approach: Breadth-First Search
    # Time: O(n), n=# of nodes
    # Space: O(w), w=width of tree
    def largestValues(self, root: Optional[TreeNode]) -> List[int]:
        if not root:
            return []
        largest = []
        queue = [root]
        while queue:
            nextq = []
            largest.append(float("-inf"))
            for node in queue:
                largest[-1] = max(largest[-1], node.val)
                if node.left:
                    nextq.append(node.left)
                if node.right:
                    nextq.append(node.right)
            queue = nextq
        return largest

    def test(self):
        for root, expected in [
            ([], []),
            ([1, 2, 3], [1, 3]),
            ([95962, None], [95962]),
            ([1, 3, 2, 5, 3, None, 9], [1, 3, 9]),
            ([-1395, 32448], [-1395, 32448]),
            (
                [-38429, None, -71908, -29625, -46622, 7696, -68937, -7431, -19170, 69587, 45237],
                [-38429, -71908, -29625, 7696, 69587],
            ),
            # ([-23351,-99006,None,-63508,86343,-66012,54651,-38287,2762,39604,None,-33734,2587,66356,-99032,50112,37759,59078,-89953,44359,-73181,-18323], [-23351,-99006,86343,54651,66356,59078]),
        ]:
            root = create_tree(root)
            output = self.largestValues(root)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
