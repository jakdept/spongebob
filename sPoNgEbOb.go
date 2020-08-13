package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input := strings.Join(os.Args[1:], " ")
	var output string

	for i := 0; i < len(input); i++ {
		switch i % 2 {
		case 0:
			output += strings.ToLower(input[i : i+1])
		case 1:
			output += strings.ToUpper(input[i : i+1])
		}
	}
	fmt.Println(output)
}
