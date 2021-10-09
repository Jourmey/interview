package main


func search(nums []int, target int) int {
	inde := index(nums, target)
	if inde < 0 {
		return 0
	}

	var a, b int
	for i := inde; i < len(nums); i++ {
		if nums[i] == target {
			b = i
		}
	}

	for i := inde; i >= 0; i-- {
		if nums[i] == target {
			a = i
		}
	}

	return b - a + 1
}

func index(nums []int, target int) int {
	i := 0
	j := len(nums) - 1
	for i <= j {
		mi := (i + j) / 2
		mv := nums[mi]
		if mv < target {
			i = mi + 1
		} else if target < mv {
			j = mi - 1
		} else {
			return mi
		}
	}
	return -1
}

func missingNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if i != nums[i] {
			return i
		}
	}

	return len(nums)
}
