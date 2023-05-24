package main

import (
	"fmt"
)

func main() {
	i := 1
	for i < 100 {
		fmt.Println(i)
		i = i * 2
	}

	// for {
	// 	fmt.Println("Hello")
	// }

	// for i := 1; i <= 100; i++ {
	// 	if i%3 == 0 {
	// 		if i%5 == 0 {
	// 			fmt.Println(i, "3でも5でも割り切れる")
	// 		} else {
	// 			fmt.Println(i, "3で割り切れる")
	// 		}
	// 	} else if i%5 == 0 {
	// 		fmt.Println(i, "5で割り切れる")
	// 	} else {
	// 		fmt.Println(i)
	// 	}
	// }

	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, "3でも5でも割り切れる")
			continue
		}
		if i%3 == 0 {
			fmt.Println(i, "3で割り切れる")
			continue
		}
		if i%5 == 0 {
			fmt.Println(i, "5で割り切れる")
			continue
		}
		fmt.Println(i)
	}

	// evenVals := []int{2, 4, 6, 8, 10, 12}
	// for i, v := range evenVals {
	// 	fmt.Println(i, v)
	// }

	evenVals := []int{2, 4, 6, 8, 10, 12}
	for _, v := range evenVals {
		fmt.Println(v)
	}

	uniqueNames := map[string]bool{"花子": true, "太郎": true, "洋子": true}
	for k, v := range uniqueNames {
		fmt.Println(k, v)
	}

	m := map[string]int{
		"a": 1,
		"c": 3,
		"b": 2,
	}

	for i := 0; i < 3; i++ {
		fmt.Println("ループ", i)
		for k, v := range m {
			fmt.Println(k, v)
		}
	}

	samples := []string{"hello", "apple_π!", "これは漢字文字列"}
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
		}
		fmt.Println()
	}

	evenVals2 := []int{2, 4, 6, 8, 10, 12}
	for _, v := range evenVals2 {
		v *= 2
	}
	fmt.Println(evenVals2)
}
