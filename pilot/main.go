package main

import "fmt"

func main() {
	var donnie struct {
		Name     string
		Life     float32
		Age      int
		Aircraft int
	}
	donnie.Name = "Donnie"
	donnie.Life = 100.0
	donnie.Age = 24
	donnie.Aircraft = AIRCRAFT1

	fmt.Println(donnie)
}

const AIRCRAFT1 = 1
