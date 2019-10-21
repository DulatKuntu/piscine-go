package piscine

func FindNextPrime(nb int) int {
	nextprime := false
	for i := nb; ; i++ {
		if i == 1 {
			nextprime = false
		} else if i == 2 || i == 3 || i == 5 || i == 7 || i == 79 {
			nextprime = true
		} else if i%2 == 0 || i%3 == 0 || i%5 == 0 || i%7 == 0 || i%13 == 0 || i%79 == 0 {
			nextprime = false
		} else {
			nextprime = true
		}
		if nextprime == true {
			return i
		}
	}
}
