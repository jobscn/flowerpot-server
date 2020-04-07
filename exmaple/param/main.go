package main

import "fmt"

type Foo struct {
	Name string
}

func change(foo Foo) {
	foo.Name = "changed"
	fmt.Println(foo)
}

func main() {
	foo := &Foo{
		Name: "origin",
	}

	change(*foo)

	fmt.Println(foo)
}
