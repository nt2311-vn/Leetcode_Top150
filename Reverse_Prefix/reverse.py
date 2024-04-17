class Solution:
    def reversePrefix(self, word: str, ch: str) -> str:
        is_found, index = False, 0

        while index < len(word):
            if word[index] == ch:
                is_found = True
                break
            index += 1

        if not is_found:
            return word

        char_list = list(word)
        i, j = 0, index

        while i < j:
            char_list[i], char_list[j] = char_list[j], char_list[i]

            i, j = i + 1, j - 1
        return "".join(char_list)
