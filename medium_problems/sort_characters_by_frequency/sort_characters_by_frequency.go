package sort_characters_by_frequency

import "sort"

func frequencySort(s string) string {
	var counts [128]int
	for i := 0; i < len(s); i++ {
		counts[s[i]]++
	}

	type Item struct {
		a     byte
		count int
	}
	var items []Item
	for i := 0; i < len(counts); i++ {
		if counts[i] != 0 {
			items = append(items, Item{byte(i), counts[i]})
		}
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i].count > items[j].count
	})
	res := make([]byte, 0, len(s))
	for _, item := range items {
		for i := 0; i < item.count; i++ {
			res = append(res, item.a)
		}
	}
	return string(res)
}
