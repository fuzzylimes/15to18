package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Must provide SF Id")
		os.Exit(1)
	}
	input := os.Args[1]
	if len(input) != 15 && len(input) != 18 {
		fmt.Println("Invalid input - must be either 15 or 18 characters")
		os.Exit(1)
	}

	output := Convert(input)

	fmt.Println("SF-15:", output[:15])
	fmt.Println("SF-18:", output)
}

func Convert(input string) string {
	if len(input) == 18 {
		return input
	}

	sum := ""
	for i := 0; i < 3; i++ {
		loop := 0
		for j := 0; j < 5; j++ {
			current := rune(input[i*5+j])
			if current >= 'A' && current <= 'Z' {
				loop = loop + (1 << j)
			}
		}
		sum += string([]rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ012345")[loop])
	}
	return input + sum
}
