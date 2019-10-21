package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	var res string
	var temp string
	var len int
	var th bool
	var isNegative bool
	len = 0
	if n == -9223372036854775808 {
		z01.PrintRune('-')
		z01.PrintRune('9')
		z01.PrintRune('2')
		z01.PrintRune('2')
		z01.PrintRune('3')
		z01.PrintRune('3')
		z01.PrintRune('7')
		z01.PrintRune('2')
		z01.PrintRune('0')
		z01.PrintRune('3')
		z01.PrintRune('6')
		z01.PrintRune('8')
		z01.PrintRune('5')
		z01.PrintRune('4')
		z01.PrintRune('7')
		z01.PrintRune('7')
		z01.PrintRune('5')
		z01.PrintRune('8')
		z01.PrintRune('0')
		z01.PrintRune('8')
		th = true
	}
	if n == 0 {
		z01.PrintRune('0')
	}
	for i := 0; i <= 21; i++ {
		if n < 0 {
			isNegative = true
			n = n * -1
		}
		if n > 0 {
			if n%10 == 0 {
				temp = "0"
			} else if n%10 == 1 {
				temp = "1"
			} else if n%10 == 2 {
				temp = "2"
			} else if n%10 == 3 {
				temp = "3"
			} else if n%10 == 4 {
				temp = "4"
			} else if n%10 == 5 {
				temp = "5"
			} else if n%10 == 6 {
				temp = "6"
			} else if n%10 == 7 {
				temp = "7"
			} else if n%10 == 8 {
				temp = "8"
			} else if n%10 == 9 {
				temp = "9"
			}
			res = temp + res
			n = n / 10
			len = len + 1
		}
	}
	if th == true {
		isNegative = false
	}
	if isNegative == true {
		z01.PrintRune('-')
	}
	for j := 0; j < len; j++ {
		if res[j] == '0' {
			z01.PrintRune('0')
		} else if res[j] == '1' {
			z01.PrintRune('1')
		} else if res[j] == '2' {
			z01.PrintRune('2')
		} else if res[j] == '3' {
			z01.PrintRune('3')
		} else if res[j] == '4' {
			z01.PrintRune('4')
		} else if res[j] == '5' {
			z01.PrintRune('5')
		} else if res[j] == '6' {
			z01.PrintRune('6')
		} else if res[j] == '7' {
			z01.PrintRune('7')
		} else if res[j] == '8' {
			z01.PrintRune('8')
		} else if res[j] == '9' {
			z01.PrintRune('9')
		}
	}
}
