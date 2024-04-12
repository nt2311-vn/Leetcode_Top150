from typing import Dict, List


class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        group_anagrams: Dict[str, List[str]] = {}

        for strElement in strs:
            key_str = "".join(sorted(strElement))

            if key_str in group_anagrams:
                group_anagrams[key_str].append(strElement)

            else:
                group_anagrams[key_str] = [strElement]

        result: List[List[str]] = []

        for group in group_anagrams.values():
            result.append(group)

        return result
