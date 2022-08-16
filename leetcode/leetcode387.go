package main

func firstUniqChar(s string) int {
	set := map[int32]int{} //char:index
	for i, char := range s {
		_, ok := set[char]
		if ok {
			set[char] = -1 //标识删除
		} else {
			set[char] = i
		}
	}
	min := -1
	for _, v := range set {
		if v >= 0 && (min == -1 || min > v) {
			min = v
		}
	}
	return min
}
