package piscine

func IsPrime(nb int) bool {
	if nb == 2 || nb == 3 || nb == 5 || nb == 7 || nb == 11 || nb == 13 {
		return true
	} else if nb%2 != 0 || nb%3 != 0 || nb%5 != 0 || nb%7 != 0 || nb%11 != 0 || nb%13 != 0 {
		return true
	} else {
		return false
	}
}
