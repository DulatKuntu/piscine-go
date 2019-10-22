package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n < 0 {
		return
	} else if n == 0 {
		z01.PrintRune(rune(0))
		return
	}
	var strrune = []rune{' '}
	for i := 0; n != 0; i++ {
		strrune = append(strrune, rune(n%10+48))
		n = n / 10
	}
	for index := range strrune {
		for j := range strrune {
			if strrune[j] > strrune[index] {
				strrune[j], strrune[index] = strrune[index], strrune[j]
			}
		}
	}
	for index := range strrune {
		z01.PrintRune(strrune[index])
	}
}
