class Solution {
  bool isPalindrome(int x) {
    final String numStr = x.toString();

    for (var i = 0; i < numStr.length / 2; i++) {
      if (numStr[i] != numStr[numStr.length - 1 - i]) {
        return false;
      }
    }
    return true;
  }
}
