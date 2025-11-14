from typing import Optional
import unittest
from pkg.listnode import ListNode, create_linked_list

# https://leetcode.com/problems/intersection-of-two-linked-lists/

# python3 -m unittest linkedlists/0160-intersection-of-two-linked-lists.py


class Solution(unittest.TestCase):
    def getIntersectionNode(self, headA: ListNode, headB: ListNode) -> Optional[ListNode]:
        lenA = 0
        node, tailA = headA, None
        while node:
            lenA += 1
            tailA = node
            node = node.next
        lenB, node, tailB = 0, headB, None
        while node:
            lenB += 1
            tailB = node
            node = node.next
        if tailA != tailB:
            return None
        while lenA > lenB:
            headA = headA.next
            lenA -= 1
        while lenB > lenA:
            headB = headB.next
            lenB -= 1
        while headA != headB:
            headA = headA.next
            headB = headB.next
        return headA

    def test(self):
        for listA, listB, skipA, skipB in [
            ([4, 1, 8, 4, 5], [5, 6, 1, 8, 4, 5], 2, 3),
            ([1, 9, 1, 2, 4], [3, 2, 4], 3, 1),
            ([2, 6, 4], [1, 5], 3, 2),
        ]:
            headA = create_linked_list(listA)
            headB = create_linked_list(listB)
            nodeA = headA
            for _ in range(skipA - 2):
                if nodeA:
                    nodeA = nodeA.next
            nodeB = headB
            for _ in range(skipB - 2):
                if nodeB:
                    nodeB = nodeB.next
            expected = None
            if nodeA and nodeB:
                nodeA.next = nodeB.next
                expected = nodeA.next
            if headA and headB:
                output = self.getIntersectionNode(headA, headB)
                self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
