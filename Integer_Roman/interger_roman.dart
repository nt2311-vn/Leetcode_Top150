class Solution {
  final numTable = {
    1: "I",
    4: "IV",
    5: "V",
    9: "IX",
    10: "X",
    40: "XL",
    50: "L",
    90: "XC",
    100: "C",
    400: "CD",
    500: "D",
    900: "CM",
    1000: "M",
  };

  final List<int> divideOrder = [
    1000,
    900,
    500,
    400,
    100,
    90,
    50,
    40,
    10,
    9,
    5,
    4,
    1
  ];
  String intToRoman(int num) {
    String romanStr = "";
    while (num > 0) {
      for (var numInt in divideOrder) {
        int wholeVal = num ~/ numInt;
        if (wholeVal >= 1) {
          num -= wholeVal * numInt;
          while (wholeVal > 0) {
            romanStr += numTable[numInt]!;
            wholeVal--;
          }
        }
      }
    }

    return romanStr;
  }
}
