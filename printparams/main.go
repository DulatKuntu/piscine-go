package main

import "github.com/01-edu/z01"

import "os"

func main() {
	for i := range os.Args {
		for _, letter := range os.Args[i] {
			z01.PrintRune(letter)
		}
		z01.PrintRune('\n')
	}

}
