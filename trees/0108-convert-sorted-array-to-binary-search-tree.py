from typing import Optional, List
import unittest
from pkg.tree import TreeNode, create_tree

# https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/

# python3 -m unittest trees/0108-convert-sorted-array-to-binary-search-tree.py


class Solution(unittest.TestCase):
    # # Time: O(n), n=len(nums)
    # # Space: O(h), h=height of tree
    # def sortedArrayToBST(self, nums: List[int]) -> Optional[TreeNode]:
    #     if len(nums) == 0: return None
    #     if len(nums) == 1: return TreeNode(val=nums[0])
    #     mid = len(nums)//2
    #     return TreeNode(
    #         val=nums[mid],
    #         left=self.sortedArrayToBST(nums[:mid]),
    #         right=self.sortedArrayToBST(nums[mid+1:]),
    #     )

    # Time: O(n), n=len(nums)
    # Space: O(h), h=height of tree
    def sortedArrayToBST(self, nums: List[int]) -> Optional[TreeNode]:
        def createTree(left: int, right: int) -> Optional[TreeNode]:
            if left == right:
                return None
            if left + 1 == right:
                return TreeNode(val=nums[left])
            mid = (left + right) // 2
            return TreeNode(
                val=nums[mid],
                left=createTree(left, mid),
                right=createTree(mid + 1, right),
            )

        return createTree(0, len(nums))

    def test(self):
        for nums, expected in [
            ([-10, -3, 0, 5, 9], [0, -3, 9, -10, None, 5]),
            ([1, 3], [3, 1]),
        ]:
            output = self.sortedArrayToBST(nums)
            expected = create_tree(expected)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
