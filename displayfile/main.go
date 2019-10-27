package main

import "os"

import "fmt"

func main() {
	counter := 0

	for i := range os.Args {
		i = i
		counter++
	}

	if counter < 2 {
		fmt.Println("File name missing")
	} else if counter > 2 {
		fmt.Println("Too many arguments")
	} else {

		fmt.Println("Almost there!!")
	}
}
