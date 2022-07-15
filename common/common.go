package common

func MaxInt() int {
	const MaxUint = ^uint(0)
	const MaxInt = int(MaxUint >> 1)
	return MaxInt
}
func MinInt() int {
	const MaxUint = ^uint(0)
	const MaxInt = int(MaxUint >> 1)
	const MinInt = -MaxInt - 1
	return MinInt
}
