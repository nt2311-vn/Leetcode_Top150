class Solution {
  int maxProfit(List<int> prices) {
    if (prices.length == 1) {
      return 0;
    }

    var buyPrice = prices[0];
    var maxProfit = 0;

    for (var i = 1; i < prices.length; i++) {
      if (prices[i] < buyPrice) {
        buyPrice = prices[i];
      } else {
        var profit = prices[i] - buyPrice;

        if (profit > 0) {
          maxProfit += profit;
          buyPrice = prices[i];
        }
      }
    }
    return maxProfit;
  }
}
