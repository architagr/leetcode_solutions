package sum

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	for i := 0; i < len(nums) && nums[i] <= 0; i++ {
		if i == 0 || nums[i-1] != nums[i] {
			twoSum(nums, i, &res)
		}
	}
	return res
}

func twoSum(nums []int, i int, res *[][]int) {
	seen := map[int]bool{}
	for j := i + 1; j < len(nums); j++ {
		complement := -nums[i] - nums[j]
		if seen[complement] {
			*res = append(*res, []int{nums[i], nums[j], complement})
			for j+1 < len(nums) && nums[j] == nums[j+1] {
				j++
			}
		}
		seen[nums[j]] = true
	}
}
