from typing import List
import unittest

# https://leetcode.com/problems/find-unique-binary-string/

# python3 -m unittest tries/1980-find-unique-binary-string.py

class TrieNode:
    def __init__(self):
        self.children = {}
        self.count = 0

    def insert(self, s: str):
        curr = self
        for c in s:
            if c not in curr.children:
                curr.children[c] = TrieNode()
            curr = curr.children[c]
            curr.count += 1

class Solution(unittest.TestCase):
    # # Approach: Trie
    # # Time: O(nn)
    # # Space: O(nn)
    # def findDifferentBinaryString(self, nums: List[str]) -> str:
    #     root = TrieNode()
    #     for num in nums:
    #         root.insert(num)
    #     n = len(nums)
    #     num = [''] * n
    #     curr = root
    #     partition_len = 1 << n
    #     for idx in range(n):
    #         partition_len //= 2
    #         for c in "01":
    #             if c not in curr.children:
    #                 curr.children[c] = TrieNode()
    #             if curr.children[c].count < partition_len:
    #                 num[idx] = c
    #                 curr = curr.children[c]
    #                 break
    #     return "".join(num)

    # Approach: Cantor's Diagonal Argument
    # Time: O(n)
    # Space: O(1)
    def findDifferentBinaryString(self, nums: List[str]) -> str:
        n = len(nums)
        num = [''] * n
        for i in range(n):
            num[i] = '0' if nums[i][i] == '1' else '1'
        return "".join(num)

    def test(self):
        for nums, expected in [
            (["01","10"], "11"),
            (["00","01"], "10"),
            (["111","011","001"], "000"),
        ]:
            output = self.findDifferentBinaryString(nums)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
