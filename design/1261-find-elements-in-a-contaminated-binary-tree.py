from collections import deque
from typing import Optional
from pkg.tree import TreeNode, create_tree
import unittest

# https://leetcode.com/problems/find-elements-in-a-contaminated-binary-tree/

# python3 -m unittest design/1261-find-elements-in-a-contaminated-binary-tree.py


class FindElements:

    def __init__(self, root: Optional[TreeNode]):
        self.values = set()
        root.val = 0
        queue = deque([root])
        while queue:
            node = queue.popleft()
            self.values.add(node.val)
            if node.left:
                node.left.val = node.val * 2 + 1
                queue.append(node.left)
            if node.right:
                node.right.val = node.val * 2 + 2
                queue.append(node.right)

    def find(self, target: int) -> bool:
        return target in self.values


class Solution(unittest.TestCase):
    def test(self):
        for commands, values, expected in [
            (
                ["FindElements", "find", "find"],
                [[[-1, None, -1]], [1], [2]],
                [None, False, True],
            ),
            (
                ["FindElements", "find", "find", "find"],
                [[[-1, -1, -1, -1, -1]], [1], [3], [5]],
                [None, True, True, False],
            ),
            (
                ["FindElements", "find", "find", "find", "find"],
                [[[-1, None, -1, -1, None, -1]], [2], [3], [4], [5]],
                [None, True, False, False, True],
            ),
        ]:
            for idx, command in enumerate(commands):
                output = None
                match command:
                    case "FindElements":
                        finder = FindElements(create_tree(values[idx][0]))
                    case "find":
                        output = finder.find(values[idx][0])
                self.assertEqual(expected[idx], output, f"expected: {expected[idx]}, output: {output}")
