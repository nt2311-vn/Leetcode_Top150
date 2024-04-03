class Solution {
  int countStudents(List<int> students, List<int> sandwiches) {
    int s = sandwiches.length;
    while (students.isNotEmpty && s > 0) {
      bool found = false;
      for (int i = 0; i < students.length; i++) {
        if (students[i] == sandwiches[s]) {
          students.removeAt(i);
          s--;
          found = true;
          break;
        }
      }

      if (s == 0) {
        break;
      }
    }

    return students.length;
  }
}
