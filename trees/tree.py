from collections import deque
from typing import List, Optional

# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val: int = 0, left = None, right = None) -> None:
        self.val = val
        self.left = left
        self.right = right

def create_tree(vals: List[int], idx: int) -> Optional[TreeNode]:
    if idx >= len(vals) or vals[idx] == None:
        return None
    return TreeNode(
        val=vals[idx],
        left=create_tree(vals, 2*idx+1),
        right=create_tree(vals, 2*idx+2)
    )

# Definition for a n-ary tree node.
class Node:
    def __init__(self, val: int = 0, children: List = None) -> None:
        self.val = val
        self.children = children or []

def create_n_ary_tree(vals: List[int]) -> Optional[Node]:
    if len(vals) == 0 or vals[0] == None:
        return None
    root = Node(vals[0])
    queue = deque()
    queue.append(root)
    idx = 2
    while idx < len(vals):
        while queue:
            node = queue.popleft()
            while idx < len(vals) and vals[idx] != None:
                child = Node(vals[idx])
                node.children.append(child)
                queue.append(child)
                idx += 1
            idx += 1
        idx += 1
    return root
