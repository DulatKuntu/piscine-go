package piscine

func Capitalize(s string) string {
	var str string
	toUpper := true
	for _, letter := range s {
		if (letter > 96 && letter < 123) || (letter > 64 && letter < 91) || (letter > 47 && letter < 58) {
			if (letter > 96 && letter < 123) && toUpper {
				str = str + string(letter-32)
				toUpper = false
			} else if (letter > 64 && letter < 91) && toUpper {
				str = str + string(letter)
				toUpper = false
			} else if (letter > 64 && letter < 91) && !toUpper {
				str = str + string(letter+32)
			} else {
				str = str + string(letter)
				toUpper = false
			}
		} else {
			str = str + string(letter)
			toUpper = true
		}
	}
	return str
}
