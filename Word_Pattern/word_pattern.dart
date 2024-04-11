class Solution {
  bool wordPattern(String pattern, String s) {
    final words = s.split(" ");
    if (words.length != pattern.length) {
      return false;
    }

    final patternTable = {};
    final wordTable = {};

    for (var i = 0; i < pattern.length; i++) {
      if (patternTable.containsKey(pattern[i]) &&
          patternTable[pattern[i]] != words[i]) {
        return false;
      }
      patternTable[pattern[i]] = words[i];

      if (wordTable.containsKey(words[i]) &&
          wordTable[words[i]] != pattern[i]) {
        return false;
      }
      wordTable[words[i]] = pattern[i];
    }

    return true;
  }
}
