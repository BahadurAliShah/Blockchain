package main

import (
	"fmt"
	"strings"
)

type StudentRecord struct {
	rollnumber int
	name       string
	address    string
}

func (s *StudentRecord) AddStudent(rollnumber int, name string, address string) *StudentRecord {
	s.rollnumber = rollnumber
	s.name = name
	s.address = address
	return s
}

func printStudent(st []StudentRecord) {
	for i, v := range st {
		fmt.Printf("%s List %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		println("Student Roll Number: ", v.rollnumber)
		println("Student Name: ", v.name)
		println("Student Address: ", v.address)
	}
}

func main() {
	st := make([]StudentRecord, 0)
	st = append(st, *new(StudentRecord).AddStudent(1, "Bob", "USA"))
	st = append(st, *new(StudentRecord).AddStudent(2, "Alice", "UK"))
	st = append(st, *new(StudentRecord).AddStudent(3, "John", "India"))

	printStudent(st)
}
