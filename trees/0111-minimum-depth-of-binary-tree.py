from collections import deque
from typing import Optional
import unittest
from pkg.tree import TreeNode, create_tree

# https://leetcode.com/problems/minimum-depth-of-binary-tree/

# python3 -m unittest trees/0111-minimum-depth-of-binary-tree.py


class Solution(unittest.TestCase):
    # # Approach: Depth-First Search
    # # Time: O(n), n=# of nodes
    # # Space: O(d), d=depth of tree
    # def minDepth(self, root: Optional[TreeNode]) -> int:
    #     if not root:
    #         return 0
    #     left = self.minDepth(root.left)
    #     right = self.minDepth(root.right)
    #     if left == 0 or right == 0:
    #         return left + right + 1
    #     return min(left, right) + 1

    # Approach: Beadth-First Search
    # Time: O(n), n=# of nodes
    # Space: O(w), w=width of tree
    def minDepth(self, root: Optional[TreeNode]) -> int:
        queue = deque()
        if root:
            queue.append(root)
        depth = 0
        while queue:
            depth += 1
            for _ in range(len(queue)):
                node = queue.popleft()
                if not node.left and not node.right:
                    return depth
                if node.left:
                    queue.append(node.left)
                if node.right:
                    queue.append(node.right)
        return 0

    def testMinDepth(self) -> None:
        for root, expected in [
            ([], 0),
            ([2], 1),
            ([3, 9, 20, None, None, 15, 7], 2),
            ([2, None, 3, None, 4, None, 5, None, 6], 5),
        ]:
            root = create_tree(root)
            output = self.minDepth(root)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
