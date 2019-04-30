package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{age: 3, name: "Alice"})
	fmt.Println(person{name: "Fred"}) // Ommited Fields will have default value
	fmt.Println(&person{name: "Anna", age: 40})

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name) //Access struct fields with a dot.

	sp := &s // You can also use dots with struct pointers - the pointers are automatically dereferenced.
	fmt.Println(sp.age)

	sp.age = 51

	fmt.Println(sp)

}
