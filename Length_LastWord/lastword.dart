class Solution {
  int lengthOfLastWord(String s) {
    final words = s.trim().split(" ");
    return words.last.length;
  }
}
