class Solution:
    def makeSmallestPalindrome(self, s: str) -> str:
        char_list = list(s)

        i, j = 0, len(char_list) - 1

        while i < j:
            if char_list[i] != char_list[j]:
                char_list[i] = min(char_list[i], char_list[j])
                char_list[j] = char_list[i]

            i += 1
            j -= 1

        return "".join(char_list)
