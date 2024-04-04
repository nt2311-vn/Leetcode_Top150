import 'dart:collection';

class Solution {
  HashMap<int, int> memoir = HashMap<int, int>();
  int climbStairs(int n) {
    if (memoir.containsKey(n)) {
      return memoir[n]!;
    }

    if (n <= 2) {
      return n;
    }

    memoir[n] = climbStairs(n - 1) + climbStairs(n - 2);

    return memoir[n]!;
  }
}
