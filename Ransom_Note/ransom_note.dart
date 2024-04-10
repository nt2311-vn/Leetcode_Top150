class Solution {
  bool canConstruct(String ransomNote, String magazine) {
    final hashMap = {};

    for (var i = 0; i < magazine.length; i++) {
      if (hashMap.containsKey(magazine[i])) {
        hashMap[magazine[i]]++;
      } else {
        hashMap[magazine[i]] = 1;
      }
    }

    for (var i = 0; i < ransomNote.length; i++) {
      if (hashMap.containsKey(ransomNote[i])) {
        hashMap[ransomNote[i]]--;
      } else {
        hashMap[ransomNote[i]] = -1;
      }

      if (hashMap[ransomNote[i]] < 0) {
        return false;
      }
    }

    return true;
  }
}
