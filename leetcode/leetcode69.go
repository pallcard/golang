package main

func mySqrt2(x int) int {
	left := 0
	right := x
	for {
		mid := left + (right-left)/2
		if mid*mid <= x && x < (mid+1)*(mid+1) {
			return mid
		} else if mid*mid > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
}

func mySqrt(x int) int {
	if x == 1 {
		return 1
	}
	temp := x / 2
	// 10 5 2
	left := true
	for temp > 0 {
		if temp*temp == x {
			return temp
		} else if temp*temp > x {
			if left {
				temp = temp / 2
			} else {
				return temp - 1
			}
		} else {
			temp++
			left = false
		}
	}
	return 0
}
