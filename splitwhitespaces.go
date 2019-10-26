package piscine

func SplitWhiteSpaces(str string) []string {
	var s string
	var arr []string
	for _, letter := range str {
		if letter == ' ' {
			arr = append(arr, s)
			s = ""
		} else {
			s = s + string(letter)
		}
	}
	arr = append(arr, s)
	return arr
}
