package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	var first rune = 0
	var second rune = 0
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			for k := '0'; k <= '9'; k++ {
				for l := '0'; l <= '9'; l++ {
					if i != k || j != l {
						second = (l - '0') + 10*(k-'0')
						first = (j - '0') + 10*(i-'0')
						if first < second {
							if (i == '9') && (j == '8') {
								z01.PrintRune(i)
								z01.PrintRune(j)
								z01.PrintRune(' ')
								z01.PrintRune(k)
								z01.PrintRune(l)
								z01.PrintRune('\n')
							} else {
								z01.PrintRune(i)
								z01.PrintRune(j)
								z01.PrintRune(' ')
								z01.PrintRune(k)
								z01.PrintRune(l)
								z01.PrintRune(',')
								z01.PrintRune(' ')
							}
						}
					}
				}
			}
		}
	}
}
