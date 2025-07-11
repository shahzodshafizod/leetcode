from collections import deque
from typing import List, Optional, Any


# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val: int = 0, left: "TreeNode" = None, right: "TreeNode" = None):
        self.val = val
        self.left = left
        self.right = right

    def __str__(self) -> str:
        if not self:
            return "None"
        return f"{self.val},{self.left.__str__()},{self.right.__str__()}"

    def __eq__(self, __value: object) -> bool:
        if not __value:
            return False
        return all(
            [
                self.val == __value.val,
                (not self.left and not __value.left or self.left and self.left.__eq__(__value.left)),
                (not self.right and not __value.right or self.right and self.right.__eq__(__value.right)),
            ]
        )


# def create_tree(vals: List[int], idx: int) -> Optional[TreeNode]:
#     if idx >= len(vals) or vals[idx] == None:
#         return None
#     return TreeNode(
#         val=vals[idx],
#         left=create_tree(vals, 2*idx+1),
#         right=create_tree(vals, 2*idx+2)
#     )


def create_tree(vals: List[Any]) -> Optional[TreeNode]:
    root = None
    queue = deque()
    if len(vals) > 0 and vals[0] is not None:
        root = TreeNode(vals[0])
        queue.append(root)
    curr = None
    for idx in range(1, len(vals)):
        child = None
        if vals[idx] is not None:
            child = TreeNode(vals[idx])
            queue.append(child)
        if not curr:
            curr = queue.popleft()
            curr.left = child
        else:
            curr.right = child
            curr = None
    return root


# Definition for a n-ary tree node.
class Node:
    def __init__(self, val: int = 0, children: List = None):
        self.val = val
        self.children = children or []


def create_n_ary_tree(vals: List[Any]) -> Optional[Node]:
    if len(vals) == 0 or vals[0] is None:
        return None
    root = Node(vals[0])
    queue = deque()
    queue.append(root)
    idx = 2
    while idx < len(vals):
        while queue:
            node = queue.popleft()
            while idx < len(vals) and vals[idx] is not None:
                child = Node(vals[idx])
                node.children.append(child)
                queue.append(child)
                idx += 1
            idx += 1
        idx += 1
    return root
