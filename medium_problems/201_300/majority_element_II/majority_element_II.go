package majority_element_II

func MajorityElementII(nums []int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	result := make([]int, 0, len(m))
	nby3 := len(nums) / 3
	for ele, count := range m {
		if count > nby3 {
			result = append(result, ele)
		}
	}
	return result
}

func MajorityElementII_approch2(nums []int) []int {
	c1, c2 := 0, 0
	v1, v2 := 0, 0
	for _, v := range nums {
		if c1 == 0 && (c2 == 0 || v2 != v) {
			v1 = v
		} else if c2 == 0 && c1 != 0 && v1 != v {
			v2 = v
		}
		if v1 == v {
			c1++
		} else if v2 == v {
			c2++
		} else {
			c1--
			c2--
		}
	}
	f := func(c1, v1 int) []int {
		if c1 == 0 {
			return nil
		}
		c1 = 0
		for _, v := range nums {
			if v1 == v {
				c1++
			}
		}
		if c1*3 > len(nums) {
			return []int{v1}
		}
		return nil
	}
	return append(f(c1, v1), f(c2, v2)...)
}
