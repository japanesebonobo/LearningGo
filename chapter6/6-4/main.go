package main

type Foo struct {
	Field1 string
	Field2 int
}

func MakeFoo(f *Foo) error {
	f.Field1 = "val"
	f.Field2 = 20
	return nil
}

func MakeFoo2() (Foo, error) {
	f := Foo{
		Field1: "val",
		Field2: 20,
	}
	return f, nil
}

func main() {
	_, _ = MakeFoo2()
}
