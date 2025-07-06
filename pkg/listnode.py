from typing import List, Optional


# Definition for singly-linked list.
class ListNode:
    def __init__(self, val: int = 0, nxt: "ListNode" = None):
        self.val = val
        self.nxt = nxt

    def __eq__(self, __value: object) -> bool:
        if not __value:
            return False
        return all(
            [
                self.val == __value.val,
                not self.nxt and not __value.nxt or self.nxt.__eq__(__value.nxt),
            ]
        )

    def __str__(self) -> str:
        curr = self
        vals = []
        while curr:
            vals.append(str(curr.val))
            curr = curr.nxt
        return ",".join(vals)


def create_linked_list(vals: List[int]) -> Optional[ListNode]:
    tail = dummy = ListNode()
    for val in vals:
        tail.nxt = ListNode(val)
        tail = tail.nxt
    return dummy.nxt
