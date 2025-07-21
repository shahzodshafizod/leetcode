from collections import defaultdict
from typing import List
import unittest

# https://leetcode.com/problems/delete-duplicate-folders-in-system/

# python3 -m unittest tries/1948-delete-duplicate-folders-in-system.py


class Solution(unittest.TestCase):
    def deleteDuplicateFolder(self, paths: List[List[str]]) -> List[List[str]]:
        class Trie:
            def __init__(self):
                self.children = defaultdict(Trie)
                self.deleted = False

            def insert(self, path: List[str]):
                curr = self
                for folder in path:
                    curr = curr.children[folder]

        root = Trie()
        for path in paths:
            root.insert(path)

        seen = defaultdict(list)

        def serialize(node: Trie) -> str:
            if not node.children:
                return ""
            key = []
            for folder, child in node.children.items():
                key.append(f"{folder}({serialize(child)})")
            key = "".join(sorted(key))
            seen[key].append(node)
            return key

        serialize(root)

        for nodes in seen.values():
            if len(nodes) > 1:
                for node in nodes:
                    node.deleted = True

        res = []

        def dfs(node: Trie, path: List[str]):
            for folder, child in node.children.items():
                if not child.deleted:
                    path.append(folder)
                    res.append(path[:])
                    dfs(child, path)
                    path.pop()

        dfs(root, [])

        return res

    def test(self):
        for paths, expected in [
            ([["a"], ["c"], ["d"], ["a", "b"], ["c", "b"], ["d", "a"]], [["d"], ["d", "a"]]),
            ([["a"], ["c"], ["a", "b"], ["c", "b"], ["a", "b", "x"], ["a", "b", "x", "y"], ["w"], ["w", "y"]], [["c"], ["c", "b"], ["a"], ["a", "b"]]),
            ([["a", "b"], ["c", "d"], ["c"], ["a"]], [["c"], ["c", "d"], ["a"], ["a", "b"]]),
        ]:
            output = self.deleteDuplicateFolder(paths)
            expected.sort()
            output.sort()
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
