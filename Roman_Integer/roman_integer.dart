class Solution {
  final hashDouble = {
    "IV": 4,
    "IX": 9,
    "XL": 40,
    "XC": 90,
    "CD": 400,
    "CM": 900
  };

  final hashSingle = {
    "I": 1,
    "V": 5,
    "X": 10,
    "L": 50,
    "C": 100,
    "D": 500,
    "M": 1000
  };
  int romanToInt(String s) {
    int result = 0;

    for (var i = 0; i < s.length; i++) {
      if (i + 1 < s.length) {
        if (hashDouble.containsKey(s.substring(i, i + 2))) {
          result += hashDouble[s.substring(i, i + 2)]!;
          i++;
          continue;
        }
      }

      result += hashSingle[s[i]]!;
    }
    return result;
  }
}
