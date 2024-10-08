from typing import Optional
import unittest
from pkg.tree import TreeNode, create_tree

# https://leetcode.com/problems/serialize-and-deserialize-binary-tree/

# python3 -m unittest design/0297-serialize-and-deserialize-binary-tree.py

class Codec:
    # Encodes a tree to a single string.
    def serialize(self, root: Optional[TreeNode]) -> str:
        if not root:
            return "N"
        return ','.join([
            str(root.val),
            self.serialize(root.left),
            self.serialize(root.right),
        ])

    # Decodes your encoded data to tree.
    def deserialize(self, data: str) -> Optional[TreeNode]:
        data = data.split(',')
        def dfs(idx: int) -> tuple[TreeNode, int]:
            if data[idx] == 'N':
                return None, idx+1
            node = TreeNode(val=int(data[idx]))
            idx += 1 # skipping current val
            node.left, idx = dfs(idx)
            node.right, idx = dfs(idx)
            return node, idx
        root, _ = dfs(0)
        return root

class TestCodec(unittest.TestCase):
    def test(self) -> None:
        ser = Codec()
        deser = Codec()
        for root, expected in [
            ([1,2,3,None,None,4,5], [1,2,3,None,None,4,5]),
            ([], []),
            ([1,2,3,None,None,4,5,6,7], [1,2,3,None,None,4,5,6,7]),
            ([1], [1]),
            ([1,-1000], [1,-1000]),
        ]:
            root = create_tree(root)
            expected = create_tree(expected)
            output = deser.deserialize(ser.serialize(root))
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")

# Your Codec object will be instantiated and called as such:
# ser = Codec()
# deser = Codec()
# ans = deser.deserialize(ser.serialize(root))
