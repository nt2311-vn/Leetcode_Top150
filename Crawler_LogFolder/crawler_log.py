from typing import List


class Solution:
    def minOperations(self, logs: List[str]) -> int:
        stack: List[bool] = []
        index = 0

        while index < len(logs):

            if logs[index] == "../":
                if len(stack) > 0:
                    stack = stack[1:]
            elif logs[index] == "./":
                pass
            else:
                stack.append(True)

            index += 1

        return len(stack)
