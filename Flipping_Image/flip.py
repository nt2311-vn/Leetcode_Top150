from typing import List


class Solution:
    def flipAndInvertImage(self, image: List[List[int]]) -> List[List[int]]:
        flip_image: List[List[int]] = []
        if len(image) == 0 or len(image[0]) == 0:
            return flip_image

        i = 0

        while i < len(image):
            flip_row = []

            j = len(image[0]) - 1

            while j >= 0:
                flip_row.append(image[i][j] ^ 1)

                j -= 1

            flip_image.append(flip_row)
            i += 1

        return flip_image
