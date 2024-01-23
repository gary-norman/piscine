package main

import (
	"fmt"
	"os"
)

func main() {
	// import file from argument
	file, err := os.Open(os.Args[1])
	// print error message
	if err != nil {
		fmt.Printf("%v", err.Error())
	}
	// print "content variable"
	println(file.Stat())
}
