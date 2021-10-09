package main

func exchange(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	i := 0
	j := len(nums) - 1
	for i < j {
		if nums[i]%2 == 0 && nums[j]%2 == 1 {
			nums[i], nums[j] = nums[j], nums[i]
			continue
		}

		if nums[i]%2 == 1 {
			i++
		}

		if nums[j]%2 == 0 {
			j--
		}
	}
	return nums
}
