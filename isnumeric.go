package piscine

func IsNumeric(str string) bool {
	boo := false
	for _, letter := range str {
		if letter > 47 && letter < 58 {
			boo = true
		} else {
			return false
		}
	}
	return boo
}
