class Solution {
  String convert(String s, int numRows) {
    if (numRows == 1 || s.length <= numRows) {
      return s;
    }

    final List<String> arrChar = List.filled(s.length, "", growable: false);
    int count = 0;
    bool zig = true;

    for (var i = 0; i < s.length; i++) {
      arrChar[count] += s[i];

      if (zig) {
        count++;
      } else {
        count--;
      }

      if (count == numRows - 1 || count == 0) {
        zig = !zig;
      }
    }

    return arrChar.join("");
  }
}
