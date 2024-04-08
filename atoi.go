package piscine

/*
Instructions

    Write a function that simulates the behaviour of the Atoi function in Go. Atoi transforms a number represented as a string in a number represented as an int.

    Atoi returns 0 if the string is not considered as a valid number. For this exercise non-valid string chains will be tested. Some will contain non-digits characters.

    For this exercise the handling of the signs + or - does have to be taken into account.

    This function will only have to return the int. For this exercise the error result of Atoi is not required.

*/

func Atoi(s string) int {
	var b int
	isNeg := false
	for i := 0; i <= len(s)-1; i++ {
		// check for double pos or neg
		if len(s) > 1 {
			if s[i] == '+' && (s[i+1] == '+' || s[i+1] == '-') || s[i] == '-' && (s[i+1] == '-' || s[i+1] == '+') {
				b = 0
				break
			}
		}
		// check for space
		if s[i] == ' ' {
			b = 0
			break
		}

		// check for digits
		if (s[i] >= '0' && s[i] <= '9') || s[i] == '+' || s[i] == '-' {
			if s[i] == '+' {
			} else if s[i] == '-' {
				isNeg = true
			} else {
				b += int(s[i] - 48)
				if i <= len(s)-2 {
					b *= 10
				}
			}
		} else {
			b = 0
			break
		}
	}
	if isNeg {
		return -(b)
	} else {
		return b
	}
}

/*
Here is a possible program to test your function :

package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.Atoi("12345"))
	fmt.Println(piscine.Atoi("0000000012345"))
	fmt.Println(piscine.Atoi("012 345"))
	fmt.Println(piscine.Atoi("Hello World!"))
	fmt.Println(piscine.Atoi("+1234"))
	fmt.Println(piscine.Atoi("-1234"))
	fmt.Println(piscine.Atoi("++1234"))
	fmt.Println(piscine.Atoi("--1234"))
}

And its output :

$ go run .
12345
12345
0
0
1234
-1234
0
0
$

*/
