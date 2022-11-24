package main

func searchRange(nums []int, target int) []int {
	left := 0
	right := len(nums) - 1

	mid := -1
	for left <= right {
		mid = left + (right-left)/2
		if nums[mid] == target {
			break
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	// not found
	if left > right {
		return []int{-1, -1}
	}
	leftIndex := mid - 1
	for leftIndex >= 0 {
		if nums[leftIndex] == target {
			leftIndex--
		} else {
			break
		}
	}
	rightIndex := mid + 1
	for rightIndex < len(nums) {
		if nums[rightIndex] == target {
			rightIndex++
		} else {
			break
		}
	}

	return []int{leftIndex + 1, rightIndex - 1}
}
