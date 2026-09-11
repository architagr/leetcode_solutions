package topkfrequentwords

import "sort"

func topKFrequent(words []string, k int) []string {
	// map to help count occurence of each word
	m := make(map[string]int)

	// increment count of each word in map
	for i := 0; i < len(words); i++ {
		m[words[i]]++
	}

	// a structure to stor word and its count
	type WordCount struct {
		word  string
		count int
	}
	wordCount := make([]WordCount, 0)

	// convert the map to array to sort the data based on count and in lexicographical order
	for key, val := range m {
		wordCount = append(wordCount, WordCount{
			word:  key,
			count: val,
		})
	}
	// sorting the array having list of words and their frequency in decending order of frequency
	// and in lexicographical order
	sort.Slice(wordCount, func(i, j int) bool {
		if wordCount[i].count == wordCount[j].count {
			return wordCount[i].word < wordCount[j].word
		}
		return wordCount[i].count > wordCount[j].count
	})
	// take only top k words from the array
	result := make([]string, 0, k)
	for i := 0; i < k; i++ {
		result = append(result, wordCount[i].word)
	}
	return result
}
