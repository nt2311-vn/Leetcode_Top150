class Solution {
  bool isPalindrome(String s) {
    final String cleaned = s.toLowerCase().replaceAll(RegExp(r'[^a-z0-9]'), '');

    int i = 0;
    int j = cleaned.length - 1;

    while (i < j) {
      if (cleaned[i] != cleaned[j]) {
        return false;
      }
      i++;
      j--;
    }

    return true;
  }
}
