package main

import (
	"flag"
	"fmt"
	"os"

	ianus "github.com/advancebsd/ianus"
)

func main() {
	helpPtr := flag.Bool("h", false, "help")
	srcPtr := flag.String("s", "", "Source file to convert")
	destPtr := flag.String("o", "", "Destination file for converted file")

	flag.Parse()

	if *helpPtr {
		fmt.Println("Static generator usage")
		fmt.Println("sg -i <input_file> -o <destination file>")
		os.Exit(0)
	}
	if *srcPtr == "" {
		panic("No source file was given")
	}
	if *destPtr == "" {
		panic("No destination file was given")
	}

	source, err := os.ReadFile(*srcPtr)
	if err != nil {
		fmt.Printf("%s is not a valid file!", *srcPtr)
		os.Exit(0)
	}

	fmt.Println(source)


}
