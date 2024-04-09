class Solution {
  String reverseWords(String s) {
    final List<String> words = s.trim().split(RegExp(r'\s+'));

    int i = 0;
    int j = words.length - 1;

    while (i < j) {
      var temp = words[j];
      words[j] = words[i];
      words[i] = temp;
      i++;
      j--;
    }

    return words.join(" ");
  }
}
