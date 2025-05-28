package problem_0001

func TwoSum(nums []int, target int) []int {
	return twoSum(nums, target)
}

func twoSum(nums []int, target int) []int {
	for i, n := range nums {
		for j, m := range nums[i+1:] {
			if n+m == target {
				return []int{i, i + j + 1}
			}
		}
	}
	return nil
}
