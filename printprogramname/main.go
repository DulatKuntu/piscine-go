package main

import "github.com/01-edu/z01"

import "os"

func main() {
	for _, letter := range os.Args[0] {
		z01.PrintRune(letter)
	}
}
