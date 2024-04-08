class Solution {
  int canCompleteCircuit(List<int> gas, List<int> cost) {
    int fill = 0;
    int energy = 0;
    int start = 0;

    for (var i = 0; i < gas.length; i++) {
      fill += gas[i] - cost[i];
      energy += gas[i] - cost[i];

      if (fill < 0) {
        start = i + 1;
        fill = 0;
      }
    }

    if (energy < 0) {
      return -1;
    }

    return start;
  }
}
