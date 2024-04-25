from typing import List


class Solution:
    def kWeakestRows(self, mat: List[List[int]], k: int) -> List[int]:
        row_info: List[RowInfo] = []

        i = 0
        while i < len(mat):
            row_info.append(RowInfo(i, count_soldiers(mat[i])))
            i += 1

        row_info = sorted(row_info, key=lambda x: (x.soldiers, x.row_index))

        weakest_rows = [info.row_index for info in row_info[:k]]
        return weakest_rows


class RowInfo:
    def __init__(self, row_index: int, soldiers: int):
        self.row_index = row_index
        self.soldiers = soldiers


def count_soldiers(row: List[int]) -> int:
    i, j = 0, len(row) - 1

    while i <= j:
        mid = (i + j) // 2

        if row[mid] == 1:
            i = mid + 1
        else:
            j = mid - 1

    return i
