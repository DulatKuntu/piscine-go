package piscine

func LastRune(s string) rune {
	run := []rune(s)
	var last int
	for index := range s {
		last = index
	}

	return run[last]
}
