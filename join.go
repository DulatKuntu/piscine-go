package piscine

func Join(strs []string, sep string) string {
	var str string
	for i, word := range strs {
		if i == 0 {
			str = str + word
		} else {
			str = str + sep + word
		}
	}
	return str
}
