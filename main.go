package main

import (
	"fmt"
	"github.com/prayogsuns/go-programming/dog"
)

// main dog.main
func main() {
	fmt.Println("Hello, Playgorund")

	// dogYears is total number of dog years equivalent to human years
	dogYears := dog.Years(3)

	// Print dogYears
	fmt.Println(dogYears)

}