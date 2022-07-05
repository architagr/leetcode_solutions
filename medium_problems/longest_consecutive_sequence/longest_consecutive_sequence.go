package longest_consecutive_sequence

func LongestConsecutive(nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	max_len := 0
	num_map := map[int]bool{}
	for _, n := range nums {
		num_map[n] = true
	}

	for n := range num_map {
		if num_map[n-1] {
			continue
		}
		tmp_n, tmp_len := n, 1
		for num_map[tmp_n+1] {
			tmp_n++
			tmp_len++
		}
		if tmp_len > max_len {
			max_len = tmp_len
		}
	}
	return max_len
}
