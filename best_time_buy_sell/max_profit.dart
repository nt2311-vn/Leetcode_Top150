class Solution {
  int maxProfit(List<int> prices) {
    if (prices.length == 1) {
      return 0;
    }
    var maxProfit = 0;
    var minPrice = prices[0];

    for (var i = 1; i < prices.length; i++) {
      if (prices[i] < minPrice) {
        minPrice = prices[i];
      } else {
        if (prices[i] - minPrice > maxProfit) {
          maxProfit = prices[i] - minPrice;
        }
      }
    }
    return maxProfit;
  }
}
