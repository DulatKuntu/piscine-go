package main

import "github.com/01-edu/z01"

import "os"

func main() {
	var counter int
	for i := range os.Args {
		counter = i
	}
	for i := counter; i > 0; i-- {
		if i > 0 {
			for _, letter := range os.Args[i] {
				z01.PrintRune(letter)
			}
			z01.PrintRune('\n')
		}

	}

}
