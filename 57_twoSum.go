package main

func main() {

	num := []int{2, 7, 11, 15}
	t := 9

	twoSum(num, t)
}

// 双指针
func twoSum(nums []int, target int) []int {
	left := 0
	right := len(nums) - 1
	for left < right {
		if nums[left]+nums[right] < target {
			left++
		} else if nums[left]+nums[right] > target {
			right--
		} else {
			return []int{nums[left], nums[right]}
		}
	}
	return nil
}
