class Solution:
    def trailingZeroes(self, n: int) -> int:

        trail_zero = 0
        while n > 0:
            trail_zero += int(n / 5)
            n = int(n / 5)

        return trail_zero
