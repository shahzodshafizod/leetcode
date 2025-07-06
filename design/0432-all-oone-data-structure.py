from typing import List
import unittest

# https://leetcode.com/problems/all-oone-data-structure/

# python3 -m unittest design/0432-all-oone-data-structure.py

# # Approach: Trie
# class TrieNode:
#     def __init__(self, val: str):
#         self.val = val
#         self.count = 0
#         self.children = {}

# class AllOne:
#     def __init__(self):
#         self.root = TrieNode(0)

#     # Time: O(len(key))
#     def inc(self, key: str) -> None:
#         curr = self.root
#         curr.count += 1
#         for c in key:
#             if c not in curr.children:
#                 curr.children[c] = TrieNode(c)
#             curr = curr.children[c]
#             curr.count += 1

#     # Time: O(len(key))
#     def dec(self, key: str) -> None:
#         curr = self.root
#         curr.count -= 1
#         for c in key:
#             curr = curr.children[c]
#             curr.count -= 1

#     # Time: O(26*len(key))
#     def getKey(self, first_count, compare) -> str:
#         key = ""
#         curr = self.root
#         while curr:
#             nxt = None
#             ch = ''
#             count = first_count
#             for c, child in curr.children.items():
#                 if child.count == 0:
#                     continue
#                 if compare(child.count, count):
#                     count = child.count
#                     nxt = child
#                     ch = c
#             key += ch
#             curr = nxt
#         return key

#     def getMaxKey(self) -> str:
#         return self.getKey(0, lambda a, b: a > b)

#     def getMinKey(self) -> str:
#         return self.getKey(int(5e4), lambda a, b: a < b)


# Approach: Hash Map + Doubly Linked List
class ListNode:
    def __init__(self, key, prev: "ListNode", nxt: "ListNode"):
        self.key = key
        self.count = 0
        self.prev = prev
        self.nxt = nxt


class AllOne:
    def __init__(self):
        self.keys = {}
        self.head = ListNode("", None, None)
        self.head.count = 5e4
        self.tail = ListNode("", self.head, None)
        self.head.nxt = self.tail

    def swap(self, nodes1: List[ListNode], nodes2: List[ListNode]) -> None:
        node1, node2 = nodes1[0], nodes2[0]
        prev = node1.prev
        nxt = node2.nxt
        node1.prev, node2.prev = node2, prev
        node1.nxt, node2.nxt = nxt, node1
        prev.nxt = node2
        nxt.prev = node1

    def inc(self, key: str) -> None:
        if key not in self.keys:
            node = ListNode(key, self.tail.prev, self.tail)
            node.prev.nxt = node
            node.nxt.prev = node
            self.keys[key] = node
        node = self.keys[key]
        node.count += 1
        while node.prev.count < node.count:
            self.swap([node.prev], [node])

    def dec(self, key: str) -> None:
        node = self.keys[key]
        node.count -= 1
        if node.count == 0:
            node.prev.nxt = node.nxt
            node.nxt.prev = node.prev
            del self.keys[key]
            return
        while node.count < node.nxt.count:
            self.swap([node], [node.nxt])

    def getMaxKey(self) -> str:
        if self.head.nxt == self.tail:
            return ""
        return self.head.nxt.key

    def getMinKey(self) -> str:
        if self.tail.prev == self.head:
            return ""
        return self.tail.prev.key


class TestAllOne(unittest.TestCase):
    def test(self):
        for commands, values, expected in [
            (
                ["AllOne", "inc", "inc", "getMaxKey", "getMinKey", "inc", "getMaxKey", "getMinKey"],
                [[], ["hello"], ["hello"], [], [], ["leet"], [], []],
                [None, None, None, "hello", "hello", None, "hello", "leet"],
            ),
            (
                [
                    "AllOne",
                    "inc",
                    "inc",
                    "inc",
                    "inc",
                    "getMaxKey",
                    "inc",
                    "inc",
                    "inc",
                    "dec",
                    "inc",
                    "inc",
                    "inc",
                    "getMaxKey",
                ],
                [
                    [],
                    ["hello"],
                    ["goodbye"],
                    ["hello"],
                    ["hello"],
                    [],
                    ["leet"],
                    ["code"],
                    ["leet"],
                    ["hello"],
                    ["leet"],
                    ["code"],
                    ["code"],
                    [],
                ],
                [None, None, None, None, None, "hello", None, None, None, None, None, None, None, "leet"],
            ),
            (
                [
                    "AllOne",
                    "inc",
                    "inc",
                    "inc",
                    "inc",
                    "inc",
                    "inc",
                    "dec",
                    "dec",
                    "getMinKey",
                    "dec",
                    "getMaxKey",
                    "getMinKey",
                ],
                [[], ["a"], ["b"], ["b"], ["c"], ["c"], ["c"], ["b"], ["b"], [], ["a"], [], []],
                [None, None, None, None, None, None, None, None, None, "a", None, "c", "c"],
            ),
        ]:
            for idx, command in enumerate(commands):
                output = None
                match command:
                    case "AllOne":
                        obj = AllOne()
                    case "inc":
                        obj.inc(values[idx][0])
                    case "dec":
                        obj.dec(values[idx][0])
                    case "getMaxKey":
                        output = obj.getMaxKey()
                    case "getMinKey":
                        output = obj.getMinKey()
                self.assertEqual(expected[idx], output, f"expected: {expected[idx]}, output: {output}")


# Your AllOne object will be instantiated and called as such:
# obj = AllOne()
# obj.inc(key)
# obj.dec(key)
# param_3 = obj.getMaxKey()
# param_4 = obj.getMinKey()
