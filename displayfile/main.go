package main

import "os"

import "fmt"

func main() {
	counter := 0
	res := false
	for _, word := range os.Args {
		if word == "quest8.txt" {
			res = true
		}
		counter++
	}
	if counter == 1 {
		fmt.Println("File name missing")
	}
	if counter > 2 {
		fmt.Println("Too many arguments")
	} else {
		if res {
			fmt.Println("Almost there!!")
		}
	}
}
