package LeetCode_Go

func twoSum(nums []int, target int) [2]int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == target-nums[i] {
				result := [...]int{i, j}
				return result
			}
		}
	}
	return [2]int{0, 0}
}
