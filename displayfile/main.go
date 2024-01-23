package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// initialise argsLen variable as number of arguments
	argslen := len(os.Args)
	// initialise variable as filename if 1 argument exists
	if argslen == 1 {
		fmt.Println("File name missing")
	} else if argslen > 2 {
		fmt.Println("Too many arguments")
	} else {
		fileName, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			fmt.Printf("%v", err.Error())
		}
		fmt.Printf("%s", fileName) // orint the contents of the file
	}
	// print "content variable"
}
