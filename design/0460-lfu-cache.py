from collections import defaultdict
import unittest

# https://leetcode.com/problems/lfu-cache/

# python3 -m unittest design/0460-lfu-cache.py


class Node:

    def __init__(self, key: int = 0, val: int = 0):
        self.key = key
        self.val = val
        self.cnt = 1
        self.prev = self.next = None

    def unlink(self) -> None:
        if self.prev:
            self.prev.next = self.next
        if self.next:
            self.next.prev = self.prev


class LRUCache:

    def __init__(self):
        self.head = Node()
        self.tail = Node()
        self.head.next = self.tail
        self.tail.prev = self.head

    def insert(self, node: Node) -> None:
        node.prev = self.head
        node.next = self.head.next
        node.prev.next = node
        node.next.prev = node

    def pop(self) -> Node:
        node = self.tail.prev
        node.unlink()
        return node

    def empty(self) -> bool:
        return self.head.next == self.tail


class LFUCache:

    def __init__(self, capacity: int):
        self.capacity = capacity
        self.nodes = {}
        self.buckets = defaultdict(LRUCache)
        self.min_cnt = 1

    def get(self, key: int) -> int:
        node = self.nodes.get(key, None)
        if not node:
            return -1
        node.unlink()  # removes from the current counts[node.cnt]
        if self.buckets[node.cnt].empty() and node.cnt == self.min_cnt:
            self.min_cnt += 1
        node.cnt += 1
        self.buckets[node.cnt].insert(node)
        return node.val

    def put(self, key: int, value: int) -> None:
        if self.get(key) != -1:
            self.nodes[key].val = value
            return
        if self.capacity == 0:
            node = self.buckets[self.min_cnt].pop()
            del self.nodes[node.key]
            self.capacity += 1
        self.nodes[key] = node = Node(key, value)
        self.buckets[node.cnt].insert(node)
        self.min_cnt = 1
        self.capacity -= 1


class Solution(unittest.TestCase):
    def test(self):
        for commands, values, expected in [
            (
                ["LFUCache", "put", "put", "get", "put", "get", "get", "put", "get", "get", "get"],
                [[2], [1, 1], [2, 2], [1], [3, 3], [2], [3], [4, 4], [1], [3], [4]],
                [None, None, None, 1, None, -1, 3, None, -1, 3, 4],
            ),
        ]:
            for idx, command in enumerate(commands):
                output = None
                match command:
                    case "LFUCache":
                        lfu = LFUCache(values[idx][0])
                    case "get":
                        output = lfu.get(values[idx][0])
                    case "put":
                        lfu.put(values[idx][0], values[idx][1])
                self.assertEqual(expected[idx], output, f"expected: {expected[idx]}, output: {output}")


# Your LFUCache object will be instantiated and called as such:
# obj = LFUCache(capacity)
# param_1 = obj.get(key)
# obj.put(key,value)
