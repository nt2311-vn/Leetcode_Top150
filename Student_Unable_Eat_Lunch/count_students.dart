class Solution {
  int countStudents(List<int> students, List<int> sandwiches) {
    int s = 0;
    while (students.isNotEmpty && s < sandwiches.length) {
      bool found = false;
      for (int i = 0; i < students.length; i++) {
        if (students[i] == sandwiches[s]) {
          students.removeAt(i);
          s++;
          found = true;
          break;
        }
      }

      if (!found) {
        break;
      }
    }

    return students.length;
  }
}
