package app

import "strconv"

func VowelConsonantCounter(s string) (int, int) {
	// store vowel letter to a map
	// using map with empty struct value's type for more efficient memory
	vowel := map[string]struct{}{}
	vowelArr := []string{"a", "i", "u", "e", "o"}

	for _, v := range vowelArr {
		vowel[v] = struct{}{}
	}

	vowelCount := 0
	consonantCount := 0

	for _, v := range s {
		// skip the char if it is digit or space
		if _, err := strconv.Atoi(string(v)); err == nil || string(v) == " "{
			continue
		}

		if _, ok := vowel[string(v)]; ok {
			vowelCount++
		}else{
			consonantCount++
		}
		
	}
	return vowelCount, consonantCount
}