package main

import "fmt"

func main() {
	x := 10
	pointerToX := &x
	fmt.Println(pointerToX)
	fmt.Println(*pointerToX)
	z := 5 + *pointerToX
	fmt.Println(z)

	var y *int
	fmt.Println(y == nil)

	var a = new(int)
	fmt.Println(a == nil)
	fmt.Println(*a)

	type person struct {
		FirstName  string
		MiddleName *string
		LastName   string
	}

	p := person{
		FirstName:  "Pat",
		MiddleName: stringp("Perry"),
		LastName:   "Peterson",
	}
	fmt.Println(p)
	fmt.Println(*p.MiddleName)
}

func stringp(s string) *string {
	return &s
}
