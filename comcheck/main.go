package main

import (
	"fmt"
	"os"
)

func main() {
	argsL := len(os.Args)
	alert := []string{"01", "galaxy", "galaxy 01"}
	for _, word := range alert {
		for i := 0; i < argsL; i++ {
			if os.Args[i] == word {
				fmt.Println("Alert!!!")
			}
		}
	}
}
