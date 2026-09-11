package counttotalnumberofcoloredcells

func coloredCells(n int) int64 {

	m := int64(n)
	return 2*m*(m-1) + 1
}
