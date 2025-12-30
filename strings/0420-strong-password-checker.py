import unittest

# https://leetcode.com/problems/strong-password-checker/

# python3 -m unittest strings/0420-strong-password-checker.py


class Solution(unittest.TestCase):
    # Approach: Optimized Greedy Algorithm
    # We analyze the problem by length categories:
    # 1. length < 6: need insertions (can fix types + repeats)
    # 2. length > 20: need deletions (strategic to help with repeats)
    # 3. 6 <= length <= 20: use replacements
    # Time: O(n) - single pass to count repeats, process deletions/replacements
    # Space: O(1) - only tracking counts and sequences
    def strongPasswordChecker(self, password: str) -> int:
        n = len(password)

        # Check character type requirements
        has_lower = any(c.islower() for c in password)
        has_upper = any(c.isupper() for c in password)
        has_digit = any(c.isdigit() for c in password)
        missing_types = (0 if has_lower else 1) + (0 if has_upper else 1) + (0 if has_digit else 1)

        # Count repeating sequences
        # We track sequences by their length % 3 (important for optimization)
        total_replacements = 0  # replacements needed for repeating sequences
        one_mod = 0   # sequences where len % 3 == 0 (one delete helps most)
        two_mod = 0   # sequences where len % 3 == 1 (two deletes help most)

        i = 2
        while i < n:
            if password[i] == password[i-1] == password[i-2]:
                length = 2
                while i < n and password[i] == password[i-1]:
                    length += 1
                    i += 1

                # Each sequence of length L needs L//3 replacements
                total_replacements += length // 3

                # Track by modulo for deletion optimization
                if length % 3 == 0:
                    one_mod += 1
                elif length % 3 == 1:
                    two_mod += 1
            else:
                i += 1

        # Case 1: Too short (length < 6)
        if n < 6:
            # Need to insert characters
            insertions_needed = 6 - n
            # Insertions can fix both missing types and some repeats
            return max(missing_types, insertions_needed)

        # Case 2: Too long (length > 20)
        elif n > 20:
            deletions_needed = n - 20
            remaining_deletions = deletions_needed

            # Optimize deletions to reduce replacements needed
            # Priority: sequences with len%3==0 need only 1 deletion to reduce replacements
            use = min(remaining_deletions, one_mod)
            total_replacements -= use
            remaining_deletions -= use

            # Next: sequences with len%3==1 need 2 deletions
            use = min(remaining_deletions, two_mod * 2)
            total_replacements -= use // 2
            remaining_deletions -= use

            # Remaining deletions: each 3 deletions saves 1 replacement
            total_replacements -= remaining_deletions // 3

            return deletions_needed + max(missing_types, total_replacements)

        # Case 3: Length is fine (6 <= length <= 20)
        else:
            # Only need replacements
            # Each replacement can fix one repeat AND potentially one missing type
            return max(missing_types, total_replacements)

    def test(self):
        for password, expected in [
            ("a", 5),        # Too short, missing types
            ("aA1", 3),      # Too short
            ("1337C0d3", 0), # Already strong
        ]:
            output = self.strongPasswordChecker(password)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
