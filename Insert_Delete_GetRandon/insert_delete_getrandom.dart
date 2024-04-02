import 'dart:collection';
import 'dart:math';

class RandomizedSet {
  List<int> arr = [];

  RandomizedSet();

  bool insert(int val) {
    if (arr.contains(val)) {
      return false;
    }
    arr.add(val);
    return true;
  }

  bool remove(int val) {
    if (!arr.contains(val)) {
      return false;
    }

    arr.remove(val);
    return true;
  }

  int getRandom() {
    var index = Random().nextInt(arr.length);
    return arr[index];
  }
}
