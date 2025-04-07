package applysubstitutions

import (
	"strings"
)

type Queue []string

func initQueue(cap int) *Queue {
	obj := make(Queue, 0, cap)
	return &obj
}
func (q *Queue) Len() int {
	return len(*q)
}
func (q *Queue) Push(v string) {
	(*q) = append((*q), v)
}
func (q *Queue) Pop() string {
	x := (*q)[0]
	(*q) = (*q)[1:]
	return x
}

func applySubstitutions(replacements [][]string, text string) string {
	data := parseAllReplacements(replacements)
	for key, value := range data {
		text = replaceKeysValue(text, key, value)
	}

	return text
}

func parseAllReplacements(replacements [][]string) map[string]string {
	result := make(map[string]string, len(replacements))

	for _, replacement := range replacements {
		result[replacement[0]] = replacement[1]

	}
	fixReplacementDependency(result)

	return result
}

func fixReplacementDependency(replacements map[string]string) {
	dependencyMap := make(map[string][]string, len(replacements))
	inverseDependencyMap := make(map[string]int, len(replacements))
	for key := range replacements {
		inverseDependencyMap[key] = 0
		validateReplacesValuesDependency(key, replacements, dependencyMap, inverseDependencyMap)
	}
	queObj := initQueue(len(replacements))

	for key, count := range inverseDependencyMap {
		if count == 0 {
			queObj.Push(key)
		}
	}
	for queObj.Len() > 0 {
		key := queObj.Pop()
		for _, dependentKeys := range dependencyMap[key] {
			replacements[dependentKeys] = replaceKeysValue(replacements[dependentKeys], key, replacements[key])
			inverseDependencyMap[dependentKeys]--
			if inverseDependencyMap[dependentKeys] == 0 {
				queObj.Push(dependentKeys)
			}
		}
	}
}
func replaceKeysValue(value, key, replaceText string) string {
	return strings.ReplaceAll(value, "%"+key+"%", replaceText)
}
func validateReplacesValuesDependency(key string, replacements map[string]string, dependencyMap map[string][]string, inverseDependencyMap map[string]int) {
	value := replacements[key]
	b := make([]byte, 0, len(value))
	for i := 0; i < len(value)-2; i++ {
		if value[i] == '%' {
			newKey := string(value[i+1])
			if _, found := replacements[newKey]; found {
				inverseDependencyMap[key]++
				list, found := dependencyMap[newKey]
				if !found {
					list = make([]string, 0)
				}
				list = append(list, key)
				dependencyMap[newKey] = list
				i += 2
				continue
			}
		}
		b = append(b, value[i])
	}
}
