package piscine

func IsPrintable(str string) bool {
	for _, letter := range str {
		if letter == '\n' || letter == '\v' || letter == '\t' || letter == '\f' || letter == '\r' {
			return false
		}
	}
	return true
}
