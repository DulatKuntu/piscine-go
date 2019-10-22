package piscine

func TrimAtoi(s string) int {
	var newstr string
	
	for _,letter:= range s {
		if (letter > 47 && letter < 56) || letter == '-' {
			newstr = newstr + string(letter)
		}
	}
	return Atoi(newstr)
}