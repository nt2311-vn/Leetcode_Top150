class Solution:
    def reverseWords(self, s: str) -> str:

        words = s.split(" ")
        i = 0

        while i < len(words):
            revword = words[i][::-1]
            words[i] = revword
            i += 1

        return " ".join(words)
