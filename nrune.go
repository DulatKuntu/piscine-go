package piscine

func NRune(s string, n int) rune {
	var run rune
	for index, item := range s {
		if index == (n - 1) {
			run = item
		}
	}

	return run
}
