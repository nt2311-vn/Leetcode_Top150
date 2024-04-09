class Solution {
  String longestCommonPrefix(List<String> strs) {
    final commonPrefix = strs[0];

    for (var i = 0; i < commonPrefix.length; i++) {
      for (var j = 0; j < strs.length; j++) {
        if (i >= strs[j].length || strs[j][i] != commonPrefix[i]) {
          return commonPrefix.substring(0, i);
        }
      }
    }

    return commonPrefix;
  }
}
