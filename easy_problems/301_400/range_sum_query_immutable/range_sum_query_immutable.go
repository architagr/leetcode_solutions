package range_sum_query_immutable

type NumArray struct {
	prefixSum []int
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	prefixSum := make([]int, n)
	sum := 0
	for i := 0; i < n; i++ {

		sum += nums[i]
		prefixSum[i] = sum
	}
	return NumArray{
		prefixSum: prefixSum,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	leftVal := 0
	if left > 0 {
		leftVal = this.prefixSum[left-1]
	}
	return this.prefixSum[right] - leftVal
}
