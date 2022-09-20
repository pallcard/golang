package main

import "strings"

func findDuplicate(paths []string) [][]string {
	fileMap := map[string][]string{}
	for _, path := range paths {
		split := strings.Split(path, " ")
		if len(split) < 2 {
			continue
		}
		dir := split[0]
		for _, file := range split[1:] {
			fileInfo := strings.Split(file, "(")
			tempFile := make([]string, 0)
			if _, ok := fileMap[fileInfo[1]]; ok {
				tempFile = fileMap[fileInfo[1]]
			}
			tempFile = append(tempFile, dir+"/"+fileInfo[0])
			fileMap[fileInfo[1]] = tempFile
		}

	}

	res := make([][]string, 0)
	for _, v := range fileMap {
		if len(v) <= 1 {
			continue
		}
		res = append(res, v)
	}

	return res
}
