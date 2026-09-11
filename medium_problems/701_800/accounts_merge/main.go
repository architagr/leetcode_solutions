package accountsmerge

import "sort"

func accountsMerge(accounts [][]string) [][]string {

	emailMap := make(map[string]int, len(accounts)*10)
	indexMap := make(map[int]int, len(accounts))

	for _, email := range accounts[0][1:] {
		emailMap[email] = 0
	}
	for i := 1; i < len(accounts); i++ {
		emails := accounts[i]
		currentIndex := i
		for _, email := range emails[1:] {
			mappedIndex, ok := emailMap[email]
			if ok {
				currentIndex = mappedIndex
			} else {
				emailMap[email] = i
			}
		}
		if currentIndex != i {
			indexMap[i] = currentIndex
			for _, email := range emails[1:] {
				if index := emailMap[email]; index == i {
					accounts[currentIndex] = append(accounts[currentIndex], email)
					emailMap[email] = currentIndex
				}
			}
		}
	}
	result := make([][]string, 0, len(accounts))
	for i, vals := range accounts {
		if _, ok := indexMap[i]; !ok {
			x := make([]string, 0, len(vals))
			x = append(x, vals[0])
			m := make(map[string]bool, len(vals))
			for _, e := range vals[1:] {
				if _, ok := m[e]; !ok {
					x = append(x, e)
					m[e] = true
				}
			}
			sort.Strings(x[1:])
			result = append(result, x)
		}
	}

	return result
}
