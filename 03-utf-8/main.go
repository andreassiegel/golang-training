package main

import "fmt"

func main() {
	for i := 60; i < 123; i++ {
		fmt.Printf("%d \t %b \t %#X \t %q \n", i, i, i, i)
	}
}
