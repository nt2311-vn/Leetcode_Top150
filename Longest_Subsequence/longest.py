from typing import List


class Solution:
    def answerQueries(self, nums: List[int], queries: List[int]) -> List[int]:
        result: List[int] = [0] * len(queries)
        nums = sorted(nums)

        result_index = 0

        while result_index < len(queries):

            sum, subsequence_count, num_index = 0, 0, 0

            while num_index < len(nums):
                if sum + nums[num_index] <= queries[result_index]:
                    sum += nums[num_index]
                    subsequence_count += 1

                else:
                    break

                num_index += 1

            result[result_index] = subsequence_count

            result_index += 1

        return result
