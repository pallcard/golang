package main

func canConstruct(ransomNote string, magazine string) bool {
	if len(magazine) < len(ransomNote) {
		return false
	}

	m := map[int32]int{}
	for _, char := range magazine {
		m[char]++
	}
	for _, char := range ransomNote {
		if _, ok := m[char]; !ok {
			return false
		} else if m[char] > 0 {
			m[char]--
		} else {
			return false
		}
	}
	return true
}
