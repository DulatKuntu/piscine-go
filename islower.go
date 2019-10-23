package piscine

func IsLower(str string) bool {
	boo := false
	for _, letter := range str {
		if letter > 96 && letter < 123 {
			boo = true
		} else {
			return false
		}
	}
	return boo
}
