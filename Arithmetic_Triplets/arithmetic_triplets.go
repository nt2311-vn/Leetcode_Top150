package main

func arithmeticTriplets(nums []int, diff int) int {
	if len(nums) < 3 {
		return 0
	}

	triplets := 0
	for i := 0; i < len(nums)-2; i++ {
		diffEqual, compareNum := 0, nums[i]
		for j := i + 1; j < len(nums); j++ {
			if nums[j]-compareNum == diff {
				diffEqual++
				compareNum = nums[j]
			}

			if diffEqual == 2 {
				triplets++
				break
			}
		}
	}

	return triplets
}
