class Solution:
    def mySqrt(self, x: int) -> int:

        if x < 2:
            return x

        left, right = 0, x

        while left <= right:
            mid = int((left + right) / 2)

            if mid <= x / mid:
                left = mid + 1
            else:
                right = mid - 1

        return right
