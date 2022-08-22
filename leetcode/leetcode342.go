package main

func isPowerOfFour(n int) bool {
	if n == 0 {
		return false
	}
	for n != 1 {
		if n%4 != 0 {
			return false
		}
		n = n / 4
	}
	return true

	//return (n & (n - 1)) == 0 && n % 3 == 1;
}
