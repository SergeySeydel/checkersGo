package main

import (
	"fmt"
)

func main() {

	s := make([][]int, 8)

	for i := range s {
		s[i] = make([]int, 8)
	}
	flag := 0
	fmt.Print("  ")
	for i := 0; i < 8; i++ {
		r := int('A')
		si := r + i

		fmt.Print(string(rune(si)))
		fmt.Print(" ")

	}
	fmt.Println()
	for i := 0; i < 8; i++ {
		fmt.Print(8 - i)
		fmt.Print(" ")

		for j := 0; j < 8; j++ {

			fmt.Print(flag)
			fmt.Print(" ")
			if flag == 1 {
				flag = 0
			} else {
				flag = 1
			}

			// fmt.Print("flag appended")
			s[i][j] = flag

		}
		if flag == 1 {
			flag = 0
		} else {
			flag = 1
		}

		fmt.Println()

	}

	// for i := range s {
	// 	for j := range s {
	// 		fmt.Print(s[i][j])
	// 		fmt.Print(" ")
	// 	}
	// 	fmt.Println()

	// }
}
