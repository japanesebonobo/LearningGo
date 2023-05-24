package main

import (
	"fmt"
)

func main() {
	samples := []string{"Hello", "apple_π!", "これは感じ文字列"}
outer:
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
			if r == 'l' || r == 'は' {
				continue outer
			}
		}
		fmt.Println()
	}
}
