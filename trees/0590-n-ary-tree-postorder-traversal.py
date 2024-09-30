import unittest
from typing import List, Optional
from design.tree import Node, create_n_ary_tree

# https://leetcode.com/problems/n-ary-tree-postorder-traversal/

# python3 -m unittest trees/0590-n-ary-tree-postorder-traversal.py

class Solution(unittest.TestCase):
    # # Iterative
    # def postorder(self, root: Optional[Node]) -> List[int]:
    #     stack = []
    #     if root != None:
    #         stack.append(root)
    #     order = []
    #     while stack:
    #         node = stack.pop()
    #         order.append(node.val)
    #         for child in node.children:
    #             stack.append(child)
    #     return list(reversed(order))

    # Recursive
    def postorder(self, root: Optional[Node]) -> List[int]:
        if root == None:
            return []
        order = []
        for child in root.children:
            order += self.postorder(child)
        order.append(root.val)
        return order

    def test(self) -> None:
        for vals, expected in [
            ([1,None,3,2,4,None,5,6], [5,6,3,2,4,1]),
            ([1,None,2,3,4,5,None,None,6,7,None,8,None,9,10,None,None,11,None,12,None,13,None,None,14], [2,6,14,11,7,3,12,8,4,13,9,10,5,1]),
            ([], []),
        ]:
            root = create_n_ary_tree(vals)
            output = self.postorder(root)
            self.assertListEqual(expected, output, f"expected: {expected}, output: {output}")
