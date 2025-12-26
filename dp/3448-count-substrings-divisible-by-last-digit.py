import unittest

# https://leetcode.com/problems/count-substrings-divisible-by-last-digit/

# python3 -m unittest dp/3448-count-substrings-divisible-by-last-digit.py


class Solution(unittest.TestCase):
    # # Approach: Brute Force - Check All Substrings
    # # For each substring, track its value modulo the last digit and check divisibility.
    # #
    # # Time Complexity: O(n^2)
    # # Space Complexity: O(1)
    # def countSubstrings(self, s: str) -> int:
    #     n = len(s)
    #     count = 0
    
    #     for i in range(n):
    #         # For each possible divisor (last digit), track value mod that divisor
    #         mod = [0] * 10
    
    #         for j in range(i, n):
    #             digit = int(s[j])
    
    #             # Update modulo values for all divisors
    #             for divisor in range(1, 10):
    #                 mod[divisor] = (mod[divisor] * 10 + digit) % divisor
    
    #             # Check if divisible by last digit
    #             if digit != 0 and mod[digit] == 0:
    #                 count += 1
    
    #     return count

    # Approach: Dynamic Programming with Space Optimization
    # Key insight: For each position, track remainder of substrings ending here
    # when divided by each possible last digit (1-9).
    #
    # Strategy:
    # 1. Use dp[divisor][remainder] = count of substrings with this remainder
    # 2. For each new digit, update all divisors (1-9)
    # 3. Count when remainder is 0 and divisor equals current digit
    # 4. Optimize space: rolling array (current and previous state)
    #
    # Mathematical Analysis:
    # - For substring with value V, extending with digit d:
    #   New value = V * 10 + d
    #   New remainder mod i = (old_remainder * 10 + d) % i
    #
    # Time Complexity: O(n * 9 * 9) = O(n)
    # Space Complexity: O(9 * 9) = O(1)
    def countSubstrings(self, s: str) -> int:
        n = len(s)
        result = 0

        # dp[divisor][remainder] = count of substrings ending at current position
        # with this remainder when divided by divisor
        dp = [[0] * 10 for _ in range(10)]

        for idx in range(n):
            digit = int(s[idx])

            # Create new state for current position
            new_dp = [[0] * 10 for _ in range(10)]

            # For each possible divisor (1-9, can't divide by 0)
            for divisor in range(1, 10):
                # New substring starting at current position
                new_dp[divisor][digit % divisor] += 1

                # Extend existing substrings
                for remainder in range(divisor):
                    if dp[divisor][remainder] > 0:
                        new_remainder = (remainder * 10 + digit) % divisor
                        new_dp[divisor][new_remainder] += dp[divisor][remainder]

            # Count substrings ending at current position divisible by last digit
            if digit != 0:
                result += new_dp[digit][0]

            dp = new_dp

        return result

    def test(self):
        for s, expected in [
            ("12936", 11),
            ("5701283", 18),
            ("1010101010", 25),
            ("1212", 10),
            ("123", 5),
        ]:
            output = self.countSubstrings(s)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
