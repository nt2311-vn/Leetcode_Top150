class Solution {
  final Map<String, String> hashComplement = {'{': '}', '(': ')', '[': ']'};
  bool isValid(String s) {
    if (s.length % 2 != 0) {
      return false;
    }

    final List<String> stack = [];

    for (var i = 0; i < s.length; i++) {
      if (stack.length > 0 && s[i] == hashComplement[stack.last]) {
        stack.removeLast();
      } else {
        stack.add(s[i]);
      }
    }

    return stack.length == 0;
  }
}
