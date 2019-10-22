package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n > 0 {
		var byt [30]int
		counter := 0
		for i := 0; n > 0; i++ {
			byt[i] = n % 10
			n = n / 10
			counter++
			if n == 0 {
				break
			}
		}
		min := byt[0]
		for i := 0; i <= counter; i++ {
			for j := i; j <= counter; j++ {
				if byt[j] < min {
					n = min
					min = byt[j]
					byt[j] = n
				}
			}
			z01.PrintRune(rune(min))
		}
	}
}
