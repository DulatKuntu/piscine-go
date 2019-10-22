package piscine

func Index(s string, toFind string) int {
	run := FirstRune(toFind)
	for index, item := range s {
		if item == run {
			return index
		}
	}
	return -1
}
