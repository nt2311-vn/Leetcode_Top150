from typing import List


class Solution:
    def finalPrices(self, prices: List[int]) -> List[int]:
        final_price: List[int] = []

        index = 0

        while index < len(prices):
            discount = 0
            next_price = index + 1

            while next_price < len(prices):
                if prices[next_price] <= prices[index]:
                    discount = prices[next_price]
                    break

                next_price += 1

            final_price.append(prices[index] - discount)
            index += 1

        return final_price
