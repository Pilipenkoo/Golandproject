package main

func twoSum(nums []int, target int) []int {
	var res []int
	for i, _ := range nums {
		if len(res) != 0 {
			break
		}
		for j := i + 1; j < len(nums); j++ {
			if target == nums[j]+nums[i] {
				res = append(res, i, j)
				break
			}
		}
	}
	return res
}
