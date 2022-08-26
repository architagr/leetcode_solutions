package minimum_sum_of_four_digit_number_after_splitting_digits

import "sort"

func MinimumSum(num int) int {
	arr := make([]int, 4)
	k := 0
	for num > 0 {
		arr[k] = num % 10
		num /= 10
		k++
	}
	sort.Ints(arr)

	return (arr[0]+arr[1])*10 + (arr[2] + arr[3])
}
