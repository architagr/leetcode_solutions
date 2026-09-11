package typeoftriangle

func triangleType(nums []int) string {
	m := make(map[int]struct{})

	m[nums[0]] = struct{}{}
	m[nums[1]] = struct{}{}
	m[nums[2]] = struct{}{}

	if nums[0]+nums[1] <= nums[2] || nums[0]+nums[2] <= nums[1] || nums[2]+nums[1] <= nums[0] {
		return "none"
	}

	if len(m) == 3 {
		return "scalene"
	} else if len(m) == 2 {
		return "isosceles"
	}
	return "equilateral"

}
