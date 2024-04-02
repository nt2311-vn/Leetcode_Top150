class Solution {
  int hIndex(List<int> citations) {
    var result = 0;

    citations.sort();

    for (var i = 0; i < citations.length; i++) {
      var position = citations.length - i;

      if (citations[i] >= position) {
        result = position;
        return result;
      }
    }

    return result;
  }
}
