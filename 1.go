package main

import "fmt"

type Person struct {
	name string
	age  int
}

func printPerson(p Person) {
	fmt.Println(p.name, p.age)
}

func main() {
	p := Person{"Bob", 20}
	printPerson(p)
}
