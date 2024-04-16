class Solution:
    def removeOuterParentheses(self, s: str) -> str:
        result = ""
        pointer = 1
        index = 1

        while index < len(s):
            if s[index] == "(":
                pointer += 1
            else:
                pointer -= 1

            if pointer == 0:
                pointer = 1
                index += 1
            else:
                result += s[index]

            index += 1

        return result
