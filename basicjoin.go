package piscine

func BasicJoin(strs []string) string {
	var str string
	for _, word := range strs {
		str = str + word
	}
	return str
}
