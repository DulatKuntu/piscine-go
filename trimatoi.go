package piscine

func TrimAtoi(s string) int {
	var newstr string

	for _, letter := range s {
		if letter == '-' && StrLen(newstr) == 0 {
			newstr = newstr + string(letter)
		}
		if letter > 47 && letter < 58 {
			newstr = newstr + string(letter)
		}
	}
	return Atoi(newstr)
}
