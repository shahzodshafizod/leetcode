import unittest
from typing import Optional
from pkg.listnode import ListNode, create_linked_list

# https://leetcode.com/problems/convert-binary-number-in-a-linked-list-to-integer/

# python3 -m unittest linkedlists/1290-convert-binary-number-in-a-linked-list-to-integer.py


class Solution(unittest.TestCase):
    def getDecimalValue(self, head: Optional[ListNode]) -> int:
        value, node = 0, head
        while node:
            value = (value << 1) | node.val
            node = node.next
        return value

    def test(self):
        for head, expected in [
            ([1, 0, 1], 5),
            ([0], 0),
        ]:
            head = create_linked_list(head)
            output = self.getDecimalValue(head)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
