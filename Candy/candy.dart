class Solution {
  int candy(List<int> ratings) {
    final List<int> candies = List.filled(ratings.length, 1, growable: false);

    for (var i = 1; i < ratings.length; i++) {
      if (ratings[i - 1] < ratings[i]) {
        candies[i] = candies[i - 1] + 1;
      }
    }

    for (var i = ratings.length - 2; i >= 0; i--) {
      if (ratings[i + 1] < ratings[i] && candies[i + 1] >= candies[i]) {
        candies[i] = candies[i + 1] + 1;
      }
    }

    return candies.reduce((value, element) => value + element);
  }
}
