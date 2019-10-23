package piscine

func ToLower(s string) string {
	var str string
	for _, letter := range s {
		if letter > 64 && letter < 91 {
			str = str + string(letter+32)
		} else {
			str = str + string(letter)
		}
	}
	return str
}
