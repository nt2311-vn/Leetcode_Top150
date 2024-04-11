class Solution {
  bool isIsomorphic(String s, String t) {
    final firstMap = {};
    final sencondMap = {};

    for (var i = 0; i < s.length; i++) {
      if (firstMap.containsKey(s[i]) && firstMap[s[i]] != t[i]) {
        return false;
      }
      firstMap[s[i]] = t[i];
    }

    for (var i = 0; i < s.length; i++) {
      if (sencondMap.containsKey(t[i]) && sencondMap[t[i]] != s[i]) {
        return false;
      }
      sencondMap[t[i]] = s[i];
    }

    return true;
  }
}
