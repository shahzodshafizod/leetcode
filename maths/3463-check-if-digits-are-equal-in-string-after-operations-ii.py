import unittest

# https://leetcode.com/problems/check-if-digits-are-equal-in-string-after-operations-ii/

# python3 -m unittest maths/3463-check-if-digits-are-equal-in-string-after-operations-ii.py


class Solution(unittest.TestCase):
    # Approach: Pascal's Triangle + Modular Arithmetic
    # The operation reduces string length by 1 each time: n → n-1 → ... → 2
    # Final two digits are weighted sums with coefficients from Pascal's triangle row (n-2)
    # We need to check if sum(C(n-2, i) * (s[i] - s[i+1])) ≡ 0 (mod 10)
    # To compute C(n-2, k) mod 10 efficiently, factor out powers of 2 and 5
    # Time: O(n), Space: O(n)
    def hasSameDigits(self, s: str) -> bool:
        n = len(s)
        if n == 2:
            return s[0] == s[1]

        # Precompute factorials as (twos, fives, other) tuples
        # where x! = 2^twos * 5^fives * other, with gcd(other, 10) = 1
        def factorize(x: int) -> tuple[int, int, int]:
            twos = 0
            while x % 2 == 0:
                x //= 2
                twos += 1
            fives = 0
            while x % 5 == 0:
                x //= 5
                fives += 1
            return (twos, fives, x % 10)

        facts = [(0, 0, 1)]  # 0! = 1
        for i in range(1, n - 1):
            prev = facts[-1]
            curr = factorize(i)
            facts.append((prev[0] + curr[0], prev[1] + curr[1], (prev[2] * curr[2]) % 10))

        # Modular inverse mod 10 (only exists for 1, 3, 7, 9)
        inv = {1: 1, 3: 7, 7: 3, 9: 9}

        # Compute C(n-2, k) mod 10
        def comb(n_val: int, k: int) -> int:
            if k < 0 or k > n_val:
                return 0
            # C(n, k) = n! / (k! * (n-k)!)
            num = facts[n_val]
            den_k = facts[k]
            den_nk = facts[n_val - k]

            # Combine denominators
            den_twos = den_k[0] + den_nk[0]
            den_fives = den_k[1] + den_nk[1]
            den_other = (den_k[2] * den_nk[2]) % 10

            # Compute result
            twos = num[0] - den_twos
            fives = num[1] - den_fives
            other = (num[2] * inv[den_other]) % 10

            # Convert back to single digit
            if twos >= 1 and fives >= 1:
                return 0
            if fives >= 1:
                return 5
            if twos == 0:
                return other
            # Powers of 2 mod 10 cycle: 2, 4, 8, 6, 2, 4, 8, 6, ...
            pow2_cycle = [6, 2, 4, 8]
            return (other * pow2_cycle[twos % 4]) % 10

        # Check if weighted sum is 0 mod 10
        total = 0
        for i in range(n - 1):
            coeff = comb(n - 2, i)
            diff = int(s[i]) - int(s[i + 1])
            total = (total + coeff * diff) % 10

        return total == 0

    def test(self):
        for s, expected in [
            ("3902", True),
            ("34789", False),
            ("111", True),
            ("12", False),
        ]:
            output = self.hasSameDigits(s)
            self.assertEqual(expected, output, f"expected: {expected}, output: {output}")
