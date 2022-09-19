package find_duplicate_file_in_system

import (
	"fmt"
	"strings"
)

func FindDuplicate(paths []string) [][]string {
	contentMap := make(map[string][]string)
	// ["root/a 1.txt(abcd) 2.txt(efgh)","root/c 3.txt(abcd)","root/c/d 4.txt(efgh)","root 4.txt(efgh)"]
	for i := 0; i < len(paths); i++ {
		dirPath, contents := getDirContent(paths[i])
		for j := 0; j < len(contents); j++ {
			filePath := fmt.Sprintf("%s/%s", dirPath, contents[j][0])
			contentMap[contents[j][1]] = append(contentMap[contents[j][1]], filePath)
		}
	}
	return removeNonDuplicateEntries(contentMap)
}

func removeNonDuplicateEntries(filePathByContent map[string][]string) [][]string {
	result := make([][]string, 0, len(filePathByContent))

	for _, paths := range filePathByContent {
		if len(paths) > 1 {
			result = append(result, paths)
		}
	}
	return result
}

func getDirContent(data string) (dirPath string, contents [][]string) {
	paths := strings.Split(data, " ")
	dirPath = paths[0]
	contents = make([][]string, len(paths)-1)
	for i := 1; i < len(paths); i++ {
		fileName, content := getFileContent(paths[i])
		contents[i-1] = []string{fileName, content}
	}
	return
}

func getFileContent(data string) (fileName, content string) {
	splitData := strings.Split(data, "(")
	fileName = splitData[0]
	content = string(splitData[1][:len(splitData[1])-1])
	return
}
