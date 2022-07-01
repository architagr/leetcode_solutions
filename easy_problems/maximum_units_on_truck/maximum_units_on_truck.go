package maximum_units_on_truck

import (
	"sort"
)

func MaxUnitsOnTruck(boxTypes [][]int, truckSize int) int {
	// sort array in decreasing order of the number of units
	// helps to get the max units on the first
	// then we can start taking boxes from the oth position
	// and fill the truck

	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})
	sum := 0
	for i := 0; i < len(boxTypes) && truckSize > 0; i++ {
		x := min(truckSize, boxTypes[i][0])
		truckSize -= x
		sum += x * boxTypes[i][1]
	}
	return sum
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
