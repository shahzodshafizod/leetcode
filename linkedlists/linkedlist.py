from typing import List, Optional

# Definition for singly-linked list.
class ListNode:
    def __init__(self, val: int = 0, next = None):
        self.val = val
        self.next = next
    
    def __eq__(self, __value: object) -> bool:
        if not self or not __value:
            return not self and not __value
        return self.val == __value.val and self.next.__eq__(__value.next)
    
    def __str__(self) -> str:
        curr = self
        vals = []
        while curr:
            vals.append(curr.val)
            curr = curr.next
        return ",".join([str(val) for val in vals])

def create_linked_list(vals: List[int]) -> Optional[ListNode]:
    tail = dummy = ListNode()
    for val in vals:
        tail.next = ListNode(val)
        tail = tail.next
    return dummy.next
