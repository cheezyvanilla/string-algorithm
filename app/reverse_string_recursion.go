package app

func ReverseString(s string) string {
	runes := []rune(s)
	if s == "" {
		return ""
	}

	return ReverseString(string(runes[1:])) + string(runes[0])
}