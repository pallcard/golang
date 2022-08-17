package main

// a  .-
// b  -...
// c  -.-.
func uniqueMorseRepresentations(words []string) int {
	set := map[string]byte{}
	wordArr := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	for _, word := range words {
		str := ""
		for _, char := range word {
			str += wordArr[char-'a']
		}
		set[str] = 0
	}
	return len(set)
}
