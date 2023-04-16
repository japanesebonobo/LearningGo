package main

import "fmt"

func main() {
	var person struct {
		name string
		age  int
		pet  string
	}

	person.name = "bob"
	person.age = 50
	person.pet = "dog"

	fmt.Println("person:", person)

	pet := struct {
		name string
		kind string
	}{
		name: "ポチ",
		kind: "dog",
	}
	fmt.Println("pet:", pet)

	type firstPerson struct {
		name string
		age  int
	}

	f := firstPerson{
		name: "Bob",
		age:  50,
	}

	var g struct {
		name string
		age  int
	}

	g = f
	fmt.Println(f == g)
}
