class Solution {
  bool isHappy(int n) {
    final memoir = {};

    while (n != 1) {
      if (memoir.containsKey(n)) {
        return false;
      }
      memoir[n] = true;
      String numStr = n.toString();
      int result = 0;

      for (var i = 0; i < numStr.length; i++) {
        result += int.parse(numStr[i]) * int.parse(numStr[i]);
      }

      if (result == 1) {
        return true;
      }
      n = result;
    }

    return true;
  }
}
