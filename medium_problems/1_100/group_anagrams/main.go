package groupanagrams

func groupAnagrams(strs []string) [][]string {
	map1 := make(map[[26]int][]string)

	for _, x := range strs {
		var smth [26]int

		for _, y := range x {
			smth[y-'a']++
		}
		map1[smth] = append(map1[smth], x)
	}

	final := [][]string{}
	for _, x := range map1 {
		final = append(final, x)
	}

	return final
}
