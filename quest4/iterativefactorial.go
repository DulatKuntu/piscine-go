package piscine

func IterativeFactorial(nb int) int {
	if nb > 0 {
		temp := 1
		for i := 1; i <= nb; i++ {
			temp = temp * i
		}
		return temp
	} else if nb == 0 {
		return 1
	} else {
		return 0
	}
}
