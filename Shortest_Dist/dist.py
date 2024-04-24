from typing import List


class Solution:
    def shortestToChar(self, s: str, c: str) -> List[int]:
        dist_list: List[int] = []

        k = 0

        while k < len(s):
            if s[k] == c:
                dist_list.append(0)
            else:
                i, j = k - 1, k + 1

                while i >= 0 and s[i] != c:
                    i -= 1

                while j < len(s) and s[j] != c:
                    j += 1

                min_dist = min(k - i, j - k)

                if i == -1:
                    min_dist = j - k

                if j == len(s):
                    min_dist = k - i

                dist_list.append(min_dist)

            k += 1

        return dist_list
