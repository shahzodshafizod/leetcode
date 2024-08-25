from typing import List, Optional

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
