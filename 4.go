package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
)

type StudentRecord struct {
	rollnumber int
	name       string
	address    string
	courses    []string
}

func (s *StudentRecord) AddStudent(rollnumber int, name string, address string, courses []string) *StudentRecord {
	s.rollnumber = rollnumber
	s.name = name
	s.address = address
	s.courses = courses
	return s
}

func printCourses(cr []string) {
	for _, v := range cr {
		fmt.Printf("%s, ", v)
	}
	fmt.Printf("\n")
}

func printStudent(st []StudentRecord) {
	for i, v := range st {
		fmt.Printf("%s List %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		println("Student Roll Number: ", v.rollnumber)
		println("Student Name: ", v.name)
		println("Student Address: ", v.address)
		fmt.Printf("Student Courses: ")
		printCourses(v.courses)
	}
}

func main() {
	st := make([]StudentRecord, 0)
	st = append(st, *new(StudentRecord).AddStudent(1, "Bob", "USA", []string{"Maths", "Physics", "Chemistry"}))
	st = append(st, *new(StudentRecord).AddStudent(2, "Alice", "UK", []string{"Maths", "Physics", "Chemistry"}))
	st = append(st, *new(StudentRecord).AddStudent(3, "John", "India", []string{"Maths", "Physics", "Chemistry"}))

	printStudent(st)

	for _, v := range st {
		stringToHash := v.name + v.address + strconv.Itoa(v.rollnumber) + strings.Join(v.courses, "")
		fmt.Printf("SHA256 of %s is %x\n", v.name, sha256.Sum256([]byte(stringToHash)))
	}
}
