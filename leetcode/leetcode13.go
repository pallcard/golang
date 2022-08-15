package main

// IX IV
// XV XC
// CD CM
func romanToInt(s string) int {
	sum := 0
	ixcMap := map[int32][]uint8{
		'I': {'X', 'V', 1},
		'X': {'L', 'C', 10},
		'C': {'D', 'M', 100},
	}
	romanMap := map[int32]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	for i, char := range s {
		switch char {
		case 'I', 'X', 'C':
			// IX 当前-I
			if i+1 < len(s) && (s[i+1] == ixcMap[char][0] || s[i+1] == ixcMap[char][1]) {
				sum -= romanMap[char]
			} else {
				sum += romanMap[char]
			}
		default:
			sum += romanMap[char]
		}
	}

	return sum
}
