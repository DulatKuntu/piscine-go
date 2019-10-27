package main

import "github.com/01-edu/z01"

import "os"

func printStr(str string) {

	for _, letter := range str {

		z01.PrintRune(letter)

		z01.PrintRune('\n')
	}
}

func isEven(nbr int) bool {
	if nbr%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	var lengthOfArg int
	for i := range os.Args {
		i = i
		lengthOfArg++
	}
	EvenMsg := "I have an odd number of arguments"
	OddMsg := "I have an even number of arguments"
	if isEven(lengthOfArg) {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
