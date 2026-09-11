package range_sum_query_mutable

type NumArray struct {
	originalArray  []int
	prefixSumArray []int
}

func Constructor(nums []int) NumArray {
	prefixSumArray := make([]int, len(nums))
	prefixSumArray[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		prefixSumArray[i] = nums[i] + prefixSumArray[i-1]
	}
	return NumArray{
		originalArray:  nums,
		prefixSumArray: prefixSumArray,
	}
}

func (this *NumArray) Update(index int, val int) {
	diff := val - this.originalArray[index]

	this.originalArray[index] = val

	for i := index; i < len(this.originalArray); i++ {
		this.prefixSumArray[i] += diff
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	ans := this.prefixSumArray[right]

	if left > 0 {
		ans -= this.prefixSumArray[left-1]
	}
	return ans
}
