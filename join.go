package piscine

func Join(strs []string, sep string) string {
	var str string
	for _, word := range strs {
		str = str + word + sep
	}
	return str
}
