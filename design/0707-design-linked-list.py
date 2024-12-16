import unittest

# https://leetcode.com/problems/design-linked-list/

# python3 -m unittest design/0707-design-linked-list.py

class ListNode:
    def __init__(self, val, prev, next):
        self.val = val
        self.prev = prev
        self.next = next

class MyLinkedList:

    def __init__(self):
        self.head = ListNode(-1, None,None)
        self.tail = ListNode(-1, self.head, None)
        self.head.next = self.tail
        self.size = 0

    def get(self, index: int) -> int:
        node = self.get_node(index)
        return node.val if node else -1

    def addAtHead(self, val: int) -> None:
        self.add_node(val, self.head, self.head.next)

    def addAtTail(self, val: int) -> None:
        self.add_node(val, self.tail.prev, self.tail)

    def addAtIndex(self, index: int, val: int) -> None:
        if index == self.size:
            self.addAtTail(val)
            return
        next = self.get_node(index)
        if not next: return
        self.add_node(val, next.prev, next)

    def add_node(self, val: int, prev: ListNode, next: ListNode) -> None:
        node = ListNode(val, prev, next)
        node.prev.next = node
        node.next.prev = node
        self.size += 1

    def deleteAtIndex(self, index: int) -> None:
        node = self.get_node(index)
        if not node: return
        node.prev.next = node.next
        node.next.prev = node.prev
        self.size -= 1

    def get_node(self, index: int) -> ListNode:
        if index < 0 or index >= self.size:
            return None
        if index > self.size//2:
            node = self.tail
            for _ in range(index, self.size):
                node = node.prev
            return node
        node = self.head.next
        for _ in range(index):
            node = node.next
        return node

class Solution(unittest.TestCase):
    def test(self):
        for commands, values, expected in [
            (
                ["MyLinkedList","addAtHead","addAtTail","addAtIndex","get","deleteAtIndex","get"],
                [[],[1],[3],[1,2],[1],[1],[1]],
                [None,None,None,None,2,None,3],
            ),
            (
                ["MyLinkedList","addAtHead","addAtHead","addAtHead","addAtIndex","deleteAtIndex","addAtHead","addAtTail","get","addAtHead","addAtIndex","addAtHead"],
                [[],[7],[2],[1],[3,0],[2],[6],[4],[4],[4],[5,0],[6]],
                [None,None,None,None,None,None,None,None,4,None,None,None],
            ),
            (
                ["MyLinkedList","addAtHead","addAtIndex","get","get","get"],
                [[],[1],[1,2],[1],[0],[2]],
                [None,None,None,2,1,-1],
            ),
            (
                ["MyLinkedList","addAtHead","get","addAtHead","addAtHead","deleteAtIndex","addAtHead","get","get","get","addAtHead","deleteAtIndex"],
                [[],[4],[1],[1],[5],[3],[7],[3],[3],[3],[1],[4]],
                [None,None,-1,None,None,None,None,4,4,4,None,None],
            ),
            (
                ["MyLinkedList","addAtHead","addAtTail","addAtIndex","get","deleteAtIndex","get"],
                [[],[1],[3],[1,2],[1],[0],[0]],
                [None,None,None,None,2,None,2],
            ),
            (
                ["MyLinkedList","addAtIndex","addAtIndex","addAtIndex","get"],
                [[],[0,10],[0,20],[1,30],[0]],
                [None,None,None,None,20],
            ),
        ]:
            for idx, command in enumerate(commands):
                output = None
                match command:
                    case "MyLinkedList":
                        ll = MyLinkedList()
                    case "get":
                        output = ll.get(values[idx][0])
                    case "addAtHead":
                        ll.addAtHead(values[idx][0])
                    case "addAtTail":
                        ll.addAtTail(values[idx][0])
                    case "addAtIndex":
                        ll.addAtIndex(values[idx][0], values[idx][1])
                    case "deleteAtIndex":
                        ll.deleteAtIndex(values[idx][0])
                self.assertEqual(expected[idx], output, f"expected: {expected[idx]}, output: {output}")

# Your MyLinkedList object will be instantiated and called as such:
# obj = MyLinkedList()
# param_1 = obj.get(index)
# obj.addAtHead(val)
# obj.addAtTail(val)
# obj.addAtIndex(index,val)
# obj.deleteAtIndex(index)
