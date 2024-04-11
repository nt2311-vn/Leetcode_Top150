class Solution {
  bool isAnagram(String s, String t) {
    if (s.length != t.length) {
      return false;
    }

    final charMap = {};

    for (var i = 0; i < s.length; i++) {
      charMap.containsKey(s[i]) ? charMap[s[i]]++ : charMap[s[i]] = 1;
    }
    for (var i = 0; i < t.length; i++) {
      charMap.containsKey(t[i]) ? charMap[t[i]]-- : charMap[t[i]] = -1;

      if (charMap[t[i]] < 0) {
        return false;
      }
    }

    return true;
  }
}
