package main

import (
	"fmt"
	"os"
)

func main() {
	// import file from argument
	file, err := os.Open(os.Args[0])
	// print error message
	if err != nil {
		fmt.Printf("%v", err.Error())
	}
	// print "content variable"
	fmt.Println(file.Stat())
}
