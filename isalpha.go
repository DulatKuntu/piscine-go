package piscine

func IsAlpha(str string) bool {
	boo := false
	for _, letter := range str {
		if (letter > 96 && letter < 123) || (letter > 64 && letter < 91) || (letter > 47 && letter < 58) {
			boo = true
		} else {
			return false
		}
	}
	return boo
}
