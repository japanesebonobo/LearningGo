package main

import (
	"fmt"
)

func main() {
	result := div(5, 2)
	fmt.Println(result)

	MyFunc(MyFancOpts{
		LastName: "Patel",
		Age:      50,
	})
	MyFunc(MyFancOpts{
		FirstName: "Joe",
		LastName:  "Smith",
	})

	fmt.Println(addTo(3))
	fmt.Println(addTo(3, 2))
	fmt.Println(addTo(3, 2, 4, 6, 8))
	a := []int{4, 3}
	fmt.Println(addTo(3, a...))
	fmt.Println(addTo(3, []int{1, 2, 3, 4, 5}...))

	x, y, _ := divAndRemainder(5, 2)
	fmt.Println(x, y)
}

func div(numerator, denominator int) int {
	if denominator == 0 {
		return 0
	}
	return numerator / denominator
}

type MyFancOpts struct {
	FirstName string
	LastName  string
	Age       int
}

func MyFunc(opts MyFancOpts) error {
	fmt.Println(opts)
	fmt.Println("【ここで必要な処理を行う】")
	return nil
}

func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

func divAndRemainder(numerator, denominator int) (result int, remainder int, err error) {
	if denominator == 0 {
		return
	}
	result, remainder = numerator/denominator, numerator%denominator
	return
}
