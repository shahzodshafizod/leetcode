from typing import Optional
import unittest
from pkg.listnode import ListNode, create_linked_list

# https://leetcode.com/problems/remove-duplicates-from-sorted-list/

# python3 -m unittest linkedlists/0083-remove-duplicates-from-sorted-list.py


class Solution(unittest.TestCase):
    def deleteDuplicates(self, head: Optional[ListNode]) -> Optional[ListNode]:
        dummy = ListNode(nxt=head, val=-10000)
        node = dummy
        while node.next:
            if node.val == node.next.val:
                node.next = node.next.next
            else:
                node = node.next
        return dummy.next

    def test(self):
        for head, expected in [
            ([1, 1, 2], [1, 2]),
            ([1, 1, 2, 3, 3], [1, 2, 3]),
            ([0, 0, 0, 0, 0], [0]),
        ]:
            head = create_linked_list(head)
            expected = create_linked_list(expected)
            output = self.deleteDuplicates(head)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
