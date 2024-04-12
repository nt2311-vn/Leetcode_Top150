class Solution {
  List<List<String>> groupAnagrams(List<String> strs) {
    final Map<String, List<String>> groupAnagrams = {};

    for (var str in strs) {
      var keyStr = str.split('')
        ..sort()
        ..join("");

      if (groupAnagrams.containsKey(keyStr)) {
        groupAnagrams[keyStr]!.add(str);
      } else {
        groupAnagrams[keyStr] = [str];
      }
    }

    final List<List<String>> result = [];
    for (var group in groupAnagrams.values) {
      result.add(group);
    }

    return result;
  }
}
