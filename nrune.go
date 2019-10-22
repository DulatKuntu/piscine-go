package piscine

func NRune(s string, n int) rune {
	if n <= 0 {
		return 0
	}
	run := []rune(s)
	return run[n-1]
}
