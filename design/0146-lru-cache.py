import unittest

# https://leetcode.com/problems/lru-cache/

# python3 -m unittest design/0146-lru-cache.py

class ListNode:
    def __init__(self, key: int, val: int):
        self.key = key
        self.val = val
        self.prev = None
        self.next = None

    def unlink(self) -> None:
        self.prev.next = self.next
        self.next.prev = self.prev

class LRUCache:

    def __init__(self, capacity: int):
        self.ptrs = {}
        self.capacity = capacity
        self.head = ListNode(-1,-1)
        self.tail = ListNode(-1,-1)
        self.head.next = self.tail
        self.tail.prev = self.head

    def get(self, key: int) -> int:
        if key not in self.ptrs:
            return -1
        node = self.ptrs[key]
        node.unlink()
        self.insert(node)
        return node.val

    def put(self, key: int, value: int) -> None:
        if key in self.ptrs:
            node = self.ptrs[key]
            node.val = value
            node.unlink()
        else:
            node = ListNode(key, value)
            self.ptrs[key] = node
            self.capacity -= 1
        self.insert(node)
        if self.capacity < 0:
            node = self.tail.prev
            node.unlink()
            del self.ptrs[node.key]
            self.capacity += 1

    def insert(self, node) -> None:
        node.prev = self.head
        node.next = self.head.next
        node.prev.next = node
        node.next.prev = node

class Solution(unittest.TestCase):
    def test(self):
        for commands, values, expected in [
            (
                ["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"],
                [[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]],
                [None, None, None, 1, None, -1, None, -1, 3, 4],
            ),
            (
                ["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"],
                [[2], [1, 0], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]],
                [None, None, None, 0, None, -1, None, -1, 3, 4],
            ),
            (
                ["LRUCache", "put", "put", "put", "put", "get", "get"],
                [[2], [2, 1], [1, 1], [2, 3], [4, 1], [1], [2]],
                [None, None, None, None, None, -1, 3],
            ),
            (
                ["LRUCache","get","put","get","put","put","get","get"],
                [[2],[2],[2,6],[1],[1,5],[1,2],[1],[2]],
                [None,-1,None,-1,None,None,2,6],
            ),
        ]:
            for idx, command in enumerate(commands):
                output = None
                match command:
                    case "LRUCache":
                        lru = LRUCache(values[idx][0])
                    case "get":
                        output = lru.get(values[idx][0])
                    case "put":
                        lru.put(values[idx][0], values[idx][1])
                self.assertEqual(expected[idx], output, f"expected: {expected[idx]}, output: {output}")

# Your LRUCache object will be instantiated and called as such:
# obj = LRUCache(capacity)
# param_1 = obj.get(key)
# obj.put(key,value)
