package main

func twoSum(nums []int, target int) []int {
	restIndex := map[int]int{} //rest:index
	for i, num := range nums {
		if index, ok := restIndex[num]; ok {
			return []int{i, index}
		}
		restIndex[target-num] = i
	}
	return nums
}
