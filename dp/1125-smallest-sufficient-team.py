from typing import List
import unittest

# https://leetcode.com/problems/smallest-sufficient-team/

# python3 -m unittest dp/1125-smallest-sufficient-team.py


class Solution(unittest.TestCase):
    # # Approach #1: Top-Down Dynamic Programming + Bit Manipulation
    # # Time: O(2^sn * pn), sn=len(reqSkills), pn=len(people)
    # # Space: O(2^sn)
    # def smallestSufficientTeam(self, req_skills: List[str], people: List[List[str]]) -> List[int]:
    #     indices = {skill: idx for idx, skill in enumerate(req_skills)}
    #     sn, pn = len(req_skills), len(people)
    #     skills = [0] * pn
    #     for person in range(pn):
    #         for skill in people[person]:
    #             skills[person] |= 1 << indices[skill]

    #     MAX = (1 << sn) - 1
    #     memo = [-1] * (MAX + 1)
    #     memo[(1 << pn) - 1] = 0

    #     def dp(smask: int) -> int:
    #         if memo[smask] != -1:
    #             return memo[smask]
    #         for person in range(pn):
    #             if smask | skills[person] != smask:
    #                 pmask = dp(smask | skills[person]) | (1 << person)
    #                 if memo[smask] == -1 or pmask.bit_count() < memo[smask].bit_count():
    #                     memo[smask] = pmask
    #         return memo[smask]

    #     mask = dp(0)
    #     team: List[int] = []
    #     for person in range(pn):
    #         if mask & (1 << person):
    #             team.append(person)
    #     return team

    # Approach #2: Bottom-Up Dynamic Programming + Bit Manipulation
    # Time: O(2^sn * pn), sn=len(reqSkills), pn=len(people)
    # Space: O(2^sn)
    def smallestSufficientTeam(self, req_skills: List[str], people: List[List[str]]) -> List[int]:
        indices = {skill: idx for idx, skill in enumerate(req_skills)}
        sn, pn = len(req_skills), len(people)
        skills = [0] * pn
        for person in range(pn):
            for skill in people[person]:
                skills[person] |= 1 << indices[skill]

        MAX = (1 << sn) - 1
        dp = [(1 << pn) - 1] * (MAX + 1)
        dp[0] = 0

        for smask in range(1, MAX + 1):
            for person in range(pn):
                prev_mask = smask & ~skills[person]
                if prev_mask != smask:
                    pmask = dp[prev_mask] | (1 << person)
                    if pmask.bit_count() < dp[smask].bit_count():
                        dp[smask] = pmask

        mask = dp[MAX]
        team: List[int] = []
        for person in range(pn):
            if mask & (1 << person):
                team.append(person)
        return team

    def test(self):
        for req_skills, people, expected in [
            (["java", "nodejs", "reactjs"], [["java"], ["nodejs"], ["nodejs", "reactjs"]], [0, 2]),
            (
                ["algorithms", "math", "java", "reactjs", "csharp", "aws"],
                [["algorithms", "math", "java"], ["algorithms", "math", "reactjs"], ["java", "csharp", "aws"], ["reactjs", "csharp"], ["csharp", "math"], ["aws", "java"]],
                [1, 2],
            ),
            # (["uhppib", "mgdipkgt", "vaostwmycy", "bjwxnfbbubpnd"], [["vaostwmycy"], ["mgdipkgt"], ["bjwxnfbbubpnd"], [], ["uhppib"]], [0, 1, 2, 4]),
        ]:
            output = self.smallestSufficientTeam(req_skills, people)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
