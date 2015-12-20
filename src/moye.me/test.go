package main

import "fmt"

//老卵的类
type Person struct {
	ID   string
	Name string
	Sex  string
}

func holy(args ...int) {
	for _, v := range args {
		fmt.Println("arg:", v)
	}
}

func main() {
	var personDB map[string]Person
	personDB = make(map[string]Person)

	personDB["001"] = Person{"001", "Tom", "Male"}
	personDB["002"] = Person{"002", "Lucy", "Female"}

	person, ok := personDB["002"]
	if ok {
		fmt.Println("name: ", person.Name)
	} else {
		fmt.Println("呒找到")
	}

	for _, ps := range personDB {
		fmt.Println("name:", ps.Name)
	}

	holy(1, 2, 3, 4, 5)
}
