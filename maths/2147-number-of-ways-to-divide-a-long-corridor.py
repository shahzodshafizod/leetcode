import unittest

# https://leetcode.com/problems/number-of-ways-to-divide-a-long-corridor/

# python3 -m unittest maths/2147-number-of-ways-to-divide-a-long-corridor.py


class Solution(unittest.TestCase):
    # Time: O(N)
    # Space: O(1)
    def numberOfWays(self, corridor: str) -> int:
        mod = int(1e9 + 7)
        seats = plants = 0
        ways = 1
        for c in corridor:
            if c == 'S':
                seats += 1
            elif seats == 2:
                plants += 1
            if seats == 3:
                ways = (ways * (plants + 1)) % mod
                seats, plants = 1, 0
        if seats != 2:
            return 0
        return ways

    def test(self):
        for corridor, expected in [
            ("SSPPSPS", 3),
            ("PPSPSP", 1),
            ("S", 0),
            ("P", 0),
            ("PPSPSPPSPSPSPSPS", 0),
            ("SPPSSSSPPS", 1),
            ("SPSPPSSPSSSS", 6),
            ("PPPSPPPSPSSPPSPSSPSSPPPPSSPSSPPSPPPSSSPSSSPSSSSPSSSSSPSSPSPPSSPSSPPSSSPSPPPSSSSSPSSPPPSSPPSSPSSSPPSP", 0),
            ("PPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPPP", 0),
            (
                "PSPSPPPSSSSPSSPSSPSPPSSSPPSPSPPPSSSPSPPSSSPSSSPPSPPPSPSPPPSPSPPSSPPSSPSPSSPPPPPSPSPSSSPSSSSPPPSPSPPS",
                663552,
            ),
            ("SSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSS", 1),
        ]:
            output = self.numberOfWays(corridor)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
