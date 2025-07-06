import unittest
from typing import List

# https://leetcode.com/problems/unique-email-addresses/

# python3 -m unittest hashes/0929-unique-email-addresses.py


class Solution(unittest.TestCase):
    # # Approach #1: Trie
    # # Time: O(n*m), n=len(emails), m=len(emails[i])
    # # Space: O(n*m)
    # def numUniqueEmails(self, emails: List[str]) -> int:
    #     class TrieNode:
    #         def __init__(self):
    #             self.children = {}
    #             self.end = False
    #     LOCAL, SKIP, DOMAIN = 0, 1, 2
    #     root = TrieNode()
    #     count = 0
    #     for email in emails:
    #         curr = root
    #         part = LOCAL
    #         for c in email:
    #             if part != DOMAIN:
    #                 match c:
    #                     case '.': continue
    #                     case '+': part = SKIP
    #                     case '@': part = DOMAIN
    #             if part == SKIP:
    #                 continue
    #             if c not in curr.children:
    #                 curr.children[c] = TrieNode()
    #             curr = curr.children[c]
    #         if not curr.end:
    #             count += 1
    #         curr.end = True
    #     return count

    # Approach #2: Hash Set
    # Time: O(n*m), n=len(emails), m=len(emails[i])
    # Space: O(n)
    def numUniqueEmails(self, emails: List[str]) -> int:
        receivers = set()
        for email in emails:
            local, domain = email.split("@")
            local = local.split("+")[0]
            local = local.replace(".", "")
            receivers.add(f"{local}@{domain}")
        return len(receivers)

    def testNumUniqueEmails(self) -> None:
        for emails, expected in [
            (["a@leetcode.com", "b@leetcode.com", "c@leetcode.com"], 3),
            (["test.email+alex@leetcode.com", "test.email.leet+alex@code.com"], 2),
            (
                ["test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"],
                2,
            ),
            # (["fg.r.u.uzj+o.pw@kziczvh.com","r.cyo.g+d.h+b.ja@tgsg.z.com","fg.r.u.uzj+o.f.d@kziczvh.com","r.cyo.g+ng.r.iq@tgsg.z.com","fg.r.u.uzj+lp.k@kziczvh.com","r.cyo.g+n.h.e+n.g@tgsg.z.com","fg.r.u.uzj+k+p.j@kziczvh.com","fg.r.u.uzj+w.y+b@kziczvh.com","r.cyo.g+x+d.c+f.t@tgsg.z.com","r.cyo.g+x+t.y.l.i@tgsg.z.com","r.cyo.g+brxxi@tgsg.z.com","r.cyo.g+z+dr.k.u@tgsg.z.com","r.cyo.g+d+l.c.n+g@tgsg.z.com","fg.r.u.uzj+vq.o@kziczvh.com","fg.r.u.uzj+uzq@kziczvh.com","fg.r.u.uzj+mvz@kziczvh.com","fg.r.u.uzj+taj@kziczvh.com","fg.r.u.uzj+fek@kziczvh.com"], 2),
        ]:
            output = self.numUniqueEmails(emails)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
