from typing import List, Optional

# Definition for singly-linked list.
class ListNode:
    def __init__(self, val: int = 0, next = None):
        self.val = val
        self.next = next
    
    def __eq__(self, __value: object) -> bool:
        if not __value: return False
        return all([
            self.val == __value.val,
            self.next and self.next.__eq__(__value.next),
        ])
    
    def __str__(self) -> str:
        curr = self
        vals = []
        while curr:
            vals.append(str(curr.val))
            curr = curr.next
        return ",".join(vals)

def create_linked_list(vals: List[int]) -> Optional[ListNode]:
    tail = dummy = ListNode()
    for val in vals:
        tail.next = ListNode(val)
        tail = tail.next
    return dummy.next
