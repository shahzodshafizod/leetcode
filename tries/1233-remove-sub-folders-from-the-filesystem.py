from typing import List
import unittest

# https://leetcode.com/problems/remove-sub-folders-from-the-filesystem/

# python3 -m unittest tries/1233-remove-sub-folders-from-the-filesystem.py


class Solution(unittest.TestCase):
    # # Approach #1: Brute-Force (Stack)
    # # Time: O(nll), n=len(folder), l=len(folder[i])
    # # Space: O(nl)
    # def removeSubfolders(self, folder: List[str]) -> List[str]:
    #     folder.sort()
    #     stack = []
    #     for f in folder:
    #         if not stack or not f.startswith(stack[-1] + "/"):
    #             stack.append(f)
    #     return stack

    # Approach #2: Tries
    # Time: O(nl), n=len(folder), l=len(folder[i])
    # Space: O(nl)
    def removeSubfolders(self, folder: List[str]) -> List[str]:
        folder.sort()
        root, res = {}, []
        for f in folder:
            curr = root
            for word in f.split('/'):
                if "$" in curr:
                    break
                if word not in curr:
                    curr[word] = {}
                curr = curr[word]
            if "$" not in curr:
                res.append(f)
                curr["$"] = {}
        return res

    def test(self):
        for folder, expected in [
            (["/a", "/a/b", "/c/d", "/c/d/e", "/c/f"], ["/a", "/c/d", "/c/f"]),
            (["/a", "/a/b/c", "/a/b/d"], ["/a"]),
            (["/a/b/c", "/a/b/ca", "/a/b/d"], ["/a/b/c", "/a/b/ca", "/a/b/d"]),
        ]:
            output = self.removeSubfolders(folder)
            expected.sort()
            output.sort()
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
