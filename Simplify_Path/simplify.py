from typing import List


class Solution:
    def simplifyPath(self, path: str) -> str:
        simplify_path = ""
        index = 0
        stack: List[str] = []

        while index < len(path):
            if path[index] == "/":
                index += 1
                continue

            start = index

            while index < len(path) and path[index] != "/":
                index += 1

            sub_str = path[start:index]

            if sub_str == "..":
                if len(stack) > 0:
                    stack = stack[: len(stack) - 1]
            elif sub_str != ".":
                stack.append(sub_str)

            index += 1

        if len(stack) == 0:
            return "/"

        for dir in stack:
            simplify_path += f"/{dir}"

        return simplify_path
