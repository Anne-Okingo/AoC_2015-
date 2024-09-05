package functions

import ()

func CountNiceString(s []string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		if IsNiceString(s[i]) {
			count++
		}
	}
	return count
}

func IsNiceString(s string) bool {
	if Contains(s) {
		if IsVowel(s) && IsDouble(s) {
			return true
		}
	}
	return false
}

func Contains(s string) bool {
	for i := 0; i < len(s); i++ {
		if i < len(s)-1 && (s[i] == 'a' && s[i+1] == 'b' || s[i] == 'c' && s[i+1] == 'd' || s[i] == 'p' && s[i+1] == 'q' || s[i] == 'x' && s[i+1] == 'y') {
			return false
		}
	}
	return true
}

func IsDouble(s string) bool {
	for i := 0; i < len(s); i++ {
		if i < len(s)-1 && s[i] == s[i+1] {
			return true
		}
	}
	return false
}

func IsVowel(s string) bool {
	vowels := "aeiou"
	mapv := make(map[rune]bool)

	for _, ch := range vowels {
		mapv[ch] = true
	}

	count := 0
	for _, c := range s {
		if mapv[c] {
			count++
		}
	}

	return count >= 3
}
