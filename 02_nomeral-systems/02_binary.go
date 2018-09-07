package main

import "fmt"

func main() {
	// string verbs
	// %d = decima
	// %b = binary
	// %x - hex
	// %q - utf-a
	for i := 0; i < 200; i++ {
		fmt.Printf("%d \t %b \t %#x \t %X \t %q \n", i, i, i, i, i)
	}
}
