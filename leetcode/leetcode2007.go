package main

import "sort"

func findOriginalArray(changed []int) []int {
	if len(changed)%2 != 0 {
		return []int{}
	}
	// 1 2 3 4 6 8
	sort.Ints(changed)
	keyMap := map[int]int{}
	res := make([]int, 0, len(changed)/2)
	for i, _ := range changed {
		if _, ok := keyMap[changed[i]/2]; ok && changed[i]%2 == 0 && keyMap[changed[i]/2] > 0 {
			//delete(set, changed[i]/2)
			keyMap[changed[i]/2] -= 1
			res = append(res, changed[i]/2)
		} else {
			keyMap[changed[i]] += 1
		}
	}

	for _, v := range keyMap {
		if v > 0 {
			return []int{}
		}
	}

	return res
}
