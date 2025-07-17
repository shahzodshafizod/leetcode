from typing import Optional
import unittest
from pkg.listnode import ListNode, create_linked_list

# https://leetcode.com/problems/remove-linked-list-elements/

# python3 -m unittest linkedlists/0203-remove-linked-list-elements.py


class Solution(unittest.TestCase):
    # # Approach: Recursive
    # # Time: O(n)
    # # Space: O(n)
    # def removeElements(self, head: Optional[ListNode], val: int) -> Optional[ListNode]:
    #     if not head: return None
    #     head.next = self.removeElements(head.next, val)
    #     return head.next if head.val == val else head

    # Approach: Iterative
    # Time: O(n)
    # Space: O(1)
    def removeElements(self, head: Optional[ListNode], val: int) -> Optional[ListNode]:
        node = dummy = ListNode(nxt=head)
        while node.next:
            if node.next.val == val:
                node.next = node.next.next
            else:
                node = node.next
        return dummy.next

    def test(self):
        for head, val, expected in [
            ([1, 2, 6, 3, 4, 5, 6], 6, [1, 2, 3, 4, 5]),
            ([], 1, []),
            ([7, 7, 7, 7], 7, []),
            ([1, 2, 2, 1], 2, [1, 1]),
            ([1], 2, [1]),
        ]:
            head = create_linked_list(head)
            output = self.removeElements(head, val)
            expected = create_linked_list(expected)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
