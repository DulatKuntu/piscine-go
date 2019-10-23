package piscine

func ToUpper(s string) string {
	var str string
	for _, letter := range s {
		if letter > 96 && letter < 123 {
			str = str + string(letter-32)
		} else {
			str = str + string(letter)
		}
	}
	return str
}
