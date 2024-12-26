from typing import Optional
import unittest
from pkg.tree import TreeNode, create_tree

# https://leetcode.com/problems/merge-two-binary-trees/

# python3 -m unittest trees/0617-merge-two-binary-trees.py

class Solution(unittest.TestCase):
    # # Approach: Depth-First Search
    # # Time: O(max(n1, n2)), n1=# of tree1 node, n2=# of tree2 nodes
    # # Space: O(max(h1, h2)), h1=height of tree1, h2=height of tree2
    # def mergeTrees(self, root1: Optional[TreeNode], root2: Optional[TreeNode]) -> Optional[TreeNode]:
    #     # if root1 and root2:
    #     #     return TreeNode(
    #     #         val=root1.val+root2.val,
    #     #         left=self.mergeTrees(root1.left, root2.left),
    #     #         right=self.mergeTrees(root1.right, root2.right)
    #     #     )
    #     # return root1 if root1 else root2 if root2 else None

    #     if not root1: return root2
    #     if not root2: return root1
    #     root1.val += root2.val
    #     root1.left = self.mergeTrees(root1.left, root2.left)
    #     root1.right = self.mergeTrees(root1.right, root2.right)
    #     return root1

    # Approach: Breadth-First Search
    # Time: O(max(n1, n2)), n1=# of tree1 node, n2=# of tree2 nodes
    # Space: O(max(w1, w2)), w1=width of tree1, w2=width of tree2
    def mergeTrees(self, root1: Optional[TreeNode], root2: Optional[TreeNode]) -> Optional[TreeNode]:
        if not root1: return root2
        stack = [(root1, root2)]
        while stack:
            node1, node2 = stack.pop()

            if not node1 or not node2: continue

            node1.val += node2.val

            if not node1.left: node1.left = node2.left
            else: stack.append((node1.left, node2.left))

            if not node1.right: node1.right = node2.right
            else: stack.append((node1.right, node2.right))
        return root1

    def test(self):
        for root1, root2, expected in [
            ([1,3,2,5], [2,1,3,None,4,None,7], [3,4,5,5,4,None,7]),
            ([1], [1,2], [2,2]),
        ]:
            root1 = create_tree(root1)
            root2 = create_tree(root2)
            output = self.mergeTrees(root1, root2)
            expected = create_tree(expected)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
