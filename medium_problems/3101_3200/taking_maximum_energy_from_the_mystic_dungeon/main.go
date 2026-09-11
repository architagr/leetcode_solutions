package takingmaximumenergyfromthemysticdungeon

func maximumEnergy(energy []int, k int) int {
	max := -10000
	n := len(energy)
	for i := k; i < n; i++ {
		val := energy[i] + energy[i-k]
		if val > energy[i] {
			energy[i] = val
		}
	}
	for i := n - 1; i >= n-k; i-- {
		if energy[i] > max {
			max = energy[i]
		}
	}
	return max
}
