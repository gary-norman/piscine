package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			for k := '0'; k <= '9'; k++ {
				if j > i && k > j {
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(k)
					if i < '7' {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}

	}
<<<<<<< HEAD
	z01.PrintRune('\n')
}
=======
  z01.PrintRune('\n')
}
>>>>>>> b2a13c4d90c322dc0ce0932715da859901392ca1
