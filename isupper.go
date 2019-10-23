package piscine

func IsUpper(str string) bool {
	boo := false
	for _, letter := range str {
		if letter > 64 && letter < 91 {
			boo = true
		} else {
			return false
		}
	}
	return boo
}
