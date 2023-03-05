package main

import (
	"flag"
	"fmt"
)

func main() {
	source := flag.String("s", "", "Source file path")
	destination := flag.String("o", "", "Destination file path")

	flag.Parse()

	fmt.Printf("Source file path: %s\n", *source)
	fmt.Printf("Destination file path: %s\n", *destination)
}
