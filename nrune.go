package piscine

func NRune(s string, n int) rune {
	if n <= 0 {
		return
	}
	run := []rune(s)
	return run[n-1]
}
